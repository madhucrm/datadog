package datadog

import (
	"context"
	"strconv"

	"github.com/terraform-providers/terraform-provider-datadog/datadog/internal/utils"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDatadogLogsPipelines() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to list all existing logs pipelines for use in other resources.",
		ReadContext: dataSourceDatadogLogsPipelinesRead,
		Schema: map[string]*schema.Schema{
			"is_read_only": {
				Description: "Filter parameter for retrieved pipelines",
				Type:        schema.TypeString,
				Optional:    true,
			},
			// Computed values
			"logs_pipelines": {
				Description: "List of logs pipelines",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "ID of the pipeline",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"filter": {
							Description: "Pipelines filter",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"query": {
										Description: "Pipeline filter criteria.",
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"name": {
							Description: "The name of the pipeline.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"is_enabled": {
							Description: "Whether or not the pipeline is enabled.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
						"is_read_only": {
							Description: "Whether or not the pipeline can be edited.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
						"type": {
							Description: "Whether or not the pipeline can be edited.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func buildTerraformLogsPipelineFilter(ddFilter datadogV1.LogsFilter) *[]map[string]interface{} {
	tfFilter := map[string]interface{}{
		"query": ddFilter.GetQuery(),
	}
	return &[]map[string]interface{}{tfFilter}
}

func dataSourceDatadogLogsPipelinesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	providerConf := meta.(*ProviderConfiguration)
	apiInstances := providerConf.DatadogApiInstances
	auth := providerConf.Auth
	logsPipelines, httpresp, err := apiInstances.GetLogsPipelinesApiV1().ListLogsPipelines(auth)
	if err != nil {
		return utils.TranslateClientErrorDiag(err, httpresp, "error querying log pipelines")
	}
	if err := utils.CheckForUnparsed(logsPipelines); err != nil {
		return diag.FromErr(err)
	}

	vStr, ok := d.GetOk("is_read_only")
	v, _ := strconv.ParseBool(vStr.(string))
	tflogsPipelines := make([]map[string]interface{}, 0)
	for _, pipeline := range logsPipelines {
		if !ok || (ok && v == *pipeline.IsReadOnly) {
			tflogsPipelines = append(tflogsPipelines, map[string]interface{}{
				"name":         pipeline.GetName(),
				"id":           pipeline.GetId(),
				"filter":       buildTerraformLogsPipelineFilter(pipeline.GetFilter()),
				"is_enabled":   pipeline.GetIsEnabled(),
				"is_read_only": pipeline.GetIsReadOnly(),
				"type":         pipeline.GetType(),
			})
		}
	}
	if err := d.Set("logs_pipelines", tflogsPipelines); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("log-pipelines")
	return nil
}
