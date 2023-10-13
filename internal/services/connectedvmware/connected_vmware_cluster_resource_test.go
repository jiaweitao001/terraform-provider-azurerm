package connectedvmware_test

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/clusters"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
	"testing"
)

type ConnectedVmwareClusterResource struct{}

func TestAccConnectedVmwareCluster_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_connected_vmware_cluster", "test")
	r := ConnectedVmwareClusterResource{}

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

func (r ConnectedVmwareClusterResource) Exists(ctx context.Context, client *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	id, err := clusters.ParseClusterID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := client.ConnectedVmware.ClusterClient.Get(ctx, *id)
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return utils.Bool(false), nil
		}
		return nil, fmt.Errorf("retrieving %s: %+v", *id, err)
	}

	return utils.Bool(true), nil
}

func (r ConnectedVmwareClusterResource) basic(data acceptance.TestData) string {
	template := r.template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_connected_vmware_cluster" "test" {
  name = "acctestcluster%d"
  resource_group_name = azurerm_resource_group.test.name
  location = azurerm_resource_group.test.location
  mo_ref_id = "aaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"
  vcenter_id = azurerm_connected_vmware_vcenter.test.id
}
`, template, data.RandomInteger)
}

func (r ConnectedVmwareClusterResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_connected_vmware_vcenter" "test" {
  name = "acctestvcenter-d"
  resource_group_name = azurerm_resource_group.test.name
  location = azurerm_resource_group.test.location
  fqdn = "ContosoVMware.contoso.com"
}
`, data.RandomInteger, data.Locations.Primary)
}
