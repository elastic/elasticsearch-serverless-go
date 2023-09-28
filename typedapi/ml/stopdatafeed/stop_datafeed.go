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

// Stops one or more datafeeds.
package stopdatafeed

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
)

const (
	datafeedidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type StopDatafeed struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	datafeedid string
}

// NewStopDatafeed type alias for index.
type NewStopDatafeed func(datafeedid string) *StopDatafeed

// NewStopDatafeedFunc returns a new instance of StopDatafeed with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewStopDatafeedFunc(tp elastictransport.Interface) NewStopDatafeed {
	return func(datafeedid string) *StopDatafeed {
		n := New(tp)

		n._datafeedid(datafeedid)

		return n
	}
}

// Stops one or more datafeeds.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-stop-datafeed.html
func New(tp elastictransport.Interface) *StopDatafeed {
	r := &StopDatafeed{
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
func (r *StopDatafeed) Raw(raw io.Reader) *StopDatafeed {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *StopDatafeed) Request(req *Request) *StopDatafeed {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *StopDatafeed) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for StopDatafeed: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == datafeedidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("datafeeds")
		path.WriteString("/")

		path.WriteString(r.datafeedid)
		path.WriteString("/")
		path.WriteString("_stop")

		method = http.MethodPost
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
func (r StopDatafeed) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the StopDatafeed query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a stopdatafeed.Response
func (r StopDatafeed) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the StopDatafeed headers map.
func (r *StopDatafeed) Header(key, value string) *StopDatafeed {
	r.headers.Set(key, value)

	return r
}

// DatafeedId Identifier for the datafeed. You can stop multiple datafeeds in a single API
// request by using a comma-separated
// list of datafeeds or a wildcard expression. You can close all datafeeds by
// using `_all` or by specifying `*` as
// the identifier.
// API Name: datafeedid
func (r *StopDatafeed) _datafeedid(datafeedid string) *StopDatafeed {
	r.paramSet |= datafeedidMask
	r.datafeedid = datafeedid

	return r
}

// AllowNoMatch Refer to the description for the `allow_no_match` query parameter.
// API name: allow_no_match
func (r *StopDatafeed) AllowNoMatch(allownomatch bool) *StopDatafeed {
	r.req.AllowNoMatch = &allownomatch

	return r
}

// Force Refer to the description for the `force` query parameter.
// API name: force
func (r *StopDatafeed) Force(force bool) *StopDatafeed {
	r.req.Force = &force

	return r
}

// Timeout Refer to the description for the `timeout` query parameter.
// API name: timeout
func (r *StopDatafeed) Timeout(duration types.Duration) *StopDatafeed {
	r.req.Timeout = duration

	return r
}