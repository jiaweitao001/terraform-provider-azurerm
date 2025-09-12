// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceconnector_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicelinker/2024-04-01/servicelinker"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

type ServiceConnectorKubernetesClusterResource struct{}

func (r ServiceConnectorKubernetesClusterResource) Exists(ctx context.Context, client *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	id, err := servicelinker.ParseScopedLinkerID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := client.ServiceConnector.ServiceLinkerClient.LinkerGet(ctx, *id)
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return utils.Bool(false), nil
		}
		return nil, fmt.Errorf("retrieving %s: %+v", *id, err)
	}
	return utils.Bool(true), nil
}

func TestAccServiceConnectorKubernetesClusterCosmosdb_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_kubernetes_cluster_connection", "test")
	r := ServiceConnectorKubernetesClusterResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.cosmosdbBasic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccServiceConnectorKubernetesClusterCosmosdb_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_kubernetes_cluster_connection", "test")
	r := ServiceConnectorKubernetesClusterResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.cosmosdbBasic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport),
	})
}

func TestAccServiceConnectorKubernetesClusterCosmosdb_secretAuth(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_kubernetes_cluster_connection", "test")
	r := ServiceConnectorKubernetesClusterResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.cosmosdbWithSecretAuth(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("authentication.0.secret"),
	})
}

func TestAccServiceConnectorKubernetesClusterStorageBlob_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_kubernetes_cluster_connection", "test")
	r := ServiceConnectorKubernetesClusterResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.storageBlob(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccServiceConnectorKubernetesClusterCosmosdb_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_kubernetes_cluster_connection", "test")
	r := ServiceConnectorKubernetesClusterResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.cosmosdbBasic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.cosmosdbUpdate(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (r ServiceConnectorKubernetesClusterResource) requiresImport(data acceptance.TestData) string {
	config := r.cosmosdbBasic(data)
	return fmt.Sprintf(`
%s

resource "azurerm_kubernetes_cluster_connection" "import" {
  name                    = azurerm_kubernetes_cluster_connection.test.name
  kubernetes_cluster_id   = azurerm_kubernetes_cluster_connection.test.kubernetes_cluster_id
  target_resource_id      = azurerm_kubernetes_cluster_connection.test.target_resource_id
  authentication {
    type            = "userAssignedIdentity"
    subscription_id = data.azurerm_subscription.test.subscription_id
    client_id       = azurerm_user_assigned_identity.test.client_id
  }
}
`, config)
}

func (r ServiceConnectorKubernetesClusterResource) storageBlob(data acceptance.TestData) string {
	template := r.template(data)
	return fmt.Sprintf(`
%[1]s

data "azurerm_subscription" "test" {}

resource "azurerm_user_assigned_identity" "test" {
  name                = "acctest%[2]s"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
}

resource "azurerm_kubernetes_cluster_connection" "test" {
  name                    = "acctestserviceconnector%[3]d"
  kubernetes_cluster_id   = azurerm_kubernetes_cluster.test.id
  target_resource_id      = azurerm_storage_account.test.id
  authentication {
    type            = "userAssignedIdentity"
    subscription_id = data.azurerm_subscription.test.subscription_id
    client_id       = azurerm_user_assigned_identity.test.client_id
  }
}
`, template, data.RandomString, data.RandomInteger)
}

func (r ServiceConnectorKubernetesClusterResource) cosmosdbBasic(data acceptance.TestData) string {
	template := r.template(data)
	return fmt.Sprintf(`
%[1]s

data "azurerm_subscription" "test" {}

resource "azurerm_user_assigned_identity" "test" {
  name                = "acctest%[2]s"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
}

resource "azurerm_kubernetes_cluster_connection" "test" {
  name                    = "acctestserviceconnector%[3]d"
  kubernetes_cluster_id   = azurerm_kubernetes_cluster.test.id
  target_resource_id      = azurerm_cosmosdb_sql_database.test.id
  authentication {
    type            = "userAssignedIdentity"
    subscription_id = data.azurerm_subscription.test.subscription_id
    client_id       = azurerm_user_assigned_identity.test.client_id
  }
}
`, template, data.RandomString, data.RandomInteger)
}

func (r ServiceConnectorKubernetesClusterResource) cosmosdbWithSecretAuth(data acceptance.TestData) string {
	template := r.template(data)
	return fmt.Sprintf(`
%[1]s

resource "azurerm_kubernetes_cluster_connection" "test" {
  name                    = "acctestserviceconnector%[2]d"
  kubernetes_cluster_id   = azurerm_kubernetes_cluster.test.id
  target_resource_id      = azurerm_cosmosdb_sql_database.test.id
  authentication {
    type   = "secret"
    name   = "foo"
    secret = "bar"
  }
}
`, template, data.RandomInteger)
}

func (r ServiceConnectorKubernetesClusterResource) cosmosdbUpdate(data acceptance.TestData) string {
	template := r.template(data)
	return fmt.Sprintf(`
%[1]s

data "azurerm_subscription" "test" {}

resource "azurerm_user_assigned_identity" "test" {
  name                = "acctest%[2]s"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
}

resource "azurerm_kubernetes_cluster_connection" "test" {
  name                    = "acctestserviceconnector%[3]d"
  kubernetes_cluster_id   = azurerm_kubernetes_cluster.test.id
  target_resource_id      = azurerm_cosmosdb_sql_database.test.id
  client_type             = "dotnet"
  authentication {
    type            = "userAssignedIdentity"
    subscription_id = data.azurerm_subscription.test.subscription_id
    client_id       = azurerm_user_assigned_identity.test.client_id
  }
}
`, template, data.RandomString, data.RandomInteger)
}

func (r ServiceConnectorKubernetesClusterResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%[2]d"
  location = "%[1]s"
}

resource "azurerm_user_assigned_identity" "test_aks" {
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  name                = "test-aks-identity"
}

resource "azurerm_kubernetes_cluster" "test" {
  name                = "acctestaks%[2]d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  dns_prefix          = "acctestaks%[2]d"

  default_node_pool {
    name       = "default"
    node_count = 1
    vm_size    = "Standard_A2_v2"
  }

  identity {
    type         = "UserAssigned"
    identity_ids = [azurerm_user_assigned_identity.test_aks.id]
  }
}

resource "azurerm_cosmosdb_account" "test" {
  name                = "acctest-ca-%[2]d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  offer_type          = "Standard"
  kind                = "GlobalDocumentDB"

  consistency_policy {
    consistency_level       = "BoundedStaleness"
    max_interval_in_seconds = 10
    max_staleness_prefix    = 200
  }

  geo_location {
    location          = azurerm_resource_group.test.location
    failover_priority = 0
  }
}

resource "azurerm_cosmosdb_sql_database" "test" {
  name                = "cosmos-sql-db"
  resource_group_name = azurerm_cosmosdb_account.test.resource_group_name
  account_name        = azurerm_cosmosdb_account.test.name
  throughput          = 400
}

resource "azurerm_storage_account" "test" {
  name                     = "acctestsa%[3]s"
  resource_group_name      = azurerm_resource_group.test.name
  location                 = azurerm_resource_group.test.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
}
`, data.Locations.Primary, data.RandomInteger, data.RandomString)
}
