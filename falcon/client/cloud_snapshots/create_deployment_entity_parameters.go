// Code generated by go-swagger; DO NOT EDIT.

package cloud_snapshots

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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewCreateDeploymentEntityParams creates a new CreateDeploymentEntityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateDeploymentEntityParams() *CreateDeploymentEntityParams {
	return &CreateDeploymentEntityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDeploymentEntityParamsWithTimeout creates a new CreateDeploymentEntityParams object
// with the ability to set a timeout on a request.
func NewCreateDeploymentEntityParamsWithTimeout(timeout time.Duration) *CreateDeploymentEntityParams {
	return &CreateDeploymentEntityParams{
		timeout: timeout,
	}
}

// NewCreateDeploymentEntityParamsWithContext creates a new CreateDeploymentEntityParams object
// with the ability to set a context for a request.
func NewCreateDeploymentEntityParamsWithContext(ctx context.Context) *CreateDeploymentEntityParams {
	return &CreateDeploymentEntityParams{
		Context: ctx,
	}
}

// NewCreateDeploymentEntityParamsWithHTTPClient creates a new CreateDeploymentEntityParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateDeploymentEntityParamsWithHTTPClient(client *http.Client) *CreateDeploymentEntityParams {
	return &CreateDeploymentEntityParams{
		HTTPClient: client,
	}
}

/*
CreateDeploymentEntityParams contains all the parameters to send to the API endpoint

	for the create deployment entity operation.

	Typically these are written to a http.Request.
*/
type CreateDeploymentEntityParams struct {

	// Body.
	Body *models.ModelsCreateDeploymentInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create deployment entity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDeploymentEntityParams) WithDefaults() *CreateDeploymentEntityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create deployment entity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDeploymentEntityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create deployment entity params
func (o *CreateDeploymentEntityParams) WithTimeout(timeout time.Duration) *CreateDeploymentEntityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create deployment entity params
func (o *CreateDeploymentEntityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create deployment entity params
func (o *CreateDeploymentEntityParams) WithContext(ctx context.Context) *CreateDeploymentEntityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create deployment entity params
func (o *CreateDeploymentEntityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create deployment entity params
func (o *CreateDeploymentEntityParams) WithHTTPClient(client *http.Client) *CreateDeploymentEntityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create deployment entity params
func (o *CreateDeploymentEntityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create deployment entity params
func (o *CreateDeploymentEntityParams) WithBody(body *models.ModelsCreateDeploymentInput) *CreateDeploymentEntityParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create deployment entity params
func (o *CreateDeploymentEntityParams) SetBody(body *models.ModelsCreateDeploymentInput) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDeploymentEntityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}