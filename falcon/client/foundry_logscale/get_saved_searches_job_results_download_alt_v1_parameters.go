// Code generated by go-swagger; DO NOT EDIT.

package foundry_logscale

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetSavedSearchesJobResultsDownloadAltV1Params creates a new GetSavedSearchesJobResultsDownloadAltV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSavedSearchesJobResultsDownloadAltV1Params() *GetSavedSearchesJobResultsDownloadAltV1Params {
	return &GetSavedSearchesJobResultsDownloadAltV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSavedSearchesJobResultsDownloadAltV1ParamsWithTimeout creates a new GetSavedSearchesJobResultsDownloadAltV1Params object
// with the ability to set a timeout on a request.
func NewGetSavedSearchesJobResultsDownloadAltV1ParamsWithTimeout(timeout time.Duration) *GetSavedSearchesJobResultsDownloadAltV1Params {
	return &GetSavedSearchesJobResultsDownloadAltV1Params{
		timeout: timeout,
	}
}

// NewGetSavedSearchesJobResultsDownloadAltV1ParamsWithContext creates a new GetSavedSearchesJobResultsDownloadAltV1Params object
// with the ability to set a context for a request.
func NewGetSavedSearchesJobResultsDownloadAltV1ParamsWithContext(ctx context.Context) *GetSavedSearchesJobResultsDownloadAltV1Params {
	return &GetSavedSearchesJobResultsDownloadAltV1Params{
		Context: ctx,
	}
}

// NewGetSavedSearchesJobResultsDownloadAltV1ParamsWithHTTPClient creates a new GetSavedSearchesJobResultsDownloadAltV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetSavedSearchesJobResultsDownloadAltV1ParamsWithHTTPClient(client *http.Client) *GetSavedSearchesJobResultsDownloadAltV1Params {
	return &GetSavedSearchesJobResultsDownloadAltV1Params{
		HTTPClient: client,
	}
}

/*
GetSavedSearchesJobResultsDownloadAltV1Params contains all the parameters to send to the API endpoint

	for the get saved searches job results download alt v1 operation.

	Typically these are written to a http.Request.
*/
type GetSavedSearchesJobResultsDownloadAltV1Params struct {

	/* JobID.

	   Job ID for a previously executed async query
	*/
	JobID string

	/* ResultFormat.

	   Result Format
	*/
	ResultFormat *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get saved searches job results download alt v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSavedSearchesJobResultsDownloadAltV1Params) WithDefaults() *GetSavedSearchesJobResultsDownloadAltV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get saved searches job results download alt v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSavedSearchesJobResultsDownloadAltV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get saved searches job results download alt v1 params
func (o *GetSavedSearchesJobResultsDownloadAltV1Params) WithTimeout(timeout time.Duration) *GetSavedSearchesJobResultsDownloadAltV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get saved searches job results download alt v1 params
func (o *GetSavedSearchesJobResultsDownloadAltV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get saved searches job results download alt v1 params
func (o *GetSavedSearchesJobResultsDownloadAltV1Params) WithContext(ctx context.Context) *GetSavedSearchesJobResultsDownloadAltV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get saved searches job results download alt v1 params
func (o *GetSavedSearchesJobResultsDownloadAltV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get saved searches job results download alt v1 params
func (o *GetSavedSearchesJobResultsDownloadAltV1Params) WithHTTPClient(client *http.Client) *GetSavedSearchesJobResultsDownloadAltV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get saved searches job results download alt v1 params
func (o *GetSavedSearchesJobResultsDownloadAltV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobID adds the jobID to the get saved searches job results download alt v1 params
func (o *GetSavedSearchesJobResultsDownloadAltV1Params) WithJobID(jobID string) *GetSavedSearchesJobResultsDownloadAltV1Params {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the get saved searches job results download alt v1 params
func (o *GetSavedSearchesJobResultsDownloadAltV1Params) SetJobID(jobID string) {
	o.JobID = jobID
}

// WithResultFormat adds the resultFormat to the get saved searches job results download alt v1 params
func (o *GetSavedSearchesJobResultsDownloadAltV1Params) WithResultFormat(resultFormat *string) *GetSavedSearchesJobResultsDownloadAltV1Params {
	o.SetResultFormat(resultFormat)
	return o
}

// SetResultFormat adds the resultFormat to the get saved searches job results download alt v1 params
func (o *GetSavedSearchesJobResultsDownloadAltV1Params) SetResultFormat(resultFormat *string) {
	o.ResultFormat = resultFormat
}

// WriteToRequest writes these params to a swagger request
func (o *GetSavedSearchesJobResultsDownloadAltV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param job_id
	qrJobID := o.JobID
	qJobID := qrJobID
	if qJobID != "" {

		if err := r.SetQueryParam("job_id", qJobID); err != nil {
			return err
		}
	}

	if o.ResultFormat != nil {

		// query param result_format
		var qrResultFormat string

		if o.ResultFormat != nil {
			qrResultFormat = *o.ResultFormat
		}
		qResultFormat := qrResultFormat
		if qResultFormat != "" {

			if err := r.SetQueryParam("result_format", qResultFormat); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}