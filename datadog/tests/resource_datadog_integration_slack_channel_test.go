package test

import (
	"context"
	"fmt"
	"github.com/terraform-providers/terraform-provider-datadog/datadog"
	"github.com/terraform-providers/terraform-provider-datadog/datadog/internal/utils"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	communityClient "github.com/zorkian/go-datadog-api"
)

func TestAccDatadogIntegrationSlackChannel_Basic(t *testing.T) {
	ctx, accProviders := testAccProviders(context.Background(), t)
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	uniqueChannelAccountName := reg.ReplaceAllString(uniqueEntityName(ctx, t), "")
	accProvider := testAccProvider(t, accProviders)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    accProviders,
		CheckDestroy: testAccCheckDatadogIntegrationSlackChannelDestroy(accProvider),
		Steps: []resource.TestStep{
			{
				// Workaround to ensure we create the slack integration before running the tests.
				Config: emptyLogsArchiveConfig(),
				Check: resource.ComposeTestCheckFunc(
					createSlackIntegration(accProvider),
				),
			},
			{
				Config: testAccCheckDatadogIntegrationSlackChannelConfigCreate(uniqueChannelAccountName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatadogIntegrationSlackChannelExists(accProvider, "datadog_integration_slack_channel.slack_channel"),
					resource.TestCheckResourceAttr(
						"datadog_integration_slack_channel.slack_channel", "channel_name", "#"+uniqueChannelAccountName),
					resource.TestCheckResourceAttr(
						"datadog_integration_slack_channel.slack_channel", "display.0.message", "true"),
					resource.TestCheckResourceAttr(
						"datadog_integration_slack_channel.slack_channel", "display.0.notified", "true"),
					resource.TestCheckResourceAttr(
						"datadog_integration_slack_channel.slack_channel", "display.0.snapshot", "true"),
					resource.TestCheckResourceAttr(
						"datadog_integration_slack_channel.slack_channel", "display.0.tags", "true"),
				),
			},
			{
				Config: testAccCheckDatadogIntegrationSlackChannelConfigUpdate(uniqueChannelAccountName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDatadogIntegrationSlackChannelExists(accProvider, "datadog_integration_slack_channel.slack_channel"),
					resource.TestCheckResourceAttr(
						"datadog_integration_slack_channel.slack_channel", "channel_name", "#"+uniqueChannelAccountName),
					resource.TestCheckResourceAttr(
						"datadog_integration_slack_channel.slack_channel", "display.0.message", "false"),
					resource.TestCheckResourceAttr(
						"datadog_integration_slack_channel.slack_channel", "display.0.notified", "false"),
					resource.TestCheckResourceAttr(
						"datadog_integration_slack_channel.slack_channel", "display.0.snapshot", "false"),
					resource.TestCheckResourceAttr(
						"datadog_integration_slack_channel.slack_channel", "display.0.tags", "false"),
				),
			},
		},
	})
}

func testAccCheckDatadogIntegrationSlackChannelConfigCreate(uniq string) string {
	return fmt.Sprintf(`
       resource "datadog_integration_slack_channel" "slack_channel" {
			display {
				message = true
				notified = true
				snapshot = true
				tags = true
			}
			channel_name = "#%s"
			account_name    = "test_account"
       }
   `, uniq)
}

func testAccCheckDatadogIntegrationSlackChannelConfigUpdate(uniq string) string {
	return fmt.Sprintf(`
       resource "datadog_integration_slack_channel" "slack_channel" {
			display {
				message = false
				notified = false
				snapshot = false
				tags = false
			}
			channel_name = "#%s"
			account_name    = "test_account"
       }
   `, uniq)
}

func emptyLogsArchiveConfig() string {
	return fmt.Sprintf(`
		data "datadog_ip_ranges" "test" {
		}
   `)
}

func testAccCheckDatadogIntegrationSlackChannelExists(accProvider *schema.Provider, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		meta := accProvider.Meta()
		providerConf := meta.(*datadog.ProviderConfiguration)
		datadogClient := providerConf.DatadogClientV1
		auth := providerConf.AuthV1

		accountName := s.RootModule().Resources[resourceName].Primary.Attributes["account_name"]
		channelName := s.RootModule().Resources[resourceName].Primary.Attributes["channel_name"]

		_, _, err := datadogClient.SlackIntegrationApi.GetSlackIntegrationChannel(auth, accountName, channelName).Execute()
		if err != nil {
			return utils.TranslateClientError(err, "error checking slack_channel existence")
		}

		return nil
	}
}

func testAccCheckDatadogIntegrationSlackChannelDestroy(accProvider *schema.Provider) func(*terraform.State) error {
	return func(s *terraform.State) error {
		meta := accProvider.Meta()
		providerConf := meta.(*datadog.ProviderConfiguration)
		datadogClient := providerConf.DatadogClientV1
		auth := providerConf.AuthV1

		for _, r := range s.RootModule().Resources {
			if r.Type != "datadog_slack_channel" {
				continue
			}

			accountName := r.Primary.Attributes["account_name"]
			channelName := r.Primary.Attributes["channel_name"]

			_, resp, err := datadogClient.SlackIntegrationApi.GetSlackIntegrationChannel(auth, accountName, channelName).Execute()

			if err != nil {
				if resp.StatusCode == 404 {
					continue // resource not found => all ok
				} else {
					return fmt.Errorf("received an error retrieving slack_channel: %s", err.Error())
				}
			} else {
				return fmt.Errorf("slack_channel %s still exists", r.Primary.ID)
			}
		}

		return nil
	}
}

func createSlackIntegration(accProvider *schema.Provider) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		providerConf := accProvider.Meta().(*datadog.ProviderConfiguration)
		client := providerConf.CommunityClient
		slackIntegration := communityClient.IntegrationSlackRequest{
			ServiceHooks: []communityClient.ServiceHookSlackRequest{
				{
					Account: communityClient.String("test_account"),
					Url:     communityClient.String("https://ddog-client-test.slack.com/fake-account-hook"),
				},
			},
		}

		if err := client.CreateIntegrationSlack(&slackIntegration); err != nil {
			return err
		}

		return nil
	}
}
