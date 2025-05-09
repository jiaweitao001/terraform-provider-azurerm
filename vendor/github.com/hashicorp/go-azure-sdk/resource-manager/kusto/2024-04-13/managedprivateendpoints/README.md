
## `github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2024-04-13/managedprivateendpoints` Documentation

The `managedprivateendpoints` SDK allows for interaction with Azure Resource Manager `kusto` (API Version `2024-04-13`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2024-04-13/managedprivateendpoints"
```


### Client Initialization

```go
client := managedprivateendpoints.NewManagedPrivateEndpointsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedPrivateEndpointsClient.CheckNameAvailability`

```go
ctx := context.TODO()
id := commonids.NewKustoClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

payload := managedprivateendpoints.ManagedPrivateEndpointsCheckNameRequest{
	// ...
}


read, err := client.CheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedPrivateEndpointsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := managedprivateendpoints.NewManagedPrivateEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "managedPrivateEndpointName")

payload := managedprivateendpoints.ManagedPrivateEndpoint{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ManagedPrivateEndpointsClient.Delete`

```go
ctx := context.TODO()
id := managedprivateendpoints.NewManagedPrivateEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "managedPrivateEndpointName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ManagedPrivateEndpointsClient.Get`

```go
ctx := context.TODO()
id := managedprivateendpoints.NewManagedPrivateEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "managedPrivateEndpointName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedPrivateEndpointsClient.List`

```go
ctx := context.TODO()
id := commonids.NewKustoClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedPrivateEndpointsClient.Update`

```go
ctx := context.TODO()
id := managedprivateendpoints.NewManagedPrivateEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "managedPrivateEndpointName")

payload := managedprivateendpoints.ManagedPrivateEndpoint{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
