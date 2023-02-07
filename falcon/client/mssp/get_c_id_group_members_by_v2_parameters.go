// Code generated by go-swagger; DO NOT EDIT.

package mssp

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

// NewGetCIDGroupMembersByV2Params creates a new GetCIDGroupMembersByV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCIDGroupMembersByV2Params() *GetCIDGroupMembersByV2Params {
	return &GetCIDGroupMembersByV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCIDGroupMembersByV2ParamsWithTimeout creates a new GetCIDGroupMembersByV2Params object
// with the ability to set a timeout on a request.
func NewGetCIDGroupMembersByV2ParamsWithTimeout(timeout time.Duration) *GetCIDGroupMembersByV2Params {
	return &GetCIDGroupMembersByV2Params{
		timeout: timeout,
	}
}

// NewGetCIDGroupMembersByV2ParamsWithContext creates a new GetCIDGroupMembersByV2Params object
// with the ability to set a context for a request.
func NewGetCIDGroupMembersByV2ParamsWithContext(ctx context.Context) *GetCIDGroupMembersByV2Params {
	return &GetCIDGroupMembersByV2Params{
		Context: ctx,
	}
}

// NewGetCIDGroupMembersByV2ParamsWithHTTPClient creates a new GetCIDGroupMembersByV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetCIDGroupMembersByV2ParamsWithHTTPClient(client *http.Client) *GetCIDGroupMembersByV2Params {
	return &GetCIDGroupMembersByV2Params{
		HTTPClient: client,
	}
}

/*
GetCIDGroupMembersByV2Params contains all the parameters to send to the API endpoint

	for the get c ID group members by v2 operation.

	Typically these are written to a http.Request.
*/
type GetCIDGroupMembersByV2Params struct {

	/* Ids.

	   CID group IDs search for
	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get c ID group members by v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCIDGroupMembersByV2Params) WithDefaults() *GetCIDGroupMembersByV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get c ID group members by v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCIDGroupMembersByV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get c ID group members by v2 params
func (o *GetCIDGroupMembersByV2Params) WithTimeout(timeout time.Duration) *GetCIDGroupMembersByV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get c ID group members by v2 params
func (o *GetCIDGroupMembersByV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get c ID group members by v2 params
func (o *GetCIDGroupMembersByV2Params) WithContext(ctx context.Context) *GetCIDGroupMembersByV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get c ID group members by v2 params
func (o *GetCIDGroupMembersByV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get c ID group members by v2 params
func (o *GetCIDGroupMembersByV2Params) WithHTTPClient(client *http.Client) *GetCIDGroupMembersByV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get c ID group members by v2 params
func (o *GetCIDGroupMembersByV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the get c ID group members by v2 params
func (o *GetCIDGroupMembersByV2Params) WithIds(ids []string) *GetCIDGroupMembersByV2Params {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get c ID group members by v2 params
func (o *GetCIDGroupMembersByV2Params) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetCIDGroupMembersByV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamGetCIDGroupMembersByV2 binds the parameter ids
func (o *GetCIDGroupMembersByV2Params) bindParamIds(formats strfmt.Registry) []string {
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