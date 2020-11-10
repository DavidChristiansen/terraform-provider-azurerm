package security

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// OnPremiseIotSensorsClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type OnPremiseIotSensorsClient struct {
	BaseClient
}

// NewOnPremiseIotSensorsClient creates an instance of the OnPremiseIotSensorsClient client.
func NewOnPremiseIotSensorsClient(subscriptionID string, ascLocation string) OnPremiseIotSensorsClient {
	return NewOnPremiseIotSensorsClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewOnPremiseIotSensorsClientWithBaseURI creates an instance of the OnPremiseIotSensorsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewOnPremiseIotSensorsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) OnPremiseIotSensorsClient {
	return OnPremiseIotSensorsClient{NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
}

// CreateOrUpdate create or update on-premise IoT sensor
// Parameters:
// onPremiseIotSensorName - name of the on-premise IoT sensor
func (client OnPremiseIotSensorsClient) CreateOrUpdate(ctx context.Context, onPremiseIotSensorName string) (result OnPremiseIotSensor, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OnPremiseIotSensorsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.OnPremiseIotSensorsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, onPremiseIotSensorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.OnPremiseIotSensorsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.OnPremiseIotSensorsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.OnPremiseIotSensorsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client OnPremiseIotSensorsClient) CreateOrUpdatePreparer(ctx context.Context, onPremiseIotSensorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"onPremiseIotSensorName": autorest.Encode("path", onPremiseIotSensorName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-06-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/onPremiseIotSensors/{onPremiseIotSensorName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client OnPremiseIotSensorsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client OnPremiseIotSensorsClient) CreateOrUpdateResponder(resp *http.Response) (result OnPremiseIotSensor, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete on-premise IoT sensor
// Parameters:
// onPremiseIotSensorName - name of the on-premise IoT sensor
func (client OnPremiseIotSensorsClient) Delete(ctx context.Context, onPremiseIotSensorName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OnPremiseIotSensorsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.OnPremiseIotSensorsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, onPremiseIotSensorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.OnPremiseIotSensorsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "security.OnPremiseIotSensorsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.OnPremiseIotSensorsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client OnPremiseIotSensorsClient) DeletePreparer(ctx context.Context, onPremiseIotSensorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"onPremiseIotSensorName": autorest.Encode("path", onPremiseIotSensorName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-06-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/onPremiseIotSensors/{onPremiseIotSensorName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client OnPremiseIotSensorsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client OnPremiseIotSensorsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DownloadActivation download sensor activation file
// Parameters:
// onPremiseIotSensorName - name of the on-premise IoT sensor
func (client OnPremiseIotSensorsClient) DownloadActivation(ctx context.Context, onPremiseIotSensorName string) (result ReadCloser, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OnPremiseIotSensorsClient.DownloadActivation")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.OnPremiseIotSensorsClient", "DownloadActivation", err.Error())
	}

	req, err := client.DownloadActivationPreparer(ctx, onPremiseIotSensorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.OnPremiseIotSensorsClient", "DownloadActivation", nil, "Failure preparing request")
		return
	}

	resp, err := client.DownloadActivationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.OnPremiseIotSensorsClient", "DownloadActivation", resp, "Failure sending request")
		return
	}

	result, err = client.DownloadActivationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.OnPremiseIotSensorsClient", "DownloadActivation", resp, "Failure responding to request")
	}

	return
}

// DownloadActivationPreparer prepares the DownloadActivation request.
func (client OnPremiseIotSensorsClient) DownloadActivationPreparer(ctx context.Context, onPremiseIotSensorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"onPremiseIotSensorName": autorest.Encode("path", onPremiseIotSensorName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-06-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/onPremiseIotSensors/{onPremiseIotSensorName}/downloadActivation", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DownloadActivationSender sends the DownloadActivation request. The method will close the
// http.Response Body if it receives an error.
func (client OnPremiseIotSensorsClient) DownloadActivationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DownloadActivationResponder handles the response to the DownloadActivation request. The method always
// closes the http.Response Body.
func (client OnPremiseIotSensorsClient) DownloadActivationResponder(resp *http.Response) (result ReadCloser, err error) {
	result.Value = &resp.Body
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK))
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get on-premise IoT sensor
// Parameters:
// onPremiseIotSensorName - name of the on-premise IoT sensor
func (client OnPremiseIotSensorsClient) Get(ctx context.Context, onPremiseIotSensorName string) (result OnPremiseIotSensor, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OnPremiseIotSensorsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.OnPremiseIotSensorsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, onPremiseIotSensorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.OnPremiseIotSensorsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.OnPremiseIotSensorsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.OnPremiseIotSensorsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client OnPremiseIotSensorsClient) GetPreparer(ctx context.Context, onPremiseIotSensorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"onPremiseIotSensorName": autorest.Encode("path", onPremiseIotSensorName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-06-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/onPremiseIotSensors/{onPremiseIotSensorName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client OnPremiseIotSensorsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client OnPremiseIotSensorsClient) GetResponder(resp *http.Response) (result OnPremiseIotSensor, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list on-premise IoT sensors
func (client OnPremiseIotSensorsClient) List(ctx context.Context) (result OnPremiseIotSensorsList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OnPremiseIotSensorsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.OnPremiseIotSensorsClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.OnPremiseIotSensorsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.OnPremiseIotSensorsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.OnPremiseIotSensorsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client OnPremiseIotSensorsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-06-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/onPremiseIotSensors", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client OnPremiseIotSensorsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client OnPremiseIotSensorsClient) ListResponder(resp *http.Response) (result OnPremiseIotSensorsList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}