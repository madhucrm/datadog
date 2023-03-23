package datadog

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/terraform-providers/terraform-provider-datadog/datadog/internal/utils"
)

func dataSourceDatadogServiceLevelObjective() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve information about an existing SLO for use in other resources.",
		ReadContext: dataSourceDatadogServiceLevelObjectiveRead,
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
			"description": {
				Description: "The description of the service level objective.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"target_threshold": {
				Description: "The primary target threshold of the service level objective.",
				Type:        schema.TypeFloat,
				Computed:    true,
			},
			"warning_threshold": {
				Description: "The primary warning threshold of the service level objective.",
				Type:        schema.TypeFloat,
				Computed:    true,
			},
			"timeframe": {
				Description: "The primary timeframe of the service level objective.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"query_numerator": {
				Description: "The numerator query of the service level objective.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"query_denominator": {
				Description: "The denominator query of the service level objective.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceDatadogServiceLevelObjectiveRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	providerConf := meta.(*ProviderConfiguration)
	apiInstances := providerConf.DatadogApiInstances
	auth := providerConf.Auth

	reqParams := datadogV1.NewListSLOsOptionalParameters()
	if v, ok := d.GetOk("id"); ok {
		reqParams.WithIds(v.(string))
	}
	if v, ok := d.GetOk("name_query"); ok {
		reqParams.WithQuery(v.(string))
	}
	if v, ok := d.GetOk("tags_query"); ok {
		reqParams.WithTagsQuery(v.(string))
	}
	if v, ok := d.GetOk("metrics_query"); ok {
		reqParams.WithMetricsQuery(v.(string))
	}

	slosResp, httpresp, err := apiInstances.GetServiceLevelObjectivesApiV1().ListSLOs(auth, *reqParams)
	if err != nil {
		return utils.TranslateClientErrorDiag(err, httpresp, "error querying service level objectives")
	}
	if err := utils.CheckForUnparsed(slosResp); err != nil {
		return diag.FromErr(err)
	}
	if len(slosResp.GetData()) > 1 {
		return diag.Errorf("your query returned more than one result, please try a more specific search criteria")
	}
	if len(slosResp.GetData()) == 0 {
		return diag.Errorf("your query returned no result, please try a less specific search criteria")
	}

	slo := slosResp.GetData()[0]

	d.SetId(slo.GetId())
	if err := d.Set("name", slo.GetName()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("type", slo.GetType()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", slo.GetDescription()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("target_threshold", slo.GetTargetThreshold()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("warning_threshold", slo.GetWarningThreshold()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("timeframe", slo.GetTimeframe()); err != nil {
		return diag.FromErr(err)
	}
	sloQ := slo.GetQuery()
	if err := d.Set("query_numerator", sloQ.GetNumerator()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("query_denominator", sloQ.GetDenominator()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
