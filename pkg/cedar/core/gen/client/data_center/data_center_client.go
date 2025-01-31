// Code generated by go-swagger; DO NOT EDIT.

package data_center

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new data center API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for data center API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DataCenterFindList(params *DataCenterFindListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DataCenterFindListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DataCenterFindList retrieves a list of data centers based on query criteria id name version state status and ids only if ids onlys is true only the id and name will be returned

  Retrieve a list of data centers based on query criteria (id, name, version, state, status and idsOnly). If idsOnlys is true, only the id and name will be returned.
*/
func (a *Client) DataCenterFindList(params *DataCenterFindListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DataCenterFindListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDataCenterFindListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dataCenterFindList",
		Method:             "GET",
		PathPattern:        "/dataCenter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DataCenterFindListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DataCenterFindListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dataCenterFindList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
