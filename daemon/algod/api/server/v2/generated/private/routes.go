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
