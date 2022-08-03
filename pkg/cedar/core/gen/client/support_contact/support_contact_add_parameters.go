// Code generated by go-swagger; DO NOT EDIT.

package support_contact

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

	"github.com/cmsgov/easi-app/pkg/cedar/core/gen/models"
)

// NewSupportContactAddParams creates a new SupportContactAddParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSupportContactAddParams() *SupportContactAddParams {
	return &SupportContactAddParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSupportContactAddParamsWithTimeout creates a new SupportContactAddParams object
// with the ability to set a timeout on a request.
func NewSupportContactAddParamsWithTimeout(timeout time.Duration) *SupportContactAddParams {
	return &SupportContactAddParams{
		timeout: timeout,
	}
}

// NewSupportContactAddParamsWithContext creates a new SupportContactAddParams object
// with the ability to set a context for a request.
func NewSupportContactAddParamsWithContext(ctx context.Context) *SupportContactAddParams {
	return &SupportContactAddParams{
		Context: ctx,
	}
}

// NewSupportContactAddParamsWithHTTPClient creates a new SupportContactAddParams object
// with the ability to set a custom HTTPClient for a request.
func NewSupportContactAddParamsWithHTTPClient(client *http.Client) *SupportContactAddParams {
	return &SupportContactAddParams{
		HTTPClient: client,
	}
}

/* SupportContactAddParams contains all the parameters to send to the API endpoint
   for the support contact add operation.

   Typically these are written to a http.Request.
*/
type SupportContactAddParams struct {

	// Body.
	Body *models.SupportContactAddRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the support contact add params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SupportContactAddParams) WithDefaults() *SupportContactAddParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the support contact add params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SupportContactAddParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the support contact add params
func (o *SupportContactAddParams) WithTimeout(timeout time.Duration) *SupportContactAddParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the support contact add params
func (o *SupportContactAddParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the support contact add params
func (o *SupportContactAddParams) WithContext(ctx context.Context) *SupportContactAddParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the support contact add params
func (o *SupportContactAddParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the support contact add params
func (o *SupportContactAddParams) WithHTTPClient(client *http.Client) *SupportContactAddParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the support contact add params
func (o *SupportContactAddParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the support contact add params
func (o *SupportContactAddParams) WithBody(body *models.SupportContactAddRequest) *SupportContactAddParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the support contact add params
func (o *SupportContactAddParams) SetBody(body *models.SupportContactAddRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SupportContactAddParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
