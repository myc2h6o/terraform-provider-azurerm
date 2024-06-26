package iotcentral

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
)

// UsersClient is the azure IoT Central is a service that makes it easy to connect, monitor, and manage your IoT
// devices at scale.
type UsersClient struct {
	BaseClient
}

// NewUsersClient creates an instance of the UsersClient client.
func NewUsersClient(subdomain string) UsersClient {
	return UsersClient{New(subdomain)}
}

// Create sends the create request.
// Parameters:
// userID - unique ID for the user.
// body - user body.
func (client UsersClient) Create(ctx context.Context, userID string, body BasicUser) (result UserModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: userID,
			Constraints: []validation.Constraint{{Target: "userID", Name: validation.MaxLength, Rule: 48, Chain: nil},
				{Target: "userID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-_]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.UsersClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, userID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.UsersClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.UsersClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.UsersClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client UsersClient) CreatePreparer(ctx context.Context, userID string, body BasicUser) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"userId": autorest.Encode("path", userID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/users/{userId}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client UsersClient) CreateResponder(resp *http.Response) (result UserModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get sends the get request.
// Parameters:
// userID - unique ID for the user.
func (client UsersClient) Get(ctx context.Context, userID string) (result UserModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: userID,
			Constraints: []validation.Constraint{{Target: "userID", Name: validation.MaxLength, Rule: 48, Chain: nil},
				{Target: "userID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-_]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.UsersClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, userID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.UsersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.UsersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.UsersClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client UsersClient) GetPreparer(ctx context.Context, userID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"userId": autorest.Encode("path", userID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/users/{userId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client UsersClient) GetResponder(resp *http.Response) (result UserModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List sends the list request.
// Parameters:
// maxpagesize - the maximum number of resources to return from one response.
func (client UsersClient) List(ctx context.Context, maxpagesize *int32) (result UserCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.List")
		defer func() {
			sc := -1
			if result.uc.Response.Response != nil {
				sc = result.uc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: maxpagesize,
			Constraints: []validation.Constraint{{Target: "maxpagesize", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "maxpagesize", Name: validation.InclusiveMaximum, Rule: int64(100), Chain: nil},
					{Target: "maxpagesize", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("iotcentral.UsersClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, maxpagesize)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.UsersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.uc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.UsersClient", "List", resp, "Failure sending request")
		return
	}

	result.uc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.UsersClient", "List", resp, "Failure responding to request")
		return
	}
	if result.uc.hasNextLink() && result.uc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client UsersClient) ListPreparer(ctx context.Context, maxpagesize *int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if maxpagesize != nil {
		queryParameters["maxpagesize"] = autorest.Encode("query", *maxpagesize)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPath("/users"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client UsersClient) ListResponder(resp *http.Response) (result UserCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client UsersClient) listNextResults(ctx context.Context, lastResults UserCollection) (result UserCollection, err error) {
	req, err := lastResults.userCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "iotcentral.UsersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "iotcentral.UsersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.UsersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client UsersClient) ListComplete(ctx context.Context, maxpagesize *int32) (result UserCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, maxpagesize)
	return
}

// Remove sends the remove request.
// Parameters:
// userID - unique ID for the user.
func (client UsersClient) Remove(ctx context.Context, userID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.Remove")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: userID,
			Constraints: []validation.Constraint{{Target: "userID", Name: validation.MaxLength, Rule: 48, Chain: nil},
				{Target: "userID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-_]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.UsersClient", "Remove", err.Error())
	}

	req, err := client.RemovePreparer(ctx, userID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.UsersClient", "Remove", nil, "Failure preparing request")
		return
	}

	resp, err := client.RemoveSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "iotcentral.UsersClient", "Remove", resp, "Failure sending request")
		return
	}

	result, err = client.RemoveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.UsersClient", "Remove", resp, "Failure responding to request")
		return
	}

	return
}

// RemovePreparer prepares the Remove request.
func (client UsersClient) RemovePreparer(ctx context.Context, userID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"userId": autorest.Encode("path", userID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/users/{userId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RemoveSender sends the Remove request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) RemoveSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RemoveResponder handles the response to the Remove request. The method always
// closes the http.Response Body.
func (client UsersClient) RemoveResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update sends the update request.
// Parameters:
// userID - unique ID for the user.
// body - user patch body.
// ifMatch - only perform the operation if the entity's etag matches one of the etags provided or * is
// provided.
func (client UsersClient) Update(ctx context.Context, userID string, body interface{}, ifMatch string) (result UserModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: userID,
			Constraints: []validation.Constraint{{Target: "userID", Name: validation.MaxLength, Rule: 48, Chain: nil},
				{Target: "userID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-_]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.UsersClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, userID, body, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.UsersClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.UsersClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.UsersClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client UsersClient) UpdatePreparer(ctx context.Context, userID string, body interface{}, ifMatch string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"userId": autorest.Encode("path", userID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/users/{userId}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client UsersClient) UpdateResponder(resp *http.Response) (result UserModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
