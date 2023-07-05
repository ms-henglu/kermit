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

// JobsClient is the azure IoT Central is a service that makes it easy to connect, monitor, and manage your IoT devices
// at scale.
type JobsClient struct {
	BaseClient
}

// NewJobsClient creates an instance of the JobsClient client.
func NewJobsClient(subdomain string) JobsClient {
	return JobsClient{New(subdomain)}
}

// Create create and execute a new job via its job definition.
// Parameters:
// jobID - unique ID of the job.
// body - job definition.
func (client JobsClient) Create(ctx context.Context, jobID string, body Job) (result Job, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: jobID,
			Constraints: []validation.Constraint{{Target: "jobID", Name: validation.MaxLength, Rule: 255, Chain: nil},
				{Target: "jobID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]*$`, Chain: nil}}},
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.Group", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "body.Batch", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.Batch.Value", Name: validation.Null, Rule: true, Chain: nil}}},
				{Target: "body.CancellationThreshold", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.CancellationThreshold.Value", Name: validation.Null, Rule: true, Chain: nil}}},
				{Target: "body.Data", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "body.Data", Name: validation.MinItems, Rule: 1, Chain: nil}}},
				{Target: "body.Organizations", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.Organizations", Name: validation.MaxItems, Rule: 1, Chain: nil},
						{Target: "body.Organizations", Name: validation.MinItems, Rule: 1, Chain: nil},
					}}}}}); err != nil {
		return result, validation.NewError("iotcentral.JobsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, jobID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.JobsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.JobsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.JobsClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client JobsClient) CreatePreparer(ctx context.Context, jobID string, body Job) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"jobId": autorest.Encode("path", jobID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	body.ID = nil
	body.Start = nil
	body.End = nil
	body.Progress = nil
	body.Status = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/jobs/{jobId}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client JobsClient) CreateResponder(resp *http.Response) (result Job, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get details about a running or completed job by job ID.
// Parameters:
// jobID - unique ID of the job.
func (client JobsClient) Get(ctx context.Context, jobID string) (result Job, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: jobID,
			Constraints: []validation.Constraint{{Target: "jobID", Name: validation.MaxLength, Rule: 255, Chain: nil},
				{Target: "jobID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.JobsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, jobID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.JobsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.JobsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.JobsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client JobsClient) GetPreparer(ctx context.Context, jobID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"jobId": autorest.Encode("path", jobID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/jobs/{jobId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client JobsClient) GetResponder(resp *http.Response) (result Job, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDevices get the list of individual device statuses by job ID.
// Parameters:
// jobID - unique ID of the job.
func (client JobsClient) GetDevices(ctx context.Context, jobID string) (result JobDeviceStatusCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.GetDevices")
		defer func() {
			sc := -1
			if result.jdsc.Response.Response != nil {
				sc = result.jdsc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: jobID,
			Constraints: []validation.Constraint{{Target: "jobID", Name: validation.MaxLength, Rule: 255, Chain: nil},
				{Target: "jobID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.JobsClient", "GetDevices", err.Error())
	}

	result.fn = client.getDevicesNextResults
	req, err := client.GetDevicesPreparer(ctx, jobID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.JobsClient", "GetDevices", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDevicesSender(req)
	if err != nil {
		result.jdsc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.JobsClient", "GetDevices", resp, "Failure sending request")
		return
	}

	result.jdsc, err = client.GetDevicesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.JobsClient", "GetDevices", resp, "Failure responding to request")
		return
	}
	if result.jdsc.hasNextLink() && result.jdsc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// GetDevicesPreparer prepares the GetDevices request.
func (client JobsClient) GetDevicesPreparer(ctx context.Context, jobID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"jobId": autorest.Encode("path", jobID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/jobs/{jobId}/devices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDevicesSender sends the GetDevices request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) GetDevicesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetDevicesResponder handles the response to the GetDevices request. The method always
// closes the http.Response Body.
func (client JobsClient) GetDevicesResponder(resp *http.Response) (result JobDeviceStatusCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getDevicesNextResults retrieves the next set of results, if any.
func (client JobsClient) getDevicesNextResults(ctx context.Context, lastResults JobDeviceStatusCollection) (result JobDeviceStatusCollection, err error) {
	req, err := lastResults.jobDeviceStatusCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "iotcentral.JobsClient", "getDevicesNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetDevicesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "iotcentral.JobsClient", "getDevicesNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetDevicesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.JobsClient", "getDevicesNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetDevicesComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobsClient) GetDevicesComplete(ctx context.Context, jobID string) (result JobDeviceStatusCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.GetDevices")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetDevices(ctx, jobID)
	return
}

// List sends the list request.
// Parameters:
// maxpagesize - the maximum number of resources to return from one response.
func (client JobsClient) List(ctx context.Context, maxpagesize *int32) (result JobCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.List")
		defer func() {
			sc := -1
			if result.jc.Response.Response != nil {
				sc = result.jc.Response.Response.StatusCode
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
		return result, validation.NewError("iotcentral.JobsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, maxpagesize)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.JobsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.jc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.JobsClient", "List", resp, "Failure sending request")
		return
	}

	result.jc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.JobsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.jc.hasNextLink() && result.jc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client JobsClient) ListPreparer(ctx context.Context, maxpagesize *int32) (*http.Request, error) {
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
		autorest.WithPath("/jobs"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client JobsClient) ListResponder(resp *http.Response) (result JobCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client JobsClient) listNextResults(ctx context.Context, lastResults JobCollection) (result JobCollection, err error) {
	req, err := lastResults.jobCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "iotcentral.JobsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "iotcentral.JobsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.JobsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobsClient) ListComplete(ctx context.Context, maxpagesize *int32) (result JobCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.List")
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

// Rerun execute a rerun of an existing job on all failed devices.
// Parameters:
// jobID - unique ID of the job.
// rerunID - unique ID of the job rerun.
func (client JobsClient) Rerun(ctx context.Context, jobID string, rerunID string) (result Job, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.Rerun")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: jobID,
			Constraints: []validation.Constraint{{Target: "jobID", Name: validation.MaxLength, Rule: 255, Chain: nil},
				{Target: "jobID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.JobsClient", "Rerun", err.Error())
	}

	req, err := client.RerunPreparer(ctx, jobID, rerunID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.JobsClient", "Rerun", nil, "Failure preparing request")
		return
	}

	resp, err := client.RerunSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.JobsClient", "Rerun", resp, "Failure sending request")
		return
	}

	result, err = client.RerunResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.JobsClient", "Rerun", resp, "Failure responding to request")
		return
	}

	return
}

// RerunPreparer prepares the Rerun request.
func (client JobsClient) RerunPreparer(ctx context.Context, jobID string, rerunID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"jobId":   autorest.Encode("path", jobID),
		"rerunId": autorest.Encode("path", rerunID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/jobs/{jobId}/rerun/{rerunId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RerunSender sends the Rerun request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) RerunSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RerunResponder handles the response to the Rerun request. The method always
// closes the http.Response Body.
func (client JobsClient) RerunResponder(resp *http.Response) (result Job, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Resume resume execution of an existing stopped job.
// Parameters:
// jobID - unique ID of the job.
func (client JobsClient) Resume(ctx context.Context, jobID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.Resume")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: jobID,
			Constraints: []validation.Constraint{{Target: "jobID", Name: validation.MaxLength, Rule: 255, Chain: nil},
				{Target: "jobID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.JobsClient", "Resume", err.Error())
	}

	req, err := client.ResumePreparer(ctx, jobID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.JobsClient", "Resume", nil, "Failure preparing request")
		return
	}

	resp, err := client.ResumeSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "iotcentral.JobsClient", "Resume", resp, "Failure sending request")
		return
	}

	result, err = client.ResumeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.JobsClient", "Resume", resp, "Failure responding to request")
		return
	}

	return
}

// ResumePreparer prepares the Resume request.
func (client JobsClient) ResumePreparer(ctx context.Context, jobID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"jobId": autorest.Encode("path", jobID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/jobs/{jobId}/resume", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResumeSender sends the Resume request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) ResumeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ResumeResponder handles the response to the Resume request. The method always
// closes the http.Response Body.
func (client JobsClient) ResumeResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Stop stop execution of a job that is currently running.
// Parameters:
// jobID - unique ID of the job.
func (client JobsClient) Stop(ctx context.Context, jobID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobsClient.Stop")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: jobID,
			Constraints: []validation.Constraint{{Target: "jobID", Name: validation.MaxLength, Rule: 255, Chain: nil},
				{Target: "jobID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.JobsClient", "Stop", err.Error())
	}

	req, err := client.StopPreparer(ctx, jobID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.JobsClient", "Stop", nil, "Failure preparing request")
		return
	}

	resp, err := client.StopSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "iotcentral.JobsClient", "Stop", resp, "Failure sending request")
		return
	}

	result, err = client.StopResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.JobsClient", "Stop", resp, "Failure responding to request")
		return
	}

	return
}

// StopPreparer prepares the Stop request.
func (client JobsClient) StopPreparer(ctx context.Context, jobID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"jobId": autorest.Encode("path", jobID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/jobs/{jobId}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client JobsClient) StopSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client JobsClient) StopResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}