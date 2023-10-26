package devcenter_test

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devcenter/2023-04-01/devboxdefinitions"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
	"testing"
)

type DevCenterDevBoxDefinitionsResource struct{}

func (r DevCenterDevBoxDefinitionsResource) Exists(ctx context.Context, client *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	id, err := devboxdefinitions.ParseDevCenterDevBoxDefinitionID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := client.DevCenter.V20230401.DevBoxDefinitions.Get(ctx, *id)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %+v", *id, err)
	}

	return utils.Bool(resp.Model != nil), nil
}

func TestAccDevCenterDevBoxDefinitions_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_dev_center_dev_box_definitions", "test")
	r := DevCenterDevBoxDefinitionsResource{}

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

func (r DevCenterDevBoxDefinitionsResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_dev_center_dev_box_definitions" "test" {
  name              = "acctestdevboxdef-%[2]d"
  location          = azurerm_resource_group.test.location
  dev_center_name   = azurerm_dev_center.test.name
  resource_group_name = azurerm_resource_group.test.name
  hibernate_support = "Enabled"
  os_storage_type   = "ssd_1024gb"
  sku {
	name = "general_a_8c32gb1024ssd_v2"
  }
}
`, r.template(data), data.RandomInteger)
}

func (r DevCenterDevBoxDefinitionsResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%[1]d"
  location = "%[2]s"
}

resource "azurerm_dev_center" "test" {
  location = azurerm_resource_group.test.location
  name     = "acctestdc-%[1]d"
  resource_group_name = azurerm_resource_group.test.name
}
`, data.RandomInteger, data.Locations.Primary)
}
