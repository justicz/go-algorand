// Package private provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package private

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
	// Aborts a catchpoint catchup.
	// (DELETE /v2/catchup/{catchpoint})
	AbortCatchup(ctx echo.Context, catchpoint string) error
	// Starts a catchpoint catchup.
	// (POST /v2/catchup/{catchpoint})
	StartCatchup(ctx echo.Context, catchpoint string) error

	// (POST /v2/register-participation-keys/{address})
	RegisterParticipationKeys(ctx echo.Context, address string, params RegisterParticipationKeysParams) error

	// (POST /v2/shutdown)
	ShutdownNode(ctx echo.Context, params ShutdownNodeParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AbortCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) AbortCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AbortCatchup(ctx, catchpoint)
	return err
}

// StartCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) StartCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.StartCatchup(ctx, catchpoint)
	return err
}

// RegisterParticipationKeys converts echo context to params.
func (w *ServerInterfaceWrapper) RegisterParticipationKeys(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"fee":              true,
		"key-dilution":     true,
		"round-last-valid": true,
		"no-wait":          true,
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
	var params RegisterParticipationKeysParams
	// ------------- Optional query parameter "fee" -------------
	if paramValue := ctx.QueryParam("fee"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "fee", ctx.QueryParams(), &params.Fee)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter fee: %s", err))
	}

	// ------------- Optional query parameter "key-dilution" -------------
	if paramValue := ctx.QueryParam("key-dilution"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "key-dilution", ctx.QueryParams(), &params.KeyDilution)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter key-dilution: %s", err))
	}

	// ------------- Optional query parameter "round-last-valid" -------------
	if paramValue := ctx.QueryParam("round-last-valid"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "round-last-valid", ctx.QueryParams(), &params.RoundLastValid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round-last-valid: %s", err))
	}

	// ------------- Optional query parameter "no-wait" -------------
	if paramValue := ctx.QueryParam("no-wait"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "no-wait", ctx.QueryParams(), &params.NoWait)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter no-wait: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RegisterParticipationKeys(ctx, address, params)
	return err
}

// ShutdownNode converts echo context to params.
func (w *ServerInterfaceWrapper) ShutdownNode(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"timeout": true,
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
	var params ShutdownNodeParams
	// ------------- Optional query parameter "timeout" -------------
	if paramValue := ctx.QueryParam("timeout"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "timeout", ctx.QueryParams(), &params.Timeout)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timeout: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ShutdownNode(ctx, params)
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

	router.DELETE("/v2/catchup/:catchpoint", wrapper.AbortCatchup, m...)
	router.POST("/v2/catchup/:catchpoint", wrapper.StartCatchup, m...)
	router.POST("/v2/register-participation-keys/:address", wrapper.RegisterParticipationKeys, m...)
	router.POST("/v2/shutdown", wrapper.ShutdownNode, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
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
