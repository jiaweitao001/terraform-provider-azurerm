package client

import (
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2020-10-01-preview/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2020-10-01-preview/datastores"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2020-10-01-preview/inventoryitems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2020-10-01-preview/vcenters"
	"github.com/hashicorp/terraform-provider-azurerm/internal/common"
)

type Client struct {
	ClusterClient        *clusters.ClustersClient
	InventoryItemsClient *inventoryitems.InventoryItemsClient
	DataStoresClient     *datastores.DataStoresClient
	VcenterClient        *vcenters.VCentersClient
}

func NewClient(o *common.ClientOptions) *Client {
	clusterClient := clusters.NewClustersClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&clusterClient.Client, o.ResourceManagerAuthorizer)

	inventoryItemsClient := inventoryitems.NewInventoryItemsClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&inventoryItemsClient.Client, o.ResourceManagerAuthorizer)

	vcenterClient := vcenters.NewVCentersClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&vcenterClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		ClusterClient:        &clusterClient,
		InventoryItemsClient: &inventoryItemsClient,
		VcenterClient:        &vcenterClient,
	}
}
