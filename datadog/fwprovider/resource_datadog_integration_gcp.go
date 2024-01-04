package fwprovider

import (
	"context"
	"sync"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	frameworkPath "github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/terraform-providers/terraform-provider-datadog/datadog/internal/utils"
)

var (
	_ resource.ResourceWithConfigure   = &integrationGcpResource{}
	_ resource.ResourceWithImportState = &integrationGcpResource{}
)

var integrationGcpMutex = sync.Mutex{}

const (
	defaultType                    = "service_account"
	defaultAuthURI                 = "https://accounts.google.com/o/oauth2/auth"
	defaultTokenURI                = "https://oauth2.googleapis.com/token"
	defaultAuthProviderX509CertURL = "https://www.googleapis.com/oauth2/v1/certs"
	defaultClientX509CertURLPrefix = "https://www.googleapis.com/robot/v1/metadata/x509/"
)

type integrationGcpResource struct {
	Api  *datadogV1.GCPIntegrationApi
	Auth context.Context
}

type integrationGcpModel struct {
	ID                             types.String `tfsdk:"id"`
	ProjectID                      types.String `tfsdk:"project_id"`
	PrivateKeyId                   types.String `tfsdk:"private_key_id"`
	PrivateKey                     types.String `tfsdk:"private_key"`
	ClientEmail                    types.String `tfsdk:"client_email"`
	ClientId                       types.String `tfsdk:"client_id"`
	Automute                       types.Bool   `tfsdk:"automute"`
	HostFilters                    types.String `tfsdk:"host_filters"`
	ResourceCollectionEnabled      types.Bool   `tfsdk:"resource_collection_enabled"`
	CspmResourceCollectionEnabled  types.Bool   `tfsdk:"cspm_resource_collection_enabled"`
	IsSecurityCommandCenterEnabled types.Bool   `tfsdk:"is_security_command_center_enabled"`
}

func NewIntegrationGcpResource() resource.Resource {
	return &integrationGcpResource{}
}

func (r *integrationGcpResource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	providerData, _ := request.ProviderData.(*FrameworkProvider)
	r.Api = providerData.DatadogApiInstances.GetGCPIntegrationApiV1()
	r.Auth = providerData.Auth
}

func (r *integrationGcpResource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = "integration_gcp"
}

func (r *integrationGcpResource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description: "Provides a Datadog Integration GCP Sts resource. This can be used to create and manage Datadog - Google Cloud Platform integration.",
		Attributes: map[string]schema.Attribute{
			"project_id": schema.StringAttribute{
				Description: "Your Google Cloud project ID found in your JSON service account key.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"private_key_id": schema.StringAttribute{
				Description: "Your private key ID found in your JSON service account key.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"private_key": schema.StringAttribute{
				Description: "Your private key name found in your JSON service account key.",
				Required:    true,
				Sensitive:   true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"client_email": schema.StringAttribute{
				Description: "Your email found in your JSON service account key.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"client_id": schema.StringAttribute{
				Description: "Your ID found in your JSON service account key.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"host_filters": schema.StringAttribute{
				Description: "Limit the GCE instances that are pulled into Datadog by using tags. Only hosts that match one of the defined tags are imported into Datadog.",
				Optional:    true,
			},
			"automute": schema.BoolAttribute{
				Description: "Silence monitors for expected GCE instance shutdowns.",
				Optional:    true,
				Computed:    true,
			},
			"resource_collection_enabled": schema.BoolAttribute{
				Description: "When enabled, Datadog scans for all resources in your GCP environment.",
				Optional:    true,
				Computed:    true,
			},
			"cspm_resource_collection_enabled": schema.BoolAttribute{
				Description: "Whether Datadog collects cloud security posture management resources from your GCP project. If enabled, requires resource_collection_enabled to also be enabled.",
				Optional:    true,
				Computed:    true,
			},
			"is_security_command_center_enabled": schema.BoolAttribute{
				Description: "When enabled, Datadog will attempt to collect Security Command Center Findings. Note: This requires additional permissions on the service account.",
				Optional:    true,
				Computed:    true,
			},
			"id": utils.ResourceIDAttribute(),
		},
	}
}

func (r *integrationGcpResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, frameworkPath.Root("id"), request, response)
}

func (r *integrationGcpResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var state integrationGcpModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	integration, diags := r.getGCPIntegration(state)
	if diags.HasError() {
		response.Diagnostics.Append(diags...)
		return
	}

	if integration == nil {
		response.State.RemoveResource(ctx)
		return
	}

	// Save data into Terraform state
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (r *integrationGcpResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var state integrationGcpModel

	integrationGcpMutex.Lock()
	defer integrationGcpMutex.Unlock()

	response.Diagnostics.Append(request.Plan.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	body, diags := r.buildIntegrationGcpRequestBodyBase(state)
	r.addDefaultsToBody(body)
	r.addRequiredFieldsToBody(body, state)
	r.addOptionalFieldsToBody(body, state)

	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	_, _, err := r.Api.CreateGCPIntegration(r.Auth, *body)
	if err != nil {
		response.Diagnostics.Append(utils.FrameworkErrorDiag(err, "error creating GCP integration"))
		return
	}
	integration, diags := r.getGCPIntegration(state)
	if diags.HasError() {
		response.Diagnostics.Append(diags...)
		return
	}
	if integration == nil {
		response.Diagnostics.AddError("error retrieving GCP integration", "")
		return
	}
	r.updateState(ctx, &state, integration)

	// Save data into Terraform state
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (r *integrationGcpResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var state integrationGcpModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	body, diags := r.buildIntegrationGcpRequestBodyBase(state)
	r.addOptionalFieldsToBody(body, state)

	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	_, _, err := r.Api.UpdateGCPIntegration(r.Auth, *body)
	if err != nil {
		response.Diagnostics.Append(utils.FrameworkErrorDiag(err, "error updating GCP integration"))
		return
	}
	integration, diags := r.getGCPIntegration(state)
	if diags.HasError() {
		response.Diagnostics.Append(diags...)
		return
	}
	if integration == nil {
		response.Diagnostics.AddError("error retrieving GCP integration", "")
		return
	}
	r.updateState(ctx, &state, integration)

	// Save data into Terraform state
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (r *integrationGcpResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var state integrationGcpModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	body, diags := r.buildIntegrationGcpRequestBodyBase(state)

	response.Diagnostics.Append(diags...)

	_, httpResp, err := r.Api.DeleteGCPIntegration(r.Auth, *body)
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			return
		}
		response.Diagnostics.Append(utils.FrameworkErrorDiag(err, "error deleting integration_gcp_sts"))
		return
	}
}

func (r *integrationGcpResource) updateState(ctx context.Context, state *integrationGcpModel, resp *datadogV1.GCPAccount) {
	projectId := types.StringValue(resp.GetProjectId())
	// ProjectID and ClientEmail are the only parameters required in all mutating API requests
	state.ID = projectId
	state.ProjectID = projectId
	state.ClientEmail = types.StringValue(resp.GetClientEmail())

	if clientId, ok := resp.GetClientIdOk(); ok {
		state.ClientId = types.StringValue(*clientId)
	}
	if privateKey, ok := resp.GetPrivateKeyOk(); ok {
		state.PrivateKey = types.StringValue(*privateKey)
	}
	if privateKeyId, ok := resp.GetPrivateKeyIdOk(); ok {
		state.PrivateKeyId = types.StringValue(*privateKeyId)
	}
	if automute, ok := resp.GetAutomuteOk(); ok {
		state.Automute = types.BoolValue(*automute)
	}
	if hostFilters, ok := resp.GetHostFiltersOk(); ok && len(*hostFilters) > 0 {
		state.HostFilters = types.StringValue(*hostFilters)
	}
	if resourceCollectionEnabled, ok := resp.GetResourceCollectionEnabledOk(); ok {
		state.ResourceCollectionEnabled = types.BoolValue(*resourceCollectionEnabled)
	}
	if isCspmEnabled, ok := resp.GetIsCspmEnabledOk(); ok {
		state.CspmResourceCollectionEnabled = types.BoolValue(*isCspmEnabled)
	}
	if isSecurityCommandCenterEnabled, ok := resp.GetIsSecurityCommandCenterEnabledOk(); ok {
		state.IsSecurityCommandCenterEnabled = types.BoolValue(*isSecurityCommandCenterEnabled)
	}
}

func (r *integrationGcpResource) getGCPIntegration(state integrationGcpModel) (*datadogV1.GCPAccount, diag.Diagnostics) {
	diags := diag.Diagnostics{}
	resp, _, err := r.Api.ListGCPIntegration(r.Auth)
	if err != nil {
		diags.Append(utils.FrameworkErrorDiag(err, "error listing GCP integration"))
		return nil, diags
	}

	var account *datadogV1.GCPAccount
	for _, integration := range resp {
		if integration.GetProjectId() == state.ProjectID.ValueString() && integration.GetClientEmail() == state.ClientEmail.ValueString() {
			if err := utils.CheckForUnparsed(integration); err != nil {
				diags.AddError("response contains unparsedObject", err.Error())
				return nil, diags
			}
			account = &integration
			break
		}
	}

	return account, diags
}

func (r *integrationGcpResource) buildIntegrationGcpRequestBodyBase(state integrationGcpModel) (*datadogV1.GCPAccount, diag.Diagnostics) {
	diags := diag.Diagnostics{}
	body := datadogV1.NewGCPAccountWithDefaults()
	body.SetProjectId(state.ProjectID.ValueString())
	body.SetClientEmail(state.ClientEmail.ValueString())

	return body, diags
}

func (r *integrationGcpResource) addDefaultsToBody(body *datadogV1.GCPAccount) {
	body.SetType(defaultType)
	body.SetAuthUri(defaultAuthURI)
	body.SetAuthProviderX509CertUrl(defaultAuthProviderX509CertURL)
	body.SetClientX509CertUrl(defaultClientX509CertURLPrefix)
	body.SetTokenUri(defaultTokenURI)
}

func (r *integrationGcpResource) addRequiredFieldsToBody(body *datadogV1.GCPAccount, state integrationGcpModel) {
	body.SetClientId(state.ClientId.ValueString())
	body.SetPrivateKey(state.PrivateKey.ValueString())
	body.SetPrivateKeyId(state.PrivateKeyId.ValueString())
}

func (r *integrationGcpResource) addOptionalFieldsToBody(body *datadogV1.GCPAccount, state integrationGcpModel) {
	if !state.Automute.IsUnknown() && !state.Automute.IsNull() {
		body.SetAutomute(state.Automute.ValueBool())
	}
	if !state.CspmResourceCollectionEnabled.IsUnknown() && !state.CspmResourceCollectionEnabled.IsNull() {
		body.SetIsCspmEnabled(state.CspmResourceCollectionEnabled.ValueBool())
	}
	if !state.ResourceCollectionEnabled.IsUnknown() && !state.ResourceCollectionEnabled.IsNull() {
		body.SetResourceCollectionEnabled(state.ResourceCollectionEnabled.ValueBool())
	}
	if !state.IsSecurityCommandCenterEnabled.IsUnknown() && !state.IsSecurityCommandCenterEnabled.IsNull() {
		body.SetIsSecurityCommandCenterEnabled(state.IsSecurityCommandCenterEnabled.ValueBool())
	}
	if !state.HostFilters.IsUnknown() && !state.HostFilters.IsNull() {
		body.SetHostFilters(state.HostFilters.ValueString())
	}
}
