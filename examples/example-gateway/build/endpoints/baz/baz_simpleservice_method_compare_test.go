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

package bazEndpoint

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uber/zanzibar/test/lib/test_gateway"
	"github.com/uber/zanzibar/test/lib/util"

	bazClient "github.com/uber/zanzibar/examples/example-gateway/build/clients/baz"
	clientsBazBase "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/base"
	clientsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/baz"
)

func TestCompareSuccessfulRequestOKResponse(t *testing.T) {
	testcompareCounter := 0

	gateway, err := testGateway.CreateGateway(t, map[string]interface{}{
		"clients.baz.serviceName": "bazService",
	}, &testGateway.Options{
		KnownTChannelBackends: []string{"baz"},
		TestBinary:            util.DefaultMainFile("example-gateway"),
		ConfigFiles:           util.DefaultConfigFiles("example-gateway"),
	})
	if !assert.NoError(t, err, "got bootstrap err") {
		return
	}
	defer gateway.Close()

	fakeCompare := func(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SimpleService_Compare_Args,
	) (*clientsBazBase.BazResponse, map[string]string, error) {
		testcompareCounter++

		var resHeaders map[string]string

		var res clientsBazBase.BazResponse

		clientResponse := []byte(`{"message":"different"}`)
		err := json.Unmarshal(clientResponse, &res)
		if err != nil {
			t.Fatal("cant't unmarshal client response json to client response struct")
			return nil, resHeaders, err
		}
		return &res, resHeaders, nil
	}

	gateway.TChannelBackends()["baz"].Register(
		"baz", "compare", "SimpleService::compare",
		bazClient.NewSimpleServiceCompareHandler(fakeCompare),
	)

	headers := map[string]string{}

	endpointRequest := []byte(`{"arg1":{"b1":true,"i3":42,"s2":"hello"},"arg2":{"b1":true,"i3":42,"s2":"hola"}}`)

	res, err := gateway.MakeRequest(
		"POST",
		"/baz/compare",
		headers,
		bytes.NewReader(endpointRequest),
	)
	if !assert.NoError(t, err, "got http error") {
		return
	}

	defer func() { _ = res.Body.Close() }()
	data, err := ioutil.ReadAll(res.Body)
	if !assert.NoError(t, err, "failed to read response body") {
		return
	}

	assert.Equal(t, 1, testcompareCounter)
	assert.Equal(t, 200, res.StatusCode)
	assert.JSONEq(t, `{"message":"different"}`, string(data))
}
