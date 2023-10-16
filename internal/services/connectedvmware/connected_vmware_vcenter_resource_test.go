package connectedvmware_test

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/vcenters"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
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

func (r ConnectedVmwareVcenterResource) Exists(ctx context.Context, client *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	id, err := vcenters.ParseVCenterID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := client.ConnectedVmware.VcenterClient.Get(ctx, *id)
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return utils.Bool(false), nil
		}
		return nil, fmt.Errorf("retrieving %s: %+v", *id, err)
	}

	return utils.Bool(true), nil
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
  extended_location {
	name = "/subscriptions/a5215e1c-837f-4133-8531-85cd161d0cfb/resourceGroups/demoRG/providers/Microsoft.ExtendedLocation/customLocations/contoso"
    type = "testlocationtype"
  }
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
