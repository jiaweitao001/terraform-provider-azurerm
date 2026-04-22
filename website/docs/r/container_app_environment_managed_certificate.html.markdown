---
subcategory: "Container Apps"
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_container_app_environment_managed_certificate"
description: |-
  Manages a Container App Environment Managed Certificate.
---

# azurerm_container_app_environment_managed_certificate

Manages a Container App Environment Managed Certificate.

~> **Note:** This resource requires a Custom Domain to be added to the Container App before the Managed Certificate can be created. The Custom Domain must have the appropriate DNS records (TXT and CNAME) configured before the Managed Certificate can be provisioned.

## Example Usage

```hcl
resource "azurerm_resource_group" "example" {
  name     = "example-resources"
  location = "West Europe"
}

resource "azurerm_log_analytics_workspace" "example" {
  name                = "example"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  sku                 = "PerGB2018"
  retention_in_days   = 30
}

resource "azurerm_container_app_environment" "example" {
  name                       = "Example-Environment"
  location                   = azurerm_resource_group.example.location
  resource_group_name        = azurerm_resource_group.example.name
  log_analytics_workspace_id = azurerm_log_analytics_workspace.example.id
}

resource "azurerm_dns_zone" "example" {
  name                = "contoso.com"
  resource_group_name = azurerm_resource_group.example.name
}

resource "azurerm_dns_txt_record" "example" {
  name                = "asuid.example"
  resource_group_name = azurerm_dns_zone.example.resource_group_name
  zone_name           = azurerm_dns_zone.example.name
  ttl                 = 300

  record {
    value = azurerm_container_app.example.custom_domain_verification_id
  }
}

resource "azurerm_dns_cname_record" "example" {
  name                = "example"
  zone_name           = azurerm_dns_zone.example.name
  resource_group_name = azurerm_resource_group.example.name
  ttl                 = 300
  record              = azurerm_container_app.example.ingress[0].fqdn
}

resource "azurerm_container_app" "example" {
  name                         = "example-app"
  container_app_environment_id = azurerm_container_app_environment.example.id
  resource_group_name          = azurerm_resource_group.example.name
  revision_mode                = "Single"

  template {
    container {
      name   = "examplecontainerapp"
      image  = "mcr.microsoft.com/k8se/quickstart:latest"
      cpu    = 0.25
      memory = "0.5Gi"
    }
  }
  ingress {
    allow_insecure_connections = false
    external_enabled           = true
    target_port                = 5000
    transport                  = "http"
    traffic_weight {
      latest_revision = true
      percentage      = 100
    }
  }
}

# Step 1: Add the custom domain without certificate binding
resource "azurerm_container_app_custom_domain" "example" {
  name             = "example.contoso.com"
  container_app_id = azurerm_container_app.example.id

  depends_on = [azurerm_dns_cname_record.example, azurerm_dns_txt_record.example]

  lifecycle {
    ignore_changes = [certificate_binding_type, container_app_environment_certificate_id, container_app_environment_managed_certificate_id]
  }
}

# Step 2: Create the managed certificate
resource "azurerm_container_app_environment_managed_certificate" "example" {
  name                         = "example-managed-cert"
  container_app_environment_id = azurerm_container_app_environment.example.id
  subject_name                 = "example.contoso.com"
  domain_control_validation    = "CNAME"

  depends_on = [azurerm_container_app_custom_domain.example]
}
```

## Arguments Reference

The following arguments are supported:

* `name` - (Required) The name of the Container App Environment Managed Certificate. Changing this forces a new resource to be created.

* `container_app_environment_id` - (Required) The ID of the Container App Environment in which to create the Managed Certificate. Changing this forces a new resource to be created.

* `subject_name` - (Required) The subject name of the certificate. This must be the domain name that the certificate is for. Changing this forces a new resource to be created.

* `domain_control_validation` - (Optional) The method used to validate domain ownership. Possible values are `CNAME`, `HTTP` and `TXT`. Defaults to `CNAME`. Changing this forces a new resource to be created.

* `tags` - (Optional) A mapping of tags to assign to the resource.

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported:

* `id` - The ID of the Container App Environment Managed Certificate.

* `provisioning_state` - The provisioning state of the Managed Certificate.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/configure#define-operation-timeouts) for certain actions:

* `create` - (Defaults to 30 minutes) Used when creating the Container App Environment Managed Certificate.
* `read` - (Defaults to 5 minutes) Used when retrieving the Container App Environment Managed Certificate.
* `update` - (Defaults to 30 minutes) Used when updating the Container App Environment Managed Certificate.
* `delete` - (Defaults to 30 minutes) Used when deleting the Container App Environment Managed Certificate.

## Import

A Container App Environment Managed Certificate can be imported using the `resource id`, e.g.

```shell
terraform import azurerm_container_app_environment_managed_certificate.example "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.App/managedEnvironments/myenv/managedCertificates/mycertificate"
```

## API Providers
<!-- This section is generated, changes will be overwritten -->
This resource uses the following Azure API Providers:

* `Microsoft.App` - 2025-07-01
