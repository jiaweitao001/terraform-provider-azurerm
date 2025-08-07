// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceconnector_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
)

type AppServiceConnectorDataSource struct{}

func TestAccAppServiceConnectorDataSource_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_app_service_connection", "test")
	d := AppServiceConnectorDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: d.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("name").Exists(),
				check.That(data.ResourceName).Key("app_service_id").Exists(),
				check.That(data.ResourceName).Key("target_resource_id").Exists(),
				check.That(data.ResourceName).Key("authentication.0.type").HasValue("systemAssignedIdentity"),
				check.That(data.ResourceName).Key("id").Exists(),
			),
		},
	})
}

func TestAccAppServiceConnectorDataSource_secretAuth(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_app_service_connection", "test")
	d := AppServiceConnectorDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: d.secretAuth(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("name").Exists(),
				check.That(data.ResourceName).Key("app_service_id").Exists(),
				check.That(data.ResourceName).Key("target_resource_id").Exists(),
				check.That(data.ResourceName).Key("authentication.0.type").HasValue("secret"),
				check.That(data.ResourceName).Key("authentication.0.name").HasValue("foo"),
				check.That(data.ResourceName).Key("id").Exists(),
			),
		},
	})
}

func (AppServiceConnectorDataSource) basic(data acceptance.TestData) string {
	resource := ServiceConnectorAppServiceResource{}
	return fmt.Sprintf(`
%s

data "azurerm_app_service_connection" "test" {
  name           = azurerm_app_service_connection.test.name
  app_service_id = azurerm_app_service_connection.test.app_service_id
}
`, resource.cosmosdbBasic(data))
}

func (AppServiceConnectorDataSource) secretAuth(data acceptance.TestData) string {
	resource := ServiceConnectorAppServiceResource{}
	return fmt.Sprintf(`
%s

data "azurerm_app_service_connection" "test" {
  name           = azurerm_app_service_connection.test.name
  app_service_id = azurerm_app_service_connection.test.app_service_id
}
`, resource.cosmosdbWithSecretAuth(data))
}
