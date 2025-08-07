---
subcategory: "App Service (Web Apps)"
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_app_service_connection"
description: |-
  Gets information about an existing App Service Connection.
---

# Data Source: azurerm_app_service_connection

Use this data source to access information about an existing App Service Connection.

## Example Usage

```hcl
data "azurerm_app_service_connection" "example" {
  name           = "example-serviceconnector"
  app_service_id = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-resources/providers/Microsoft.Web/sites/example-app"
}

output "target_resource_id" {
  value = data.azurerm_app_service_connection.example.target_resource_id
}

output "authentication_type" {
  value = data.azurerm_app_service_connection.example.authentication[0].type
}
```

## Arguments Reference

The following arguments are supported:

* `name` - (Required) The name of the App Service Connection.

* `app_service_id` - (Required) The ID of the App Service.

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported:

* `target_resource_id` - The ID of the target resource.

* `client_type` - The application client type.

* `vnet_solution` - The type of VNet solution.

* `authentication` - The authentication info. An `authentication` block as defined below.

* `secret_store` - An option to store secret value in secure place. A `secret_store` block as defined below.

---

An `authentication` block exports the following:

* `type` - The authentication type.

* `name` - Username or account name for secret auth.

* `client_id` - Client ID for userAssignedIdentity or servicePrincipal auth.

* `subscription_id` - Subscription ID for userAssignedIdentity auth.

* `principal_id` - Principal ID for servicePrincipal auth.

~> **Note:** For security reasons, sensitive values such as `secret` and `certificate` are not exposed in the data source.

---

A `secret_store` block exports the following:

* `key_vault_id` - The key vault id to store secret.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `read` - (Defaults to 5 minutes) Used when retrieving the App Service Connection.
