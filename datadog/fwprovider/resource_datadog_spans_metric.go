package fwprovider

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	frameworkPath "github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/terraform-providers/terraform-provider-datadog/datadog/internal/utils"
)

var (
	_ resource.ResourceWithConfigure   = &SpansMetricResource{}
	_ resource.ResourceWithImportState = &SpansMetricResource{}
)

type SpansMetricResource struct {
	Api  *datadogV2.SpansMetricsApi
	Auth context.Context
}

type SpansMetricModel struct {
	ID      types.String    `tfsdk:"id"`
	GroupBy []*GroupByModel `tfsdk:"group_by"`
	Compute *ComputeModel   `tfsdk:"compute"`
	Filter  *FilterModel    `tfsdk:"filter"`
}

type GroupByModel struct {
	Path    types.String `tfsdk:"path"`
	TagName types.String `tfsdk:"tag_name"`
}

type ComputeModel struct {
	AggregationType    types.String `tfsdk:"aggregation_type"`
	IncludePercentiles types.Bool   `tfsdk:"include_percentiles"`
	Path               types.String `tfsdk:"path"`
}

type FilterModel struct {
	Query types.String `tfsdk:"query"`
}

func NewSpansMetricResource() resource.Resource {
	return &SpansMetricResource{}
}

func (r *SpansMetricResource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	providerData, ok := request.ProviderData.(*FrameworkProvider)
	if !ok {
		response.Diagnostics.AddError("Unexpected Resource Configure Type", "")
		return
	}

	r.Api = providerData.DatadogApiInstances.GetSpansMetricsApiV2()
	r.Auth = providerData.Auth
}

func (r *SpansMetricResource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "spans_metric"
}

func (r *SpansMetricResource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description: "Provides a Datadog SpansMetric resource. This can be used to create and manage Datadog spans_metric.",
		Attributes: map[string]schema.Attribute{
			"id": utils.ResourceIDAttribute(),
		},
		Blocks: map[string]schema.Block{
			"group_by": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"path": schema.StringAttribute{
							Optional:    true,
							Description: "The path to the value the span-based metric will be aggregated over.",
						},
						"tag_name": schema.StringAttribute{
							Optional:    true,
							Description: "Eventual name of the tag that gets created. By default, the path attribute is used as the tag name.",
						},
					},
				},
			},
			"compute": schema.SingleNestedBlock{
				Attributes: map[string]schema.Attribute{
					"aggregation_type": schema.StringAttribute{
						Optional:    true,
						Description: "The type of aggregation to use.",
					},
					"include_percentiles": schema.BoolAttribute{
						Optional:    true,
						Description: "Toggle to include or exclude percentile aggregations for distribution metrics. Only present when the `aggregation_type` is `distribution`.",
					},
					"path": schema.StringAttribute{
						Optional:    true,
						Description: "The path to the value the span-based metric will aggregate on (only used if the aggregation type is a \"distribution\").",
					},
				},
			},
			"filter": schema.SingleNestedBlock{
				Attributes: map[string]schema.Attribute{
					"query": schema.StringAttribute{
						Optional:    true,
						Description: "The search query - following the span search syntax.",
					},
				},
			},
		},
	}
}

func (r *SpansMetricResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, frameworkPath.Root("id"), request, response)
}

func (r *SpansMetricResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var state SpansMetricModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	id := state.ID.ValueString()
	resp, httpResp, err := r.Api.GetSpansMetric(r.Auth, id)
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			response.State.RemoveResource(ctx)
			return
		}
		response.Diagnostics.Append(utils.FrameworkErrorDiag(err, "error retrieving API Key"))
		return
	}
	if err := utils.CheckForUnparsed(resp); err != nil {
		response.Diagnostics.AddError("response contains unparsedObject", err.Error())
		return
	}

	r.updateState(ctx, &state, &resp)

	// Save data into Terraform state
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (r *SpansMetricResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var state SpansMetricModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	body, diags := r.buildSpansMetricRequestBody(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	resp, _, err := r.Api.CreateSpansMetric(r.Auth, *body)
	if err != nil {
		response.Diagnostics.Append(utils.FrameworkErrorDiag(err, "error retrieving SpansMetric"))
		return
	}
	if err := utils.CheckForUnparsed(resp); err != nil {
		response.Diagnostics.AddError("response contains unparsedObject", err.Error())
		return
	}
	r.updateState(ctx, &state, &resp)

	// Save data into Terraform state
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (r *SpansMetricResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var state SpansMetricModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	id := state.ID.ValueString()

	body, diags := r.buildSpansMetricUpdateRequestBody(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	resp, _, err := r.Api.UpdateSpansMetric(r.Auth, id, *body)
	if err != nil {
		response.Diagnostics.Append(utils.FrameworkErrorDiag(err, "error retrieving SpansMetric"))
		return
	}
	if err := utils.CheckForUnparsed(resp); err != nil {
		response.Diagnostics.AddError("response contains unparsedObject", err.Error())
		return
	}
	r.updateState(ctx, &state, &resp)

	// Save data into Terraform state
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (r *SpansMetricResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var state SpansMetricModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	id := state.ID.ValueString()

	_, httpResp, err := r.Api.DeleteSpansMetric(r.Auth, id)
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			return
		}
		response.Diagnostics.Append(utils.FrameworkErrorDiag(err, "error deleting spans_metric"))
		return
	}
}

func (r *SpansMetricResource) updateState(ctx context.Context, state *SpansMetricModel, resp *datadogV2.SpansMetricResponse) {
	state.ID = types.StringValue(resp.Data.GetId())

	data := resp.GetData()
	attributes := data.GetAttributes()

	if groupBy, ok := attributes.GetGroupByOk(); ok && len(*groupBy) > 0 {
		state.GroupBy = []*GroupByModel{}
		for _, groupByDd := range *groupBy {
			groupByTfItem := GroupByModel{}
			if path, ok := groupByDd.GetPathOk(); ok {
				groupByTfItem.Path = types.StringValue(*path)
			}
			if tagName, ok := groupByDd.GetTagNameOk(); ok {
				groupByTfItem.TagName = types.StringValue(*tagName)
			}

			state.GroupBy = append(state.GroupBy, &groupByTfItem)
		}
	}

	if compute, ok := attributes.GetComputeOk(); ok {
		computeTf := ComputeModel{}
		if aggregationType, ok := compute.GetAggregationTypeOk(); ok {
			computeTf.AggregationType = types.StringValue(*aggregationType)
		}
		if includePercentiles, ok := compute.GetIncludePercentilesOk(); ok {
			computeTf.IncludePercentiles = types.BoolValue(*includePercentiles)
		}
		if path, ok := compute.GetPathOk(); ok {
			computeTf.Path = types.StringValue(*path)
		}

		state.Compute = &computeTf
	}

	if filter, ok := attributes.GetFilterOk(); ok {
		filterTf := FilterModel{}
		if query, ok := filter.GetQueryOk(); ok {
			filterTf.Query = types.StringValue(*query)
		}

		state.Filter = &filterTf
	}
}

func (r *SpansMetricResource) buildSpansMetricRequestBody(ctx context.Context, state *SpansMetricModel) (*datadogV2.SpansMetricCreateRequest, diag.Diagnostics) {
	diags := diag.Diagnostics{}
	attributes := datadogV2.NewSpansMetricCreateAttributesWithDefaults()

	if state.GroupBy != nil {
		var groupBy []datadogV2.SpansMetricGroupBy
		for _, groupByTFItem := range state.GroupBy {
			groupByDDItem := datadogV2.NewSpansMetricGroupBy()

			groupByDDItem.SetPath(groupByTFItem.Path.ValueString())
			if !groupByTFItem.TagName.IsNull() {
				groupByDDItem.SetTagName(groupByTFItem.TagName.ValueString())
			}
		}
		attributes.SetGroupBy(groupBy)
	}

	var compute datadogV2.SpansMetricCompute

	compute.SetAggregationType(state.Compute.AggregationType.ValueString())
	if !state.Compute.IncludePercentiles.IsNull() {
		compute.SetIncludePercentiles(state.Compute.IncludePercentiles.ValueBool())
	}
	if !state.Compute.Path.IsNull() {
		compute.SetPath(state.Compute.Path.ValueString())
	}

	attributes.SetCompute(compute)

	if state.Filter != nil {
		var filter datadogV2.SpansMetricFilter

		if !state.Filter.Query.IsNull() {
			filter.SetQuery(state.Filter.Query.ValueString())
		}

		attributes.SetFilter(filter)
	}

	req := datadogV2.NewSpansMetricCreateRequestWithDefaults()
	req.Data = *datadogV2.NewSpansMetricCreateDataWithDefaults()
	req.Data.SetAttributes(*attributes)

	return req, diags
}

func (r *SpansMetricResource) buildSpansMetricUpdateRequestBody(ctx context.Context, state *SpansMetricModel) (*datadogV2.SpansMetricUpdateRequest, diag.Diagnostics) {
	diags := diag.Diagnostics{}
	attributes := datadogV2.NewSpansMetricUpdateAttributesWithDefaults()

	if state.GroupBy != nil {
		var groupBy []datadogV2.SpansMetricGroupBy
		for _, groupByTFItem := range state.GroupBy {
			groupByDDItem := datadogV2.NewSpansMetricGroupBy()

			groupByDDItem.SetPath(groupByTFItem.Path.ValueString())
			if !groupByTFItem.TagName.IsNull() {
				groupByDDItem.SetTagName(groupByTFItem.TagName.ValueString())
			}
		}
		attributes.SetGroupBy(groupBy)
	}

	if state.Compute != nil {
		var compute datadogV2.SpansMetricUpdateCompute

		if !state.Compute.IncludePercentiles.IsNull() {
			compute.SetIncludePercentiles(state.Compute.IncludePercentiles.ValueBool())
		}

		attributes.SetCompute(compute)
	}

	if state.Filter != nil {
		var filter datadogV2.SpansMetricFilter

		if !state.Filter.Query.IsNull() {
			filter.SetQuery(state.Filter.Query.ValueString())
		}

		attributes.SetFilter(filter)
	}

	req := datadogV2.NewSpansMetricUpdateRequestWithDefaults()
	req.Data = *datadogV2.NewSpansMetricUpdateDataWithDefaults()
	req.Data.SetAttributes(*attributes)

	return req, diags
}
