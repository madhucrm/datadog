package datadog

import (
	"context"
	"encoding/json"

	"github.com/terraform-providers/terraform-provider-datadog/datadog/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const serviceDefinitionPath = "/api/v2/services/definitions"

type responseDataAttributes struct {
	Schema map[string]interface{}
}

type responseData struct {
	Attributes responseDataAttributes
}

type responseListData struct {
	Data []responseData
}

type responseSingleData struct {
	Data responseData
}

func resourceDatadogServiceCatalogJSON() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides a Datadog service catalog JSON resource. This can be used to create and manage Datadog service definitions using the JSON definition.",
		CreateContext: resourceDatadogServiceCatalogJSONCreate,
		ReadContext:   resourceDatadogServiceCatalogJSONRead,
		UpdateContext: resourceDatadogServiceCatalogJSONCreate,
		DeleteContext: resourceDatadogServiceCatalogJSONDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		CustomizeDiff: customdiff.ForceNewIfChange("definition", func(ctx context.Context, old, new, meta interface{}) bool {
			oldAttrMap, _ := structure.ExpandJsonFromString(old.(string))
			newAttrMap, _ := structure.ExpandJsonFromString(new.(string))

			oldType, ok := oldAttrMap["dd-service"].(string)
			if !ok {
				return true
			}

			newType, ok := newAttrMap["dd-service"].(string)
			if !ok {
				return true
			}

			return oldType != newType
		}),
		Schema: map[string]*schema.Schema{
			"definition": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsJSON,
				Description:  "The JSON formatted definition of the service.",
			},
		},
	}
}

func resourceDatadogServiceCatalogJSONRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	providerConf := meta.(*ProviderConfiguration)
	datadogClientV1 := providerConf.DatadogClientV1
	authV2 := providerConf.AuthV1

	id := d.Id()
	respByte, httpResp, err := utils.SendRequest(authV2, datadogClientV1, "GET", serviceDefinitionPath+"/"+id, nil)
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	var responseData responseSingleData
	err = json.Unmarshal(respByte, &responseData)
	if err != nil {
		return diag.FromErr(err)
	}

	return updateServiceCatalogJSONState(d, responseData.Data)
}

func resourceDatadogServiceCatalogJSONCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	providerConf := meta.(*ProviderConfiguration)
	datadogClientV1 := providerConf.DatadogClientV1
	authV2 := providerConf.AuthV1

	definition := d.Get("definition").(string)

	respByte, httpresp, err := utils.SendRequest(authV2, datadogClientV1, "POST", serviceDefinitionPath, &definition)
	if err != nil {
		return utils.TranslateClientErrorDiag(err, httpresp, "error creating resource")
	}

	var responseData responseListData
	err = json.Unmarshal(respByte, &responseData)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(responseData.Data[0].Attributes.Schema["dd-service"].(string))

	return updateServiceCatalogJSONState(d, responseData.Data[0])
}

func resourceDatadogServiceCatalogJSONDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	providerConf := meta.(*ProviderConfiguration)
	datadogClientV1 := providerConf.DatadogClientV1
	authV2 := providerConf.AuthV1
	id := d.Id()
	_, httpResp, err := utils.SendRequest(authV2, datadogClientV1, "DELETE", serviceDefinitionPath+"/"+id, nil)
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}
	return nil
}

func updateServiceCatalogJSONState(d *schema.ResourceData, response responseData) diag.Diagnostics {
	serviceString, err := structure.FlattenJsonToString(response.Attributes.Schema)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("definition", serviceString); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
