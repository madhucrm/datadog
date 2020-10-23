package datadog

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"testing"

	datadogV1 "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// config
func testAccCheckDatadogServiceLevelObjectiveConfig(uniq string) string {
	return fmt.Sprintf(`
resource "datadog_service_level_objective" "foo" {
  name = "%s"
  type = "metric"
  description = "some description about foo SLO"
  query {
	numerator = "sum:my.metric{type:good}.as_count()"
	denominator = "sum:my.metric{*}.as_count()"
  }

  thresholds {
	timeframe = "7d"
	target = 99.5
	warning = 99.8
  }

  thresholds {
	timeframe = "30d"
	target = 99
  }

  thresholds {
	timeframe = "90d"
	target = 99
  }

  tags = ["foo:bar", "baz"]
}`, uniq)
}

func testAccCheckDatadogServiceLevelObjectiveInvalidMonitorConfig(uniq string) string {
	return fmt.Sprintf(`
resource "datadog_service_level_objective" "bar" {
	name               = "%s"
	type               = "monitor"
	description        = "My custom monitor SLO"
	monitor_ids = [1, 2, 3]
	validate = true
	thresholds {
	timeframe = "7d"
	target = 99.9
	warning = 99.99
	}
	
	thresholds {
	timeframe = "30d"
	target = 99.9
	warning = 99.99
	}
	
	tags = ["foo:bar", "baz"]
}`, uniq)
}

func testAccCheckDatadogServiceLevelObjectiveConfigUpdated(uniq string) string {
	return fmt.Sprintf(`
resource "datadog_service_level_objective" "foo" {
  name = "%s"
  type = "metric"
  description = "some updated description about foo SLO"
  query {
	numerator = "sum:my.metric{type:good}.as_count()"
	denominator = "sum:my.metric{type:good}.as_count() + sum:my.metric{type:bad}.as_count()"
  }

  thresholds {
	timeframe = "7d"
	target = 99.5
	warning = 99.8
  }

  thresholds {
	timeframe = "30d"
	target = 98
	warning = 99.0
  }

  thresholds {
	timeframe = "90d"
	target = 99.9
  }

  tags = ["foo:bar", "baz"]
}`, uniq)
}

func testAccCheckDatadogServiceLevelObjectiveConfigMonitor(uniq string) string {
	return fmt.Sprintf(`
resource "datadog_monitor" "foo" {
  name = "%s"
  type = "query alert"
  message = "some message Notify: @hipchat-channel"
  query = "avg(last_1h):anomalies(avg:system.cpu.system{name:cassandra}, 'basic', 3, direction='above', alert_window='last_5m', interval=20, count_default_zero='true') >= 1"
}

resource "datadog_service_level_objective" "foo" {
  name = "%s"
  type = "monitor"
  description = "some updated description about foo SLO"

  thresholds {
	timeframe = "7d"
	target = 99.5
	warning = 99.8
  }

  monitor_ids = [
    datadog_monitor.foo.id
  ]
}`, uniq, uniq)
}

func testAccCheckDatadogServiceLevelObjectiveConfigForceRecreate(uniq string) string {
	return fmt.Sprintf(`
resource "datadog_monitor" "foo" {
  name = "%s"
  type = "metric alert"
  message = "some message Notify: @hipchat-channel"
  query = "avg(last_1h):avg:aws.ec2.cpu{environment:foo,host:foo} by {host} > 2"
}

resource "datadog_service_level_objective" "foo" {
  name = "%s"
  type = "monitor"
  description = "some updated description about foo SLO"

  thresholds {
	timeframe = "7d"
	target = 99.5
	warning = 99.8
  }

  monitor_ids = [
    datadog_monitor.foo.id
  ]
}`, uniq, uniq)
}

// tests

func TestAccDatadogServiceLevelObjective_Basic(t *testing.T) {
	accProviders, clock, cleanup := testAccProviders(t, initRecorder(t))
	sloName := uniqueEntityName(clock, t)
	sloNameUpdated := sloName + "-updated"
	defer cleanup(t)
	accProvider := testAccProvider(t, accProviders)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    accProviders,
		CheckDestroy: testAccCheckDatadogServiceLevelObjectiveDestroy(accProvider),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDatadogServiceLevelObjectiveConfig(sloName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatadogServiceLevelObjectiveExists(accProvider, "datadog_service_level_objective.foo"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "name", sloName),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "description", "some description about foo SLO"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "type", "metric"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "query.0.numerator", "sum:my.metric{type:good}.as_count()"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "query.0.denominator", "sum:my.metric{*}.as_count()"),
					// Thresholds are a TypeList, that are sorted by timeframe alphabetically.
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "thresholds.#", "3"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "thresholds.0.timeframe", "7d"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "thresholds.0.target", "99.5"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "thresholds.0.warning", "99.8"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "thresholds.1.timeframe", "30d"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "thresholds.1.target", "99"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "thresholds.2.timeframe", "90d"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "thresholds.2.target", "99"),
					// Tags are a TypeSet => use a weird way to access members by their hash
					// TF TypeSet is internally represented as a map that maps computed hashes
					// to actual values. Since the hashes are always the same for one value,
					// this is the way to get them.
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "tags.#", "2"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "tags.2644851163", "baz"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "tags.1750285118", "foo:bar"),
				),
			},
			{
				Config: testAccCheckDatadogServiceLevelObjectiveConfigUpdated(sloNameUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatadogServiceLevelObjectiveExists(accProvider, "datadog_service_level_objective.foo"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "name", sloNameUpdated),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "description", "some updated description about foo SLO"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "type", "metric"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "query.0.numerator", "sum:my.metric{type:good}.as_count()"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "query.0.denominator", "sum:my.metric{type:good}.as_count() + sum:my.metric{type:bad}.as_count()"),
					// Thresholds are a TypeList, that are sorted by timeframe alphabetically.
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "thresholds.#", "3"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "thresholds.0.timeframe", "7d"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "thresholds.0.target", "99.5"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "thresholds.0.warning", "99.8"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "thresholds.1.timeframe", "30d"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "thresholds.1.target", "98"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "thresholds.1.warning", "99"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "thresholds.2.timeframe", "90d"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "thresholds.2.target", "99.9"),
					// Tags are a TypeSet => use a weird way to access members by their hash
					// TF TypeSet is internally represented as a map that maps computed hashes
					// to actual values. Since the hashes are always the same for one value,
					// this is the way to get them.
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "tags.#", "2"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "tags.2644851163", "baz"),
					resource.TestCheckResourceAttr(
						"datadog_service_level_objective.foo", "tags.1750285118", "foo:bar"),
				),
			},
		},
	})
}

func TestAccDatadogServiceLevelObjective_InvalidMonitor(t *testing.T) {
	accProviders, clock, cleanup := testAccProviders(t, initRecorder(t))
	sloName := uniqueEntityName(clock, t)
	defer cleanup(t)
	accProvider := testAccProvider(t, accProviders)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    accProviders,
		CheckDestroy: testAccCheckDatadogServiceLevelObjectiveDestroy(accProvider),
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckDatadogServiceLevelObjectiveInvalidMonitorConfig(sloName),
				ExpectError: regexp.MustCompile("error finding monitor to add to SLO"),
			},
		},
	})
}

func TestAccDatadogServiceLevelObjective_NewMonitorForceRecreate(t *testing.T) {
	accProviders, clock, cleanup := testAccProviders(t, initRecorder(t))
	sloName := uniqueEntityName(clock, t)
	defer cleanup(t)
	accProvider := testAccProvider(t, accProviders)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    accProviders,
		CheckDestroy: testAccCheckDatadogServiceLevelObjectiveDestroy(accProvider),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDatadogServiceLevelObjectiveConfigMonitor(sloName),
				Check: func(firstState *terraform.State) error {
					firstSloId, _ := getSloIdHelper(firstState, accProvider)
					resource.TestCheckResourceAttr("datadog_service_level_objective.foo", "id", firstSloId)
					resource.Test(t, resource.TestCase{
						PreCheck:     func() { testAccPreCheck(t) },
						Providers:    accProviders,
						CheckDestroy: testAccCheckDatadogServiceLevelObjectiveDestroy(accProvider),
						Steps: []resource.TestStep{
							{
								Config: testAccCheckDatadogServiceLevelObjectiveConfigForceRecreate(sloName),
								Check:  checkThatSloHasBeenForcedToBeRecreated(accProvider, firstSloId),
							},
						},
					})
					return nil
				},
			},
		},
	})
}

func checkThatSloHasBeenForcedToBeRecreated(accProvider *schema.Provider, previousSloId string) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		secondSloId, _ := getSloIdHelper(s, accProvider)
		if secondSloId == previousSloId {
			return fmt.Errorf("slo id may have change if the resource as been recreated")
		}
		return nil
	}
}

// helpers

func getSloIdHelper(s *terraform.State, accProvider *schema.Provider) (string, error) {
	providerConf := accProvider.Meta().(*ProviderConfiguration)
	datadogClientV1 := providerConf.DatadogClientV1
	authV1 := providerConf.AuthV1
	for _, r := range s.RootModule().Resources {
		sloResp, _, err := datadogClientV1.ServiceLevelObjectivesApi.GetSLO(authV1, r.Primary.ID).Execute()
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "not found") {
				continue
			}
			return "", fmt.Errorf("received an error retrieving service level objective  %s", err)
		}
		data := sloResp.GetData()
		return data.GetId(), err
	}
	return "", fmt.Errorf("service level objective not found in current state")
}

func destroyServiceLevelObjectiveHelper(s *terraform.State, authV1 context.Context, datadogClientV1 *datadogV1.APIClient) error {
	for _, r := range s.RootModule().Resources {
		if r.Primary.ID != "" {
			if _, _, err := datadogClientV1.ServiceLevelObjectivesApi.GetSLO(authV1, r.Primary.ID).Execute(); err != nil {
				if strings.Contains(strings.ToLower(err.Error()), "not found") {
					continue
				}
				return fmt.Errorf("received an error retrieving service level objective %s", err)
			}
			return fmt.Errorf("service Level Objective still exists")
		}
	}
	return nil
}

func existsServiceLevelObjectiveHelper(s *terraform.State, authV1 context.Context, datadogClientV1 *datadogV1.APIClient) error {
	for _, r := range s.RootModule().Resources {
		if _, _, err := datadogClientV1.ServiceLevelObjectivesApi.GetSLO(authV1, r.Primary.ID).Execute(); err != nil {
			return fmt.Errorf("received an error retrieving service level objective %s", err)
		}
	}
	return nil
}

func testAccCheckDatadogServiceLevelObjectiveDestroy(accProvider *schema.Provider) func(*terraform.State) error {
	return func(s *terraform.State) error {
		providerConf := accProvider.Meta().(*ProviderConfiguration)
		datadogClientV1 := providerConf.DatadogClientV1
		authV1 := providerConf.AuthV1
		if err := destroyServiceLevelObjectiveHelper(s, authV1, datadogClientV1); err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckDatadogServiceLevelObjectiveExists(accProvider *schema.Provider, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		providerConf := accProvider.Meta().(*ProviderConfiguration)
		datadogClientV1 := providerConf.DatadogClientV1
		authV1 := providerConf.AuthV1

		if err := existsServiceLevelObjectiveHelper(s, authV1, datadogClientV1); err != nil {
			return err
		}
		return nil
	}
}
