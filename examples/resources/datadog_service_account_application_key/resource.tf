
# Create new service_account_application_key resource


resource "datadog_service_account_application_key" "foo" {
    service_account_id = "00000000-0000-1234-0000-000000000000"
    name = "Application Key for managing dashboards"
    scopes = ["dashboards_read", "dashboards_write", "dashboards_public_share"]
}