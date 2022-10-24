package connectedvmware_test

import (
	"fmt"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
	"testing"
)

type ConnectedVmwareVcenterResource struct{}

func TestAccConnectedVmwareVcenter_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_connected_vmware_vcenter", "test")
	r := ConnectedVmwareVcenterResource{}

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

func (r ConnectedVmwareVcenterResource) basic(data acceptance.TestData) string {
	template := r.template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_connected_vmware_vcenter" "test" {
  name = "acctestvcenter-%d"
  resource_group_name = azurerm_resource_group.test.name
  location = azurerm_resource_group.test.location
  fqdn = "ContosoVMware.contoso.com"
}
`, template, data.RandomInteger)
}

func (r ConnectedVmwareVcenterResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name = "acctestRG-%[1]d"
  location = "%[2]s"
}

`, data.RandomInteger, data.Locations.Primary, data.RandomString)
}
