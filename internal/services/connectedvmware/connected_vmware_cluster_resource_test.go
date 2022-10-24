package connectedvmware_test

import (
	"fmt"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"testing"
)

type ConnectedVmwareClusterResource struct{}

func TestAccConnectedVmwareCluster_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_connected_vmware_cluster", "test")
	r := ConnectedVmwareClusterResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
		},
	})
}

func (r ConnectedVmwareClusterResource) basic(data acceptance.TestData) string {
	template := r.template(data)

}

func (r ConnectedVmwareClusterResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name = "acctestRG-%[1]d"
  location = "%[2]s"
}

resource "azurerm_connected_vmware_clusters" "test" {
  name = "acctestcluster%[3]s"
  resource_group_name = azurerm_resource_group.test.name
  location = azurerm_resource_group.test.location
  mo_ref_id = "aaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"
  vcenter_id = 
}
`)
}
