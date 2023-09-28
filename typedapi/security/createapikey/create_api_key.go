// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/d70d15b514ca03d715b6eb83fe5183246ded8717

// Creates an API key for access without requiring basic authentication.
package createapikey

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/elasticsearch-serverless-go/typedapi/types"
	"github.com/elastic/elasticsearch-serverless-go/typedapi/types/enums/refresh"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type CreateApiKey struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int
}

// NewCreateApiKey type alias for index.
type NewCreateApiKey func() *CreateApiKey

// NewCreateApiKeyFunc returns a new instance of CreateApiKey with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewCreateApiKeyFunc(tp elastictransport.Interface) NewCreateApiKey {
	return func() *CreateApiKey {
		n := New(tp)

		return n
	}
}

// Creates an API key for access without requiring basic authentication.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-create-api-key.html
func New(tp elastictransport.Interface) *CreateApiKey {
	r := &CreateApiKey{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),

		req: NewRequest(),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *CreateApiKey) Raw(raw io.Reader) *CreateApiKey {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *CreateApiKey) Request(req *Request) *CreateApiKey {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *CreateApiKey) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if len(r.deferred) > 0 {
		for _, f := range r.deferred {
			deferredErr := f(r.req)
			if deferredErr != nil {
				return nil, deferredErr
			}
		}
	}

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for CreateApiKey: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_security")
		path.WriteString("/")
		path.WriteString("api_key")

		method = http.MethodPut
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.buf)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.buf)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Elastic-Api-Version") == "" {
		if r.buf.Len() > 0 {
			req.Header.Set("Elastic-Api-Version", "2023-10-31")
		}
	}
	if req.Header.Get("Content-Type") == "" {
		if r.buf.Len() > 0 {
			req.Header.Set("Content-Type", "application/json")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/json")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r CreateApiKey) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the CreateApiKey query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a createapikey.Response
func (r CreateApiKey) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	return nil, errorResponse
}

// Header set a key, value pair in the CreateApiKey headers map.
func (r *CreateApiKey) Header(key, value string) *CreateApiKey {
	r.headers.Set(key, value)

	return r
}

// Refresh If `true` (the default) then refresh the affected shards to make this
// operation visible to search, if `wait_for` then wait for a refresh to make
// this operation visible to search, if `false` then do nothing with refreshes.
// API name: refresh
func (r *CreateApiKey) Refresh(refresh refresh.Refresh) *CreateApiKey {
	r.values.Set("refresh", refresh.String())

	return r
}

// Expiration Expiration time for the API key. By default, API keys never expire.
// API name: expiration
func (r *CreateApiKey) Expiration(duration types.Duration) *CreateApiKey {
	r.req.Expiration = duration

	return r
}

// Metadata Arbitrary metadata that you want to associate with the API key. It supports
// nested data structure. Within the metadata object, keys beginning with `_`
// are reserved for system usage.
// API name: metadata
func (r *CreateApiKey) Metadata(metadata types.Metadata) *CreateApiKey {
	r.req.Metadata = metadata

	return r
}

// Name Specifies the name for this API key.
// API name: name
func (r *CreateApiKey) Name(name string) *CreateApiKey {
	r.req.Name = &name

	return r
}

// RoleDescriptors An array of role descriptors for this API key. This parameter is optional.
// When it is not specified or is an empty array, then the API key will have a
// point in time snapshot of permissions of the authenticated user. If you
// supply role descriptors then the resultant permissions would be an
// intersection of API keys permissions and authenticated user’s permissions
// thereby limiting the access scope for API keys. The structure of role
// descriptor is the same as the request for create role API. For more details,
// see create or update roles API.
// API name: role_descriptors
func (r *CreateApiKey) RoleDescriptors(roledescriptors map[string]types.RoleDescriptor) *CreateApiKey {

	r.req.RoleDescriptors = roledescriptors

	return r
}