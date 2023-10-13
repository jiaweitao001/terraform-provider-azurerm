package client

import (
	"fmt"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/datastores"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/hosts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/inventoryitems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/resourcepools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/vcenters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/virtualmachineinstances"
	"github.com/hashicorp/terraform-provider-azurerm/internal/common"
)

type Client struct {
	ClusterClient        *clusters.ClustersClient
	InventoryItemsClient *inventoryitems.InventoryItemsClient
	DataStoresClient     *datastores.DataStoresClient
	VcenterClient        *vcenters.VCentersClient
	HostClient           *hosts.HostsClient
	ResourcepoolClient   *resourcepools.ResourcePoolsClient
	VirtualMachineClient *virtualmachineinstances.VirtualMachineInstancesClient
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	clusterClient, err := clusters.NewClustersClientWithBaseURI(o.Environment.ResourceManager)
	if err != nil {
		return nil, fmt.Errorf("building Connected VMware Cluster client: %+v", err)
	}
	o.Configure(clusterClient.Client, o.Authorizers.ResourceManager)

	inventoryItemsClient, err := inventoryitems.NewInventoryItemsClientWithBaseURI(o.Environment.ResourceManager)
	if err != nil {
		return nil, fmt.Errorf("building Connected VMware Inventory Item client: %+v", err)
	}
	o.Configure(inventoryItemsClient.Client, o.Authorizers.ResourceManager)

	vcenterClient, err := vcenters.NewVCentersClientWithBaseURI(o.Environment.ResourceManager)
	if err != nil {
		return nil, fmt.Errorf("building Connected VMware VCenter client: %+v", err)
	}
	o.Configure(vcenterClient.Client, o.Authorizers.ResourceManager)

	hostClient, err := hosts.NewHostsClientWithBaseURI(o.Environment.ResourceManager)
	if err != nil {
		return nil, fmt.Errorf("building Connected VMware Host client: %+v", err)
	}
	o.Configure(hostClient.Client, o.Authorizers.ResourceManager)

	resourcepoolClient, err := resourcepools.NewResourcePoolsClientWithBaseURI(o.Environment.ResourceManager)
	if err != nil {
		return nil, fmt.Errorf("building Connected VMware Resource Pool client: %+v", err)
	}
	o.Configure(resourcepoolClient.Client, o.Authorizers.ResourceManager)

	virtualMachineClient, err := virtualmachineinstances.NewVirtualMachineInstancesClientWithBaseURI(o.Environment.ResourceManager)
	if err != nil {
		return nil, fmt.Errorf("building Connected VMware Virtual Machine client: %+v", err)
	}
	o.Configure(resourcepoolClient.Client, o.Authorizers.ResourceManager)

	return &Client{
		ClusterClient:        clusterClient,
		InventoryItemsClient: inventoryItemsClient,
		VcenterClient:        vcenterClient,
		HostClient:           hostClient,
		ResourcepoolClient:   resourcepoolClient,
		VirtualMachineClient: virtualMachineClient,
	}, nil
}
