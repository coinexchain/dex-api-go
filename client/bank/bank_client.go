// Code generated by go-swagger; DO NOT EDIT.

package bank

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new bank API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for bank API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetAddressBalances(params *GetAddressBalancesParams) (*GetAddressBalancesOK, *GetAddressBalancesNoContent, error)

	GetBankParams(params *GetBankParamsParams) (*GetBankParamsOK, error)

	SendCoins(params *SendCoinsParams) (*SendCoinsOK, error)

	SetMemoRequired(params *SetMemoRequiredParams) (*SetMemoRequiredOK, error)

	TransferSupervisedCoins(params *TransferSupervisedCoinsParams) (*TransferSupervisedCoinsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAddressBalances gets the account balances
*/
func (a *Client) GetAddressBalances(params *GetAddressBalancesParams) (*GetAddressBalancesOK, *GetAddressBalancesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAddressBalancesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAddressBalances",
		Method:             "GET",
		PathPattern:        "/bank/balances/{address}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAddressBalancesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetAddressBalancesOK:
		return value, nil, nil
	case *GetAddressBalancesNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for bank: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetBankParams gets the current bankx parameters
*/
func (a *Client) GetBankParams(params *GetBankParamsParams) (*GetBankParamsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBankParamsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBankParams",
		Method:             "GET",
		PathPattern:        "/bank/parameters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetBankParamsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBankParamsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBankParams: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SendCoins sends coins from one account to another
*/
func (a *Client) SendCoins(params *SendCoinsParams) (*SendCoinsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendCoinsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sendCoins",
		Method:             "POST",
		PathPattern:        "/bank/accounts/{address}/transfers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SendCoinsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SendCoinsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sendCoins: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetMemoRequired marks if memo is required to receive coins
*/
func (a *Client) SetMemoRequired(params *SetMemoRequiredParams) (*SetMemoRequiredOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetMemoRequiredParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setMemoRequired",
		Method:             "POST",
		PathPattern:        "/bank/accounts/memo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SetMemoRequiredReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetMemoRequiredOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setMemoRequired: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TransferSupervisedCoins operates a supervised transfer
*/
func (a *Client) TransferSupervisedCoins(params *TransferSupervisedCoinsParams) (*TransferSupervisedCoinsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTransferSupervisedCoinsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "transferSupervisedCoins",
		Method:             "POST",
		PathPattern:        "/bank/accounts/{address}/supervised_transfers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TransferSupervisedCoinsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TransferSupervisedCoinsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for transferSupervisedCoins: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
