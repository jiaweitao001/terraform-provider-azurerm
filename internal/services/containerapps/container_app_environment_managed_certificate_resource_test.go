// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package containerapps_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/managedenvironments"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

type ContainerAppEnvironmentManagedCertificateResource struct{}

func (r ContainerAppEnvironmentManagedCertificateResource) Exists(ctx context.Context, client *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	id, err := managedenvironments.ParseManagedCertificateID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := client.ContainerApps.ManagedEnvironmentClient.ManagedCertificatesGet(ctx, *id)
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("retrieving %s: %+v", *id, err)
	}

	return pointer.To(resp.Model != nil), nil
}

func TestAccContainerAppEnvironmentManagedCertificate_basic(t *testing.T) {
	if os.Getenv("ARM_TEST_DNS_ZONE") == "" || os.Getenv("ARM_TEST_DATA_RESOURCE_GROUP") == "" {
		t.Skipf("Skipping as either ARM_TEST_DNS_ZONE or ARM_TEST_DATA_RESOURCE_GROUP is not set")
	}

	data := acceptance.BuildTestData(t, "azurerm_container_app_environment_managed_certificate", "test")
	r := ContainerAppEnvironmentManagedCertificateResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccContainerAppEnvironmentManagedCertificate_requiresImport(t *testing.T) {
	if os.Getenv("ARM_TEST_DNS_ZONE") == "" || os.Getenv("ARM_TEST_DATA_RESOURCE_GROUP") == "" {
		t.Skipf("Skipping as either ARM_TEST_DNS_ZONE or ARM_TEST_DATA_RESOURCE_GROUP is not set")
	}

	data := acceptance.BuildTestData(t, "azurerm_container_app_environment_managed_certificate", "test")
	r := ContainerAppEnvironmentManagedCertificateResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport),
	})
}

func TestAccContainerAppEnvironmentManagedCertificate_complete(t *testing.T) {
	if os.Getenv("ARM_TEST_DNS_ZONE") == "" || os.Getenv("ARM_TEST_DATA_RESOURCE_GROUP") == "" {
		t.Skipf("Skipping as either ARM_TEST_DNS_ZONE or ARM_TEST_DATA_RESOURCE_GROUP is not set")
	}

	data := acceptance.BuildTestData(t, "azurerm_container_app_environment_managed_certificate", "test")
	r := ContainerAppEnvironmentManagedCertificateResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccContainerAppEnvironmentManagedCertificate_update(t *testing.T) {
	if os.Getenv("ARM_TEST_DNS_ZONE") == "" || os.Getenv("ARM_TEST_DATA_RESOURCE_GROUP") == "" {
		t.Skipf("Skipping as either ARM_TEST_DNS_ZONE or ARM_TEST_DATA_RESOURCE_GROUP is not set")
	}

	data := acceptance.BuildTestData(t, "azurerm_container_app_environment_managed_certificate", "test")
	r := ContainerAppEnvironmentManagedCertificateResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (r ContainerAppEnvironmentManagedCertificateResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider azurerm {
  features {}
}

%s

resource "azurerm_container_app_environment_managed_certificate" "test" {
  name                         = "acctest-mcert%[2]d"
  container_app_environment_id = azurerm_container_app_environment.test.id
  subject_name                 = trimprefix(azurerm_dns_txt_record.test.fqdn, "asuid.")

  depends_on = [azurerm_container_app_custom_domain.test]
}
`, r.template(data), data.RandomInteger)
}

func (r ContainerAppEnvironmentManagedCertificateResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_container_app_environment_managed_certificate" "import" {
  name                         = azurerm_container_app_environment_managed_certificate.test.name
  container_app_environment_id = azurerm_container_app_environment_managed_certificate.test.container_app_environment_id
  subject_name                 = azurerm_container_app_environment_managed_certificate.test.subject_name
}
`, r.basic(data))
}

func (r ContainerAppEnvironmentManagedCertificateResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider azurerm {
  features {}
}

%s

resource "azurerm_container_app_environment_managed_certificate" "test" {
  name                         = "acctest-mcert%[2]d"
  container_app_environment_id = azurerm_container_app_environment.test.id
  subject_name                 = trimprefix(azurerm_dns_txt_record.test.fqdn, "asuid.")
  domain_control_validation    = "CNAME"

  tags = {
    env = "test"
  }

  depends_on = [azurerm_container_app_custom_domain.test]
}
`, r.template(data), data.RandomInteger)
}

func (r ContainerAppEnvironmentManagedCertificateResource) template(data acceptance.TestData) string {
	dnsZone := os.Getenv("ARM_TEST_DNS_ZONE")
	dataResourceGroup := os.Getenv("ARM_TEST_DATA_RESOURCE_GROUP")

	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-CAE-%[1]d"
  location = "%[2]s"
}

data "azurerm_dns_zone" "test" {
  name                = "%[3]s"
  resource_group_name = "%[4]s"
}

resource "azurerm_dns_txt_record" "test" {
  name                = "asuid.acctestcapp%[1]d"
  resource_group_name = data.azurerm_dns_zone.test.resource_group_name
  zone_name           = data.azurerm_dns_zone.test.name
  ttl                 = 300

  record {
    value = azurerm_container_app.test.custom_domain_verification_id
  }
}

resource "azurerm_dns_cname_record" "test" {
  name                = "acctestcapp%[1]d"
  resource_group_name = data.azurerm_dns_zone.test.resource_group_name
  zone_name           = data.azurerm_dns_zone.test.name
  ttl                 = 300
  record              = azurerm_container_app.test.ingress[0].fqdn
}

resource "azurerm_log_analytics_workspace" "test" {
  name                = "acctestCAEnv-%[1]d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  sku                 = "PerGB2018"
  retention_in_days   = 30
}

resource "azurerm_container_app_environment" "test" {
  name                       = "acctest-CAEnv%[1]d"
  resource_group_name        = azurerm_resource_group.test.name
  location                   = azurerm_resource_group.test.location
  log_analytics_workspace_id = azurerm_log_analytics_workspace.test.id
}

resource "azurerm_container_app" "test" {
  name                         = "acctest-capp-%[1]d"
  resource_group_name          = azurerm_resource_group.test.name
  container_app_environment_id = azurerm_container_app_environment.test.id
  revision_mode                = "Single"

  template {
    container {
      name   = "acctest-cont-%[1]d"
      image  = "jackofallops/azure-containerapps-python-acctest:v0.0.1"
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

resource "azurerm_container_app_custom_domain" "test" {
  name             = trimprefix(azurerm_dns_txt_record.test.fqdn, "asuid.")
  container_app_id = azurerm_container_app.test.id

  depends_on = [azurerm_dns_cname_record.test]

  lifecycle {
    ignore_changes = [certificate_binding_type, container_app_environment_certificate_id]
  }
}
`, data.RandomInteger, data.Locations.Primary, dnsZone, dataResourceGroup)
}
