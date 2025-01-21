package dynatrace_test

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2023-04-27/singlesignon"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"testing"
)

type SingleSignOnResource struct {
	dynatraceInfo dynatraceInfo
}

func NewDynatraceSingleSignOnResource() SingleSignOnResource {
	return SingleSignOnResource{
		dynatraceInfo: dynatraceInfo{
			UserCountry:     "westus2",
			UserEmail:       "alice@microsoft.com",
			UserFirstName:   "Alice",
			UserLastName:    "Bobab",
			UserPhoneNumber: "12345678",
		},
	}
}

func (r SingleSignOnResource) preCheck(t *testing.T) {
	if r.dynatraceInfo.UserCountry == "" {
		t.Skipf("DYNATRACE_USER_COUNTRY must be set for acceptance tests")
	}
	if r.dynatraceInfo.UserEmail == "" {
		t.Skipf("DYNATRACE_USER_EMAIL must be set for acceptance tests")
	}
	if r.dynatraceInfo.UserFirstName == "" {
		t.Skipf("DYNATRACE_USER_FIRST_NAME must be set for acceptance tests")
	}
	if r.dynatraceInfo.UserLastName == "" {
		t.Skipf("DYNATRACE_USER_LAST_NAME must be set for acceptance tests")
	}
	if r.dynatraceInfo.UserPhoneNumber == "" {
		t.Skipf("DYNATRACE_USER_PHONE_NUMBER must be set for acceptance tests")
	}
}

func (d SingleSignOnResource) Exists(ctx context.Context, client *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	id, err := singlesignon.ParseSingleSignOnConfigurationID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := client.Dynatrace.SingleSignOnClient.Get(ctx, *id)
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("retrieving %s: %+v", *id, err)
	}
	return pointer.To(true), nil
}

func TestAccDynatraceSingleSignOnResource_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_dynatrace_single_sign_on", "test")
	r := NewDynatraceSingleSignOnResource()
	//r.preCheck(t)

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

func (r SingleSignOnResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azurerm_dynatrace_single_sign_on" "test" {
	  name                = "acctestsso%d"
	  resource_group_name = azurerm_resource_group.test.name
	  monitor_name        = azurerm_dynatrace_monitor.test.name
	  aad_domains         = ["example.com"]
	  enterprise_app_id   = "00000000-0000-0000-0000-000000000000"
	  single_sign_on_state = "Existing"
	  single_sign_on_url   = "https://example.com"
}
`, MonitorsResource{}.basic(data), data.RandomInteger)
}
