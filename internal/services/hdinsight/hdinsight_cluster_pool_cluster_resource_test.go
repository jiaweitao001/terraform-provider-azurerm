package hdinsight_test

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hdinsight/2024-05-01/hdinsights"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
	"os"
	"testing"
)

type ClusterPoolClusterResource struct{}

func (r ClusterPoolClusterResource) Exists(ctx context.Context, client *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	id, err := hdinsights.ParseClusterID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := client.HDInsight2024.Hdinsights.ClustersGet(ctx, *id)
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("retrieving %s: %+v", *id, err)
	}
	return utils.Bool(true), nil
}

func TestAccHDInsightClusterPoolCluster_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_hdinsight_cluster_pool_cluster", "test")
	r := ClusterPoolClusterResource{}

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

func (r ClusterPoolClusterResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

provider "azurerm" {
  features {}
}

resource "azurerm_hdinsight_cluster_pools" "test" {
  name = "acctestpool-%[2]d"
  resource_group_name = azurerm_resource_group.test.name
  location = azurerm_resource_group.test.location

  compute_profile {
	vm_size = "Standard_D3_v2"
  }
}

resource "azurerm_hdinsight_cluster_pool_cluster" "test" {
  name = "acctestcluster-%[2]d"
  cluster_pool_name = azurerm_hdinsight_cluster_pools.test.name
  resource_group_name = azurerm_hdinsight_cluster_pools.test.resource_group_name
  location = azurerm_hdinsight_cluster_pools.test.location
  cluster_type = "Kafka"
  cluster_profile {
	cluster_version = "1.2.0"
	oss_version = "2.3.0"
	authorization_profile {
	  user_ids = ["%[3]s"]
	}
  }
  compute_profile {
	node {
	  vm_size = "Standard_D16a_v4"
	  type = "cluster"
	  count = 2
	}
  }
}
`, r.template(data), data.RandomInteger, os.Getenv("ARM_CLIENT_ID"))
}

func (r ClusterPoolClusterResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name = "acctest-hdinsightcluster%[1]d"
  location = "%[2]s"
}

resource "azurerm_virtual_network" "test" {
  name                = "clustervn%[1]d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  address_space = ["10.0.0.0/16"]
}

resource "azurerm_subnet" "test" {
  name                 = "clustersn%[1]d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefixes = ["10.0.2.0/24"]
}

resource "azurerm_user_assigned_identity" "test" {
  name                = "clusteridentity%[1]d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
}
`, data.RandomInteger, data.Locations.Primary)
}
