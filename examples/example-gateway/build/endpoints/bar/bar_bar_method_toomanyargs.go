// Code generated by zanzibar
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package barEndpoint

import (
	"context"

	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"

	clientsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/bar/bar"
	clientsFooBaseBase "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/foo/base/base"
	clientsFooFoo "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/foo/foo"
	endpointsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/bar/bar"
	endpointsFooFoo "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/foo/foo"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar/module"
)

// BarTooManyArgsHandler is the handler for "/bar/too-many-args-path"
type BarTooManyArgsHandler struct {
	Clients  *module.ClientDependencies
	endpoint *zanzibar.RouterEndpoint
}

// NewBarTooManyArgsHandler creates a handler
func NewBarTooManyArgsHandler(
	g *zanzibar.Gateway,
	deps *module.Dependencies,
) *BarTooManyArgsHandler {
	handler := &BarTooManyArgsHandler{
		Clients: deps.Client,
	}
	handler.endpoint = zanzibar.NewRouterEndpoint(
		deps.Default.Logger, deps.Default.Scope,
		"bar", "tooManyArgs",
		handler.HandleRequest,
	)
	return handler
}

// Register adds the http handler to the gateway's http router
func (h *BarTooManyArgsHandler) Register(g *zanzibar.Gateway) error {
	g.HTTPRouter.Register(
		"POST", "/bar/too-many-args-path",
		h.endpoint,
	)
	// TODO: register should return errors on route conflicts
	return nil
}

// HandleRequest handles "/bar/too-many-args-path".
func (h *BarTooManyArgsHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {
	if !req.CheckHeaders([]string{"x-uuid", "x-token"}) {
		return
	}
	var requestBody endpointsBarBar.Bar_TooManyArgs_Args
	if ok := req.ReadAndUnmarshalBody(&requestBody); !ok {
		return
	}

	workflow := TooManyArgsEndpoint{
		Clients: h.Clients,
		Logger:  req.Logger,
		Request: req,
	}

	response, cliRespHeaders, err := workflow.Handle(ctx, req.Header, &requestBody)
	if err != nil {
		switch errValue := err.(type) {

		case *endpointsBarBar.BarException:
			res.WriteJSON(
				403, cliRespHeaders, errValue,
			)
			return

		case *endpointsFooFoo.FooException:
			res.WriteJSON(
				418, cliRespHeaders, errValue,
			)
			return

		default:
			req.Logger.Warn("Workflow for endpoint returned error", zap.Error(errValue))
			res.SendErrorString(500, "Unexpected server error")
			return
		}
	}
	// TODO(sindelar): implement check headers on response
	// TODO(jakev): implement writing fields into response headers

	res.WriteJSON(200, cliRespHeaders, response)
}

// TooManyArgsEndpoint calls thrift client Bar.TooManyArgs
type TooManyArgsEndpoint struct {
	Clients *module.ClientDependencies
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w TooManyArgsEndpoint) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsBarBar.Bar_TooManyArgs_Args,
) (*endpointsBarBar.BarResponse, zanzibar.Header, error) {
	clientRequest := convertToTooManyArgsClientRequest(r)

	clientHeaders := map[string]string{}

	var ok bool
	var h string
	h, ok = reqHeaders.Get("X-Token")
	if ok {
		clientHeaders["X-Token"] = h
	}
	h, ok = reqHeaders.Get("X-Uuid")
	if ok {
		clientHeaders["X-Uuid"] = h
	}

	clientRespBody, cliRespHeaders, err := w.Clients.Bar.TooManyArgs(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		case *clientsBarBar.BarException:
			serverErr := convertTooManyArgsBarException(
				errValue,
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, nil, serverErr

		case *clientsFooFoo.FooException:
			serverErr := convertTooManyArgsFooException(
				errValue,
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, nil, serverErr

		default:
			w.Logger.Warn("Could not make client request", zap.Error(errValue))
			// TODO(sindelar): Consider returning partial headers

			return nil, nil, err

		}
	}

	// Filter and map response headers from client to server response.

	// TODO: Add support for TChannel Headers with a switch here
	resHeaders := zanzibar.ServerHTTPHeader{}

	resHeaders.Set("X-Token", cliRespHeaders["X-Token"])
	resHeaders.Set("X-Uuid", cliRespHeaders["X-Uuid"])

	response := convertTooManyArgsClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertToTooManyArgsClientRequest(in *endpointsBarBar.Bar_TooManyArgs_Args) *clientsBarBar.Bar_TooManyArgs_Args {
	out := &clientsBarBar.Bar_TooManyArgs_Args{}

	if in.Request != nil {
		out.Request = &clientsBarBar.BarRequest{}
		convertToTooManyArgsStringFieldClientRequest(in, out)
		convertToTooManyArgsBoolFieldClientRequest(in, out)
		out.Request.BinaryField = []byte(in.Request.BinaryField)
		convertToTooManyArgsTimestampClientRequest(in, out)
		convertToTooManyArgsEnumFieldClientRequest(in, out)
		convertToTooManyArgsLongFieldClientRequest(in, out)
	} else {
		out.Request = nil
	}
	if in.Foo != nil {
		out.Foo = &clientsFooFoo.FooStruct{}
		convertToTooManyArgsFooStringClientRequest(in, out)
		convertToTooManyArgsFooI32ClientRequest(in, out)
		convertToTooManyArgsFooI16ClientRequest(in, out)
		convertToTooManyArgsFooDoubleClientRequest(in, out)
		convertToTooManyArgsFooBoolClientRequest(in, out)
		out.Foo.FooMap = make(map[string]string, len(in.Foo.FooMap))
		for key1, value2 := range in.Foo.FooMap {
			out.Foo.FooMap[key1] = string(value2)
		}
		if in.Foo.Message != nil {
			out.Foo.Message = &clientsFooBaseBase.Message{}
			convertToTooManyArgsBodyClientRequest(in, out)
		} else {
			out.Foo.Message = nil
		}
	} else {
		out.Foo = nil
	}

	return out
}
func convertToTooManyArgsStringFieldClientRequest(in *endpointsBarBar.Bar_TooManyArgs_Args, out *clientsBarBar.Bar_TooManyArgs_Args) {
	out.Request.StringField = string(in.Request.StringField)
}
func convertToTooManyArgsBoolFieldClientRequest(in *endpointsBarBar.Bar_TooManyArgs_Args, out *clientsBarBar.Bar_TooManyArgs_Args) {
	out.Request.BoolField = bool(in.Request.BoolField)
}
func convertToTooManyArgsTimestampClientRequest(in *endpointsBarBar.Bar_TooManyArgs_Args, out *clientsBarBar.Bar_TooManyArgs_Args) {
	out.Request.Timestamp = clientsBarBar.Timestamp(in.Request.Timestamp)
}
func convertToTooManyArgsEnumFieldClientRequest(in *endpointsBarBar.Bar_TooManyArgs_Args, out *clientsBarBar.Bar_TooManyArgs_Args) {
	out.Request.EnumField = clientsBarBar.Fruit(in.Request.EnumField)
}
func convertToTooManyArgsLongFieldClientRequest(in *endpointsBarBar.Bar_TooManyArgs_Args, out *clientsBarBar.Bar_TooManyArgs_Args) {
	out.Request.LongField = clientsBarBar.Long(in.Request.LongField)
}
func convertToTooManyArgsFooStringClientRequest(in *endpointsBarBar.Bar_TooManyArgs_Args, out *clientsBarBar.Bar_TooManyArgs_Args) {
	out.Foo.FooString = string(in.Foo.FooString)
}
func convertToTooManyArgsFooI32ClientRequest(in *endpointsBarBar.Bar_TooManyArgs_Args, out *clientsBarBar.Bar_TooManyArgs_Args) {
	out.Foo.FooI32 = (*int32)(in.Foo.FooI32)
}
func convertToTooManyArgsFooI16ClientRequest(in *endpointsBarBar.Bar_TooManyArgs_Args, out *clientsBarBar.Bar_TooManyArgs_Args) {
	out.Foo.FooI16 = (*int16)(in.Foo.FooI16)
}
func convertToTooManyArgsFooDoubleClientRequest(in *endpointsBarBar.Bar_TooManyArgs_Args, out *clientsBarBar.Bar_TooManyArgs_Args) {
	out.Foo.FooDouble = (*float64)(in.Foo.FooDouble)
}
func convertToTooManyArgsFooBoolClientRequest(in *endpointsBarBar.Bar_TooManyArgs_Args, out *clientsBarBar.Bar_TooManyArgs_Args) {
	out.Foo.FooBool = (*bool)(in.Foo.FooBool)
}
func convertToTooManyArgsBodyClientRequest(in *endpointsBarBar.Bar_TooManyArgs_Args, out *clientsBarBar.Bar_TooManyArgs_Args) {
	out.Foo.Message.Body = string(in.Foo.Message.Body)
}

func convertTooManyArgsBarException(
	clientError *clientsBarBar.BarException,
) *endpointsBarBar.BarException {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBarBar.BarException{}
	return serverError
}
func convertTooManyArgsFooException(
	clientError *clientsFooFoo.FooException,
) *endpointsFooFoo.FooException {
	// TODO: Add error fields mapping here.
	serverError := &endpointsFooFoo.FooException{}
	return serverError
}

func convertTooManyArgsClientResponse(in *clientsBarBar.BarResponse) *endpointsBarBar.BarResponse {
	out := &endpointsBarBar.BarResponse{}

	convertToTooManyArgsStringFieldClientResponse(in, out)
	convertToTooManyArgsIntWithRangeClientResponse(in, out)
	convertToTooManyArgsIntWithoutRangeClientResponse(in, out)
	out.MapIntWithRange = make(map[endpointsBarBar.UUID]int32, len(in.MapIntWithRange))
	for key1, value2 := range in.MapIntWithRange {
		out.MapIntWithRange[endpointsBarBar.UUID(key1)] = int32(value2)
	}
	out.MapIntWithoutRange = make(map[string]int32, len(in.MapIntWithoutRange))
	for key3, value4 := range in.MapIntWithoutRange {
		out.MapIntWithoutRange[key3] = int32(value4)
	}
	out.BinaryField = []byte(in.BinaryField)

	return out
}
func convertToTooManyArgsStringFieldClientResponse(in *clientsBarBar.BarResponse, out *endpointsBarBar.BarResponse) {
	out.StringField = string(in.StringField)
}
func convertToTooManyArgsIntWithRangeClientResponse(in *clientsBarBar.BarResponse, out *endpointsBarBar.BarResponse) {
	out.IntWithRange = int32(in.IntWithRange)
}
func convertToTooManyArgsIntWithoutRangeClientResponse(in *clientsBarBar.BarResponse, out *endpointsBarBar.BarResponse) {
	out.IntWithoutRange = int32(in.IntWithoutRange)
}
