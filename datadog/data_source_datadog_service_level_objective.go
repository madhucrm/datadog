package datadog

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-datadog/datadog/internal/utils"
)

func dataSourceDatadogServiceLevelObjective() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve information about an existing SLO for use in other resources.",
		Read:        dataSourceDatadogServiceLevelObjectiveRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "A SLO ID to limit the search.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name_query": {
				Description: "Filter results based on SLO names.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tags_query": {
				Description: "Filter results based on a single SLO tag.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"metrics_query": {
				Description: "Filter results based on SLO numerator and denominator.",
				Type:        schema.TypeString,
				Optional:    true,
			},

			// Computed values
			"name": {
				Description: "Name of the Datadog service level objective",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description: "The type of the service level objective. The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API [documentation page](https://docs.datadoghq.com/api/v1/service-level-objectives/#create-a-slo-object). Available values are: `metric` and `monitor`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceDatadogServiceLevelObjectiveRead(d *schema.ResourceData, meta interface{}) error {
	providerConf := meta.(*ProviderConfiguration)
	datadogClientV1 := providerConf.DatadogClientV1
	authV1 := providerConf.AuthV1

	req := datadogClientV1.ServiceLevelObjectivesApi.ListSLOs(authV1)
	if v, ok := d.GetOk("id"); ok {
		req = req.Ids(v.(string))
	}
	if v, ok := d.GetOk("name_query"); ok {
		req = req.Query(v.(string))
	}
	if v, ok := d.GetOk("tags_query"); ok {
		req = req.TagsQuery(v.(string))
	}
	if v, ok := d.GetOk("metrics_query"); ok {
		req = req.MetricsQuery(v.(string))
	}

	slosResp, _, err := req.Execute()
	if err != nil {
		return utils.TranslateClientError(err, "error querying service level objectives")
	}
	if len(slosResp.GetData()) > 1 {
		return fmt.Errorf("your query returned more than one result, please try a more specific search criteria")
	}
	if len(slosResp.GetData()) == 0 {
		return fmt.Errorf("your query returned no result, please try a less specific search criteria")
	}

	slo := slosResp.GetData()[0]

	d.SetId(slo.GetId())
	if err := d.Set("name", slo.GetName()); err != nil {
		return err
	}
	if err := d.Set("type", slo.GetType()); err != nil {
		return err
	}

	return nil
}
