// Package generated provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get account information.
	// (GET /v2/accounts/{address})
	AccountInformation(ctx echo.Context, address string) error
	// Get a list of unconfirmed transactions currently in the transaction pool by address.
	// (GET /v2/accounts/{address}/transactions/pending)
	GetPendingTransactionsByAddress(ctx echo.Context, address string, params GetPendingTransactionsByAddressParams) error
	// Get the block for the given round.
	// (GET /v2/blocks/{round})
	GetBlock(ctx echo.Context, round uint64, params GetBlockParams) error
	// Get the current supply reported by the ledger.
	// (GET /v2/ledger/supply)
	GetSupply(ctx echo.Context) error
	// Gets the current node status.
	// (GET /v2/status)
	GetStatus(ctx echo.Context) error
	// Gets the node status after waiting for the given round.
	// (GET /v2/status/wait-for-block-after/{round}/)
	WaitForBlock(ctx echo.Context, round uint64) error
	// Broadcasts a raw transaction to the network.
	// (POST /v2/transactions)
	RawTransaction(ctx echo.Context) error
	// Provide debugging information for a transaction (or group).
	// (POST /v2/transactions/dryrun)
	TransactionDryRun(ctx echo.Context) error
	// Get parameters for constructing a new transaction
	// (GET /v2/transactions/params)
	TransactionParams(ctx echo.Context) error
	// Get a list of unconfirmed transactions currently in the transaction pool.
	// (GET /v2/transactions/pending)
	GetPendingTransactions(ctx echo.Context, params GetPendingTransactionsParams) error
	// Get a specific pending transaction.
	// (GET /v2/transactions/pending/{txid})
	PendingTransactionInformation(ctx echo.Context, txid string, params PendingTransactionInformationParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AccountInformation converts echo context to params.
func (w *ServerInterfaceWrapper) AccountInformation(ctx echo.Context) error {

	validQueryParams := map[string]bool{}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AccountInformation(ctx, address)
	return err
}

// GetPendingTransactionsByAddress converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactionsByAddress(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"max":    true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPendingTransactionsByAddressParams
	// ------------- Optional query parameter "max" -------------
	if paramValue := ctx.QueryParam("max"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactionsByAddress(ctx, address, params)
	return err
}

// GetBlock converts echo context to params.
func (w *ServerInterfaceWrapper) GetBlock(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameter("simple", false, "round", ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetBlockParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetBlock(ctx, round, params)
	return err
}

// GetSupply converts echo context to params.
func (w *ServerInterfaceWrapper) GetSupply(ctx echo.Context) error {

	validQueryParams := map[string]bool{}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetSupply(ctx)
	return err
}

// GetStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetStatus(ctx echo.Context) error {

	validQueryParams := map[string]bool{}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStatus(ctx)
	return err
}

// WaitForBlock converts echo context to params.
func (w *ServerInterfaceWrapper) WaitForBlock(ctx echo.Context) error {

	validQueryParams := map[string]bool{}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameter("simple", false, "round", ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.WaitForBlock(ctx, round)
	return err
}

// RawTransaction converts echo context to params.
func (w *ServerInterfaceWrapper) RawTransaction(ctx echo.Context) error {

	validQueryParams := map[string]bool{}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RawTransaction(ctx)
	return err
}

// TransactionDryRun converts echo context to params.
func (w *ServerInterfaceWrapper) TransactionDryRun(ctx echo.Context) error {

	validQueryParams := map[string]bool{}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TransactionDryRun(ctx)
	return err
}

// TransactionParams converts echo context to params.
func (w *ServerInterfaceWrapper) TransactionParams(ctx echo.Context) error {

	validQueryParams := map[string]bool{}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TransactionParams(ctx)
	return err
}

// GetPendingTransactions converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactions(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"max":    true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPendingTransactionsParams
	// ------------- Optional query parameter "max" -------------
	if paramValue := ctx.QueryParam("max"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactions(ctx, params)
	return err
}

// PendingTransactionInformation converts echo context to params.
func (w *ServerInterfaceWrapper) PendingTransactionInformation(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "txid" -------------
	var txid string

	err = runtime.BindStyledParameter("simple", false, "txid", ctx.Param("txid"), &txid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter txid: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PendingTransactionInformationParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PendingTransactionInformation(ctx, txid, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/v2/accounts/:address", wrapper.AccountInformation, m...)
	router.GET("/v2/accounts/:address/transactions/pending", wrapper.GetPendingTransactionsByAddress, m...)
	router.GET("/v2/blocks/:round", wrapper.GetBlock, m...)
	router.GET("/v2/ledger/supply", wrapper.GetSupply, m...)
	router.GET("/v2/status", wrapper.GetStatus, m...)
	router.GET("/v2/status/wait-for-block-after/:round/", wrapper.WaitForBlock, m...)
	router.POST("/v2/transactions", wrapper.RawTransaction, m...)
	router.POST("/v2/transactions/dryrun", wrapper.TransactionDryRun, m...)
	router.GET("/v2/transactions/params", wrapper.TransactionParams, m...)
	router.GET("/v2/transactions/pending", wrapper.GetPendingTransactions, m...)
	router.GET("/v2/transactions/pending/:txid", wrapper.PendingTransactionInformation, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"/9O7MMdamgAA",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
