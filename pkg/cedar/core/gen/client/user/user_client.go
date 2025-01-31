// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	UserAdd(params *UserAddParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserAddOK, error)

	UserFindByID(params *UserFindByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserFindByIDOK, error)

	UserFindByUsername(params *UserFindByUsernameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserFindByUsernameOK, error)

	UserFindList(params *UserFindListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserFindListOK, error)

	UserIDUpdate(params *UserIDUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserIDUpdateOK, error)

	UserNameUpdate(params *UserNameUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserNameUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  UserAdd adds users to a c e d a r application this interface takes an array of user documents

  Add users to a CEDAR application. This interface takes an array of User documents.
*/
func (a *Client) UserAdd(params *UserAddParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserAddOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserAddParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "userAdd",
		Method:             "POST",
		PathPattern:        "/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserAddReader{formats: a.formats},
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
	success, ok := result.(*UserAddOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for userAdd: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UserFindByID retrieves an existing user from a c e d a r application this interface takes an id

  Retrieve an existing user from a CEDAR application. This interface takes an id.
*/
func (a *Client) UserFindByID(params *UserFindByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserFindByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserFindByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "userFindById",
		Method:             "GET",
		PathPattern:        "/user/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserFindByIDReader{formats: a.formats},
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
	success, ok := result.(*UserFindByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for userFindById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UserFindByUsername retrieves an existing user in a c e d a r system based on the user name this interface takes username

  Retrieve an existing user in a CEDAR system based on the user name. This interface takes username.
*/
func (a *Client) UserFindByUsername(params *UserFindByUsernameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserFindByUsernameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserFindByUsernameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "userFindByUsername",
		Method:             "GET",
		PathPattern:        "/user/username/{username}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserFindByUsernameReader{formats: a.formats},
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
	success, ok := result.(*UserFindByUsernameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for userFindByUsername: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UserFindList retrieves a list of users from a c e d a r application this interface takes in application id user name first name last name phone and email

  Retrieve a list of users from a CEDAR application. This interface takes in application, id, userName, firstName, lastName, phone and email.
*/
func (a *Client) UserFindList(params *UserFindListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserFindListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserFindListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "userFindList",
		Method:             "GET",
		PathPattern:        "/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserFindListReader{formats: a.formats},
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
	success, ok := result.(*UserFindListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for userFindList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UserIDUpdate updates a list existing user in a c e d a r application this interface takes an arrary of user documents

  Update a list existing user in a CEDAR application.This interface takes an arrary of User documents.
*/
func (a *Client) UserIDUpdate(params *UserIDUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserIDUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserIDUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "userIdUpdate",
		Method:             "PUT",
		PathPattern:        "/user/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserIDUpdateReader{formats: a.formats},
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
	success, ok := result.(*UserIDUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for userIdUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UserNameUpdate updates a list existing user in a c e d a r system this interface takes an array of user documents

  Update a list existing user in a CEDAR system. This interface takes an array of User documents.
*/
func (a *Client) UserNameUpdate(params *UserNameUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserNameUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserNameUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "userNameUpdate",
		Method:             "PUT",
		PathPattern:        "/user/username/{username}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserNameUpdateReader{formats: a.formats},
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
	success, ok := result.(*UserNameUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for userNameUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
