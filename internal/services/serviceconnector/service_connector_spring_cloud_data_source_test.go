package serviceconnector_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
)

type SpringCloudConnectorDataSource struct{}

func TestAccDataSourceServiceConnectorSpringCloud_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_spring_cloud_connection", "test")
	d := SpringCloudConnectorDataSource{}

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

func (SpringCloudConnectorDataSource) complete(data acceptance.TestData) string {
	config := ServiceConnectorSpringCloudResource{}.complete(data)
	return fmt.Sprintf(`
%s

data "azurerm_spring_cloud_connection" "test" {
  name            = azurerm_spring_cloud_connection.test.name
  spring_cloud_id = azurerm_spring_cloud_java_deployment.test.id
}
`, config)
}
