// Code generated by go-swagger; DO NOT EDIT.

package ods

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
	"github.com/go-openapi/swag"
)

// NewGetMaliciousFilesByIdsParams creates a new GetMaliciousFilesByIdsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMaliciousFilesByIdsParams() *GetMaliciousFilesByIdsParams {
	return &GetMaliciousFilesByIdsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMaliciousFilesByIdsParamsWithTimeout creates a new GetMaliciousFilesByIdsParams object
// with the ability to set a timeout on a request.
func NewGetMaliciousFilesByIdsParamsWithTimeout(timeout time.Duration) *GetMaliciousFilesByIdsParams {
	return &GetMaliciousFilesByIdsParams{
		timeout: timeout,
	}
}

// NewGetMaliciousFilesByIdsParamsWithContext creates a new GetMaliciousFilesByIdsParams object
// with the ability to set a context for a request.
func NewGetMaliciousFilesByIdsParamsWithContext(ctx context.Context) *GetMaliciousFilesByIdsParams {
	return &GetMaliciousFilesByIdsParams{
		Context: ctx,
	}
}

// NewGetMaliciousFilesByIdsParamsWithHTTPClient creates a new GetMaliciousFilesByIdsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMaliciousFilesByIdsParamsWithHTTPClient(client *http.Client) *GetMaliciousFilesByIdsParams {
	return &GetMaliciousFilesByIdsParams{
		HTTPClient: client,
	}
}

/*
GetMaliciousFilesByIdsParams contains all the parameters to send to the API endpoint

	for the get malicious files by ids operation.

	Typically these are written to a http.Request.
*/
type GetMaliciousFilesByIdsParams struct {

	/* Ids.

	   The scan IDs to retrieve the scan entities
	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get malicious files by ids params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMaliciousFilesByIdsParams) WithDefaults() *GetMaliciousFilesByIdsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get malicious files by ids params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMaliciousFilesByIdsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get malicious files by ids params
func (o *GetMaliciousFilesByIdsParams) WithTimeout(timeout time.Duration) *GetMaliciousFilesByIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get malicious files by ids params
func (o *GetMaliciousFilesByIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get malicious files by ids params
func (o *GetMaliciousFilesByIdsParams) WithContext(ctx context.Context) *GetMaliciousFilesByIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get malicious files by ids params
func (o *GetMaliciousFilesByIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get malicious files by ids params
func (o *GetMaliciousFilesByIdsParams) WithHTTPClient(client *http.Client) *GetMaliciousFilesByIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get malicious files by ids params
func (o *GetMaliciousFilesByIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the get malicious files by ids params
func (o *GetMaliciousFilesByIdsParams) WithIds(ids []string) *GetMaliciousFilesByIdsParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get malicious files by ids params
func (o *GetMaliciousFilesByIdsParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetMaliciousFilesByIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Ids != nil {

		// binding items for ids
		joinedIds := o.bindParamIds(reg)

		// query array param ids
		if err := r.SetQueryParam("ids", joinedIds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetMaliciousFilesByIds binds the parameter ids
func (o *GetMaliciousFilesByIdsParams) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []string

		idsIIV := idsIIR // string as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: "multi"
	idsIS := swag.JoinByFormat(idsIC, "multi")

	return idsIS
}