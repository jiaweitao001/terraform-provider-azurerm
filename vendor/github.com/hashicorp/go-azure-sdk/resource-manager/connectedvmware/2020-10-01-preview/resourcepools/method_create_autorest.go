package resourcepools

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

type CreateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Create ...
func (c ResourcePoolsClient) Create(ctx context.Context, id ResourcePoolId, input ResourcePool) (result CreateOperationResponse, err error) {
	req, err := c.preparerForCreate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcepools.ResourcePoolsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcepools.ResourcePoolsClient", "Create", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateThenPoll performs Create then polls until it's completed
func (c ResourcePoolsClient) CreateThenPoll(ctx context.Context, id ResourcePoolId, input ResourcePool) error {
	result, err := c.Create(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing Create: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Create: %+v", err)
	}

	return nil
}

// preparerForCreate prepares the Create request.
func (c ResourcePoolsClient) preparerForCreate(ctx context.Context, id ResourcePoolId, input ResourcePool) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCreate sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (c ResourcePoolsClient) senderForCreate(ctx context.Context, req *http.Request) (future CreateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}
	future.Poller, err = polling.NewLongRunningPollerFromResponse(ctx, resp, c.Client)
	return
}
