package serviceconnector_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
)

type AppServiceConnectorDataSource struct{}

func TestAccDataSourceServiceConnectorAppService_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_app_service_connection", "test")
	d := AppServiceConnectorDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: d.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("target_resource_id").Exists(),
				check.That(data.ResourceName).Key("client_type").Exists(),
				check.That(data.ResourceName).Key("auth_info.0.type").Exists(),
				check.That(data.ResourceName).Key("vnet_solution").Exists(),
			),
		},
	})
}

func (AppServiceConnectorDataSource) complete(data acceptance.TestData) string {
	config := ServiceConnectorAppServiceResource{}.complete(data)
	return fmt.Sprintf(`
%s

data "azurerm_app_service_connection" "test" {
  name           = azurerm_app_service_connection.test.name
  app_service_id = azurerm_linux_web_app.test.id
}
`, config)
}
