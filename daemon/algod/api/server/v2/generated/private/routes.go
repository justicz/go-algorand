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

	validQueryParams := map[string]bool{
		"pretty": true,
	}

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

	validQueryParams := map[string]bool{
		"pretty": true,
	}

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
		"pretty":           true,
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
		"pretty":  true,
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

	"H4sIAAAAAAAC/+x9f3PcNrLgV8HNe1WxfcMZ+Vd2rarUnmLlhy6O47KUvbu1fBsM2TODiAQYAtRo4tN3",
	"v+oGQIIkODOytd5Nvf3L1gBoNBrdje5Go/lhkqqiVBKk0ZPjD5OSV7wAAxX9xdNU1dIkIsO/MtBpJUoj",
	"lJwc+zamTSXkajKdCPy15GY9mU4kL6Dtg+Onkwp+q0UF2eTYVDVMJzpdQ8ERsNmW2LuBdJOsVOJAnFgQ",
	"Z6eT2x0NPMsq0HqI5U8y3zIh07zOgJmKS81TbNJsI8yambXQzA1mQjIlgaklM+tOZ7YUkGd65hf5Ww3V",
	"Nlilm3x8Sbctikmlchji+VIVCyHBYwUNUs2GMKNYBkvqtOaG4QyIq+9oFNPAq3TNlqrag6pFIsQXZF1M",
	"jt9NNMgMKtqtFMQ1/XdZAfwOieHVCszk/TS2uKWBKjGiiCztzFG/Al3nRjPqS2tciWuQDEfN2I+1NmwB",
	"jEv29tuX7OnTpy9wIQU3BjLHZKOramcP12SHT44nGTfgm4e8xvOVqrjMkqb/229f0vznboGH9uJaQ1xY",
	"TrCFnZ2OLcAPjLCQkAZWtA8d7scREaFof17AUlVw4J7Yzve6KeH8/9RdSblJ16US0kT2hVErs81RHRYM",
	"36XDGgQ6/UukVIVA3x0lL95/eDx9fHT7H+9Okr+5P58/vT1w+S8buHsoEO2Y1lUFMt0mqwo4ScuayyE9",
	"3jp+0GtV5xlb82vafF6QqndjGY61qvOa5zXyiUgrdZKvlGbcsVEGS17nhvmJWS1zVFMIzXE7E5qVlboW",
	"GWRT1L6btUjXLOXagqB+bCPyHHmw1pCN8Vp8dTuE6TYkCeL1UfSgBf3rEqNd1x5KwA1pgyTNlYbEqD3H",
	"kz9xuMxYeKC0Z5W+22HFLtbAaHJssIct0U4iT+f5lhna14xxzTjzR9OUiSXbqpptaHNycUXj3WqQagVD",
	"otHmdM5RFN4x8g2IESHeQqkcuCTiebkbkkwuxaquQLPNGszanXkV6FJJDUwtfoXU4Lb/z/OfXjNVsR9B",
	"a76CNzy9YiBTlY3vsZs0doL/qhVueKFXJU+v4sd1LgoRQflHfiOKumCyLhZQ4X7588EoVoGpKzmGkIW4",
	"h88KfjOc9KKqZUqb207bMdSQlYQuc76dsbMlK/jNV0dTh45mPM9ZCTITcsXMjRw10nDu/egllapldoAN",
	"Y3DDglNTl5CKpYCMNVB2YOKm2YePkHfDp7WsAnQ8kFF0mln2oCPhJsIzKLrYwkq+goBlZuxnp7mo1agr",
	"kI2CY4stNZUVXAtV62bQCI409W7zWioDSVnBUkR47NyRA7WH7ePUa+EMnFRJw4WEDDUvIa0MWE00ilMw",
	"4W5nZnhEL7iGL5+NHeBt64G7v1T9Xd+54wftNnVKrEhGzkVsdQIbN5s64w9w/sK5tVgl9ufBRorVBR4l",
	"S5HTMfMr7p8nQ61JCXQI4Q8eLVaSm7qC40v5CP9iCTs3XGa8yvCXwv70Y50bcS5W+FNuf3qlViI9F6sR",
	"Yja4Rr0pGlbYfxBeXB2bm6jT8Eqpq7oMF5R2vNLFlp2djm2yhXlXxjxpXNnQq7i48Z7GXUeYm2YjR5Ac",
	"pV3JseMVbCtAbHm6pH9ulsRPfFn9HiMmcq47YSka4KIEb91v+BPKOlhngJdlLlKO1JzTuXn8IcDkPytY",
	"To4n/zFvQyRz26rnDq6dsbttD6AozfYhLv+khX//GLQjY1gEzUxIu13UdWqdxPvHB6FGMSHLtYfD17lK",
	"rz4Kh7JSJVRG2P1dIJyh6BB4tgaeQcUybvis9bKs4TUiADTwexpHbhNUkTPvJ/oPzxk2o1hy4+05tGWF",
	"RqtOBZGnDE1Ae7DYmbADmaaKFdbqY2it3QnLl+3kVmM3KvadI8v7PrTI7nxjDU1GI/wicOmtG3myUNXH",
	"8UuPESRrnWPGEWpjDuPKuztLXesycfSJGNi2Qw9QG48c6tmQQn3wh9AqkOyWOueG/wOooxHqfVCnC+gz",
	"Uee02la1vAfxhqpSVcTgw5NFUg9hoND7dJPF5+JGWn+exluAvKr4drB8O62b5JC1dxfszUnNSqgScyNZ",
	"Bot6FepAtqxUwTjLaCAJ3GuVwbnhptb3wE0tsBYZVD8hCnyhasM4kypDxsDOcT4bCZ6R107BBhOyrllb",
	"/bYANMdSXq/WhqEdo4ZsF0bnEp7aHUhIF+kRX6NxEm0vO50NzOQV8GzLFgCSqYUz6J2rQYvkFAcwPsTv",
	"uLxFqzFCO3iVlUpBa8gSd5+xFzV/N0KbbHaQifAmfJtJmFZsyauPxNUow/M9eFKfIba6Pa2cEzTE+rDp",
	"d+1ff/JwF3mFPo1lAjwaUZJzMDBGwr00qcuR+LfTjheiQJFgkkulIVUy01FgOdcm2ScK2KmjwnFbA+6L",
	"cT8BHvHyXnFtrJ8lZEbHvBVhmofG0BTjCF9DpYWScch/tY0x2CnqHqlrzRwEpuuyVJWBLLYGdM7H53oN",
	"N81cahnALitlVKpy3Ohawz7IY1QK4Dti2ZVYAnHjHP0mEDFcHMVUUbduo6TsINESYhci575XQN0wBjiC",
	"CNqEzUhiHKF7nNMEHqcTbVRZok4ySS2bcWNkOre9T8zPbd8hc3HT6spMAc5uPE4O842lrI3+rrlmDg9W",
	"8CvU92WlVs4hHOKMwphoIVNIdnE+iuU59gpFYI+Qjhgw7n4pmK0nHD3+jTLdKBPs2YWxBd/Rmnpjw5sX",
	"ret/DwbCKRguct0YAU0MtZ2Fwq39q/AN1xSAlybfIg8vRVXYGws6O7T/zZoYmZvFxuZbsZQZq2DDq8z3",
	"GFq47mJEZnAT17fc+ZUZ3DARR3TZzCYMS/0dgrt0mcXPDQr7W+R07EKIGpAfC5FWitt7HiS8PbNMc5VR",
	"QcERO7pxcGfs+JxCrhJ7rRQ5rWy7v3by4b5wq+Jw/faMClqzI5s1UCQbtWePiOEmL1lZgYaxhZRK5Ulj",
	"s/eDlgM905/pSqRXkDFkSLJ6nPr7oosTTsIe4KbqJqy7WW+9QVWWICF7OGPsRDISIuf09I663uTyC7Nr",
	"/huaNavpholLRoucXcrYseXvpz6RizyY3bxjEzY+cSoLZPdE5kaOMBDfUHgVwUU5cmco45xGBrptoMoD",
	"prJYHKI+v6MsBt7ZZZGRtduqL10vCkGpDEG3KeoKf7s0dJeEmTF2QdKC5qqGa6h4Tve02kd5hGaFQK9H",
	"12kKkB1fyqSDSaoKN/GD9r9WEC/ro6OnwI4e9sdog3aKs8ytDPTHfsWOpraJyMW+YpeTy8kAUgWFuobM",
	"eichX9tRe8H+twbupfxpoIpYwbfWr/GyyHS9XIpUWKLnCjXZSvXMDamoBSpED9A70EyYKSlvoiiZaXZf",
	"WgGMH4/34UBHoKKBhodHVfGtv1Po8o5mcMNTXCUnJbNlG2SUhs+Gp5xRZRICiMaFdszoInb25syHQj5S",
	"7vpBkenEunO78bvoOXQdcgTsOttvtA2IEcXgEPE/YaXCXRcue8BfMedCmwGSzrOkcG3DkJFDZ8b+j6pZ",
	"ykl+y9pAY9Sriixl8qBwBjpF/ZzONmkpBDkUYP1tann0qL/wR4/cngvNlrDxKTfYsU+OR4+sEChtXqqi",
	"FDncQ9htzfV6uNMLruHpE3b+/cnzx0/+/uT5l7gYsvd5wRZbPFgfuOsgps02h4fx05FCcFHoXz7ziQ9d",
	"uHvjlYRwA/sQDrkA1NqWYqwNCyIdP1mT9ET85ixietE60SqJpJviamZ710xwD1pqAPrs1E9ISklrOqpv",
	"pxP0WfPtPShOC4hV4CxF3YneaNuqlmGalJMDvdUGimEI0g79+4gN+9a7WgOLRclcSEgKJWEbzQwWEn6k",
	"xqi9Q6I2MpiU3tjYvivawb+HVneeQ3bzU+lLux2wxJsmaeseNr8Ptxd9DhPEyFqHvGScpbmgyJ6S2lR1",
	"ai4lp0hDz5zssYWPn4zHnl76LvFgVyQW5UBdSq6Rhk38YRbTZEuIRBa/BfAhKF2vVqB75iVbAlxK10tI",
	"VkthaC6yzhO7YSVUpPhmtidaVEueU6jsd6gUW9Sme4RRHou1EG0oHKdhankpuWE5cG3Yj0Je3BA47z96",
	"npFgNqq6aqgQt/9XIEELncTPhu9s6/dcr/3ysaNXNm6wjfYi/DbZZWugkyj7fx/85fjdSfI3nvx+lLz4",
	"7/P3H57dPnw0+PHJ7Vdf/b/uT09vv3r4l/+M7ZTHPZZl4TA/O3Xm3dkpneFtFHyA+2eL4hZCJlEmQ7er",
	"EJKS9Xq8xR6gJeIZ6GEbT3e7finNjURGuua5yLj5OHboq7iBLFrp6HFNZyN6QTm/1vcxt3GlkpKnV3TR",
	"OVkJs64Xs1QVc2/WzleqMXHnGYdCSWrL5rwUc11COr9+vOdo/AR9xSLqivKY7PVjkIcSMe/dVVHH00SI",
	"Ng/fJnKhp3UKSyEFth9fyowbPl9wLVI9rzVUX/OcyxRmK8WOmQN5yg2nAEUvrjb2VIayjB02Zb3IRcqu",
	"wvOt5fexONXl5Tuk+uXl+8E1z/A0clNFGd9OkGyEWavaJC42OR7kaANBBNmGyXbNOmUOtt1mF/t08OP6",
	"j5elTnKV8jzRhhuIL78sc1x+cGZqRoMoC4VpoyqvWVDduIAL7u9r5S66Kr7xycG1Bs1+KXj5TkjzniUu",
	"OHBSlq8Q5jni8YsTYNS62xI6juCB+UMtMB3zAmnl1ky5c2oSQT23o/wLGR0nHTYR7agPylp7C/KxhEJQ",
	"36scd/ej6RTAiFKnNusEhSq6Ko28RQIRvOniK9Qw/moKnXrkPvfGYAEsXUN6BRnF3ymCOe0M9zfCTl97",
	"mRXaPguwGUiUu0rO6gJYXWbcnWhcbvtJhBqM8ZmTb+EKtheqTX29S9bg7XTiguwJ8syYhJRIj0C1qmVX",
	"Xnygvrf57oqBAuFlyVa5WjixatjiuOELP2Zcgqy+vwfpiTFFQ4Yd/F7yKkIIy/wjJPiIhSK8T2L92PJK",
	"XhmRitKu/7BUyTedMQhkn1aP6nG17KvrgTaNqm/bOVlwHdfcgC24HyhD/eQLP5ON+3B7OUZPSx3jLnII",
	"bnm0k2xekQnhl23fyo2hFucSqGR7nHo0uhQJz+21u50T1+2dHN3KHnLC7b0kQi7y1+miGxwXOG8O13z0",
	"nmI0p/ssuCMPngo1GdtesfWFYdpk79tXuz6z26dz+xzuyfRO+djTiUuFim2HknS8Z5DDiruwPCVZOUZx",
	"qH2hgw1CPH5aLtHpZ0nsup1rrVJh7yZbXe7mALT+HjFmwxXsYAgxNg7QpngmAWavVSibcnUXJCUICoBy",
	"D5siocHfsD+O1T6fdnblXvtvqDtaIZq2zxvsNg5jKtNJVCWNmeadXsx2WcDAQYixKKqmYZRhGMvQkAMd",
	"x0lHsyZXsdgTWhVAbHjuhwX2OnsglnjIPwzC2hWs0KNtvUCUVh/W+Lye+LUykCxFpU1CDmh0edjpW03G",
	"4LfYNa5+OqRi9v2lyOLah6a9gm2SibyO77ab94dTnPZ147joenEFWzpkgKdrtqD3wngKdabHPjumtikn",
	"Oxf8yi74Fb+39R7GS9gVJ66UMr05/iBc1dMnu4QpwoAx5hju2ihJd6gX8n1OITexbO/geQa5k6gw7TOF",
	"UXd9IEyZh73L/AqwGNe8FlJ0LYGhu3MVNhHH5toEz20jKcVo28YyREJovtOdHZARCeNlKbKbnmtucR5J",
	"ZSH34A5ugPUnBjQm3nHA9tC3dcOj+XsV+FCCZZjgRLbPsmW4tiEnIYfTE/V9i7oAnv8A279iX5p3cjud",
	"fJrn3yNKi0oD+GDaRAyzN1xUPS8yYMfw14B+A/pYzolYrn5H7hxNiTODhbdnvW8a9ovyAcWlrQPcCRze",
	"kSV4WVbqmueJu20dE51KXTvRoe7+cvbzmw9pDryy8bedOFO/8l8DZ7tPyUFsFBW9EMAnR+CCCGZyrzI9",
	"4KX4bu3h+3CGHc+4C1upQDMl+8kVaKSRD0muS8G36CHauOtQAGRdJMgEic5FGg8MyIVGPpJ1Qc8VtgYY",
	"dR4x9xBiLUai47IWASzspg+4XekhGcwRJSYFbXbQbqFcialait9qYCIDabCpcslWHbcGTTufMXuHo5xy",
	"cz/+EI9n/TqEXeJvg/annN8IauzkJiR2H95hbDiSYu1dRU/AJqiNPwQhvTvc7YQzDtT5jnsZx3dOSuyt",
	"87ob4w0rTQ23BxnOViXYX+bKBxzWFtGROaJlq8ZZypeOQbH28VHrbFM2d/PwNgyV+7TmAUu3A32MnHLF",
	"bb4hz7WKgKnlhktbhQbHWRq60Rqst4+jNqqil0AaorfFQifLSv0OcR90iRsVyStzpKSMMBo9i7yw6Kvk",
	"Jp7S1hdrZbLFY5S1x6yQoJF1795GNAdxeRD0pkRZH5ri0rK1rZjTuUaNC0eY+jC38FvhcDgPVFXONwse",
	"eyuO5gLidNJer3SCaEYxP9jvgm7ywx3vBTc1TV9hn8+UULXJn0ODZozdLwL2+8OzfAapKHgej6lmRP3u",
	"A8pMrIQtO1RrCOraOEC2XpvlIlcbyF5gtaQ5W7KjaVA5y+1GJq6FFoscqMdj22PBNZ2GTaC0GYLLA2nW",
	"mro/OaD7upZZBZlZa0tYrZiSbqfI5Wmi1gswGwDJjqjf4xfsAcXrtbiGh0hFZ+NMjh+/oPQI+8dR7LBz",
	"9cV26ZWMFMv/coolzsd0YWFh4CHloM6iT7lsUchxFbZDmuzQQ2SJejqtt1+WCi75CuL3sMUenOxY2k0K",
	"9/XoIjNb0UybSm2ZMPH5wXDUTyMpUqj+LBou/79AATKKaVUgP7VFa+ykHpwtj+bqRni8fCNdjpT+HUfP",
	"2fu8fo49y2Orpius17yALlmnjNsXj/QUxb2UdQpxNpI6DNV1fJJqZIP9uenGsgdSyaRA2cketsl3Af/F",
	"Jqbrt+i0xuuufsLLbtCHmloIJRklbN0hLA900keTuK7i6+Q1TvXz21fuYChUFSsm0GpDd0hUYCoB11GJ",
	"7SeRNZZJc1x4yscMFFty4aQsI4a3b6Ld1z5SwsvShc6W9oqKurGU55ErmeBgHlpw2Uh91l1beV8BxcCA",
	"2xNa9EUpfqtBm9jLMGqwiUnkMSNZHFFAZnTGzph9SYUb23kLQ2ebKOrcvquAbAWVo25d5opnU4ZwLr45",
	"ecXsrNo9+6QXPFQQY2Vf5TVMFAlQBYUMDktZ8HWZ4mlMd6wYgtwVgYRr1obeFmvDizKW9Yk9LnwHSi29",
	"5iL3iQKk8kPazNipPW211+V2kvb1JWumc/KdrxS9dufG8HRNx1hH6VvOi/rG7iJyPHP6jevhE6d1UECv",
	"qUXWVAewzyuNYrh1xAEyg2rKFNoaG6FtEVC4hm6iaZN17cwon3jaXV5VS2n5JH4m7HgV8DFk98jZKzgf",
	"Xopi1iP8HfWBVnWVwl058pxGRd9q9QviDCrnScgubmRTqMoXd065VFKk9FIqKDvaoOwKih4SwjngUVnf",
	"RfUC7uQzIlwRfg0u+R0VR2v1eDXoCDdyVthW3FTLHfZPQ5Ur0flagdFOr0E29e+GnO8kpAZX7YFqywZa",
	"Es+Z/k3fUMtRst3IMf8tttERL1yCzJWQ9JDVLd3l4lgPhWoWGnSLhGErBdrh1H1gpN/hmNnFjTyTGdy8",
	"n/kahwTDBnIRdRtFH4I68TH1N+6RlqrYS+zLKHzb/ty5krGTnpSlmzQaGxm7g2l2L1YIqq0SsCeSGhDa",
	"HZ8OcAhmBw/tvJujIxK5B67pjglKOloHuz3yPv4b9A4tm9hntvbGPfreQMgIGq+EhLYkZ0Trp1E9TztF",
	"QjgyTqcVN+n6YEV1ATynC4yYltLGxWA+FVRvZ4kktEY/x/g2tnXARrRB06F9DcDltqkEiuwe2GkvqQSx",
	"I+Sw0BcZSs4uyiinqlf6K3Ipl6A29sXXulp9yP9DM8cONxW3InOX48VeWw6hZkKjXV8s8kgWyWnTGDzc",
	"pHS1xZb+jb1OHl+Bu++6c3aDv9yigXc2GbuQBiYf7n2ixeojd6Udf4/b0pOBcI9i3P8NqpXwkd/goblV",
	"PE2xQLpDVr7YJfkJTV54l2dJ0cXoENQn3O39jVcanJJqHMmjeds+g+RW+9og21g2TTqa/MWNy+w0nO0q",
	"qmIfJMcg2ItC+xDafgsg6mGPXQ7au0FsHow+zKYcmFYEeydB/V3zEKEffIoGK7lwEeRWRKLZJVEGOCjj",
	"pN3gSKbIxAOJreQjc6wOkr0hlSKCHd7Y72HPqw5J7WOMnnmoKrhn0gZH6B1JO8xFOHR5tA7imFrDcJ0H",
	"b0CHtiO0P4TwrV4YEndcnM3iEHGO57TjcNInliD+1cVQm3w2bdCpo+Dmje36X0cL+dlnV9ywDTAupSKJ",
	"cqFGxlmhMsiZdnVdcljxdOteSupLmXLJMlEBFUcRBRWU40xv+GoFFT2xrSiQ5AMOBC2yW7XIs31s42B8",
	"TX0jL5f/mW+Ph0Jskb2TOdHfWlro7re2zTT/qPe1qSoK6+93yB99Zdq8XKNICqHfFkHcFQ5cVFxaT2RA",
	"IYISfK8gUg1tzaWEPDraXsj8kzik4L+qEZwLIeNNfRawhOmRoV1zd4V+Sg8/UnZiOtGQ1pUwW0rG8p6J",
	"+Hs0jfy7Rn5dyfXm6tndfNrPf7g7gVba2y82fKdsOZYC3SVyHQxV3PnmhhdlDk6PfvXF4k/w9M/PsqOn",
	"j/+0+PPR86MUnj1/cXTEXzzjj188fQxP/vz82RE8Xn75YvEke/LsyeLZk2dfPn+RPn32ePHsyxd/+sJ/",
	"LsEi2n6K4H9T6YXk5M1ZcoHIthvFS/EDbO3rceROXx6Dp6S5oeAinxz7n/6HlxMUoOALb+7XibtimayN",
	"KfXxfL7ZbGbhkPmKqh4mRtXpeu7nGRY4enPWxOjtdQbJkg3AoqDTeSFMTuk11Pb2m/MLdvLmbNaqg8nx",
	"5Gh2NHtM1VJKkLwUk+PJU/qJuH5N+z5fA88NSsbtdDIvwFQi1e4vp8JnrjII/nT9ZO7DevMPLp/gdldb",
	"N6HDve0JBgSPQ+cfgr8SkYVw6enk/INPdgmabJHh+QeKGga/uyqh8w9t2d5by905xCI9vqpc252qxVEF",
	"em1/RYb2F7JCd0snN7tzluGu4KiXTQnj8Mub7/6Lfqfufe/rHU+Ojv79vQGqAfvsjpTY6dd04gCReb/m",
	"GfPXhjT3488395mkBzeoqJhVxLfTyfPPufoziaLAc0Y9g/SaIUv8LK+k2kjfE0/Nuih4tfXirTvKwhcs",
	"J93MV5qqW1bimt4DUPnU2D3tiNKhDzvcWenQ1yr+rXQ+l9L5Y3/G499K54+mdM6tUjhc6ThDyOZvzG31",
	"uNY+8k88h+8eu5bdmOZyhj57QFFlCZuHLgfEgo28oW1u3FXm03RsISSfz+VmnQ0021sHtPNc+wfY6n1q",
	"7mIN7Jf2S+W/UNYpXdVMmarYLzzPg9/og5PehJ2NfPa8eVd56DfPb2+nMbSWAD4HlnJdXSFZVPdX4F/g",
	"Whp07neHaQ1tLboljH761JbsCjWbY8HHR0dHsWcqfZxdtMtiTDnHG5XkcA35cKvHkOg9xN31ocDRr3AM",
	"30+HXmeE6/x3dZsn1aPfTew+Cr4LdqdKfmHYhgtXwT0owmO/dVII4z8parOkXN5ic3bEP0OZIMjdX6n9",
	"1CPuj1fQ9HaHstPr2mRqI8cVFz2W4rnLCqY83cbZNop5AI2mmjH/Sbh86z9yyjhlbKnadL897Gtr9Opf",
	"N9WfVkLSBCTlNItNf+dBcqn7DMZQCZ47zF7br4b09F70E4wWx7jcx4T+U3npcANk5x76Gi2dv+coCmjs",
	"2U8QJUS5odtvgOdzl8PT+9Veygc/dmtfR36dN4md0cZ+MCPWOv9gbly8Igi80ZY1Ibd375HylMPsdrON",
	"Ix3P53TzvVbazCeoeboxprDxfUPUD54FPHFv39/+/wAAAP//MTlmmg+EAAA=",
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
