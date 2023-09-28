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

// Translates SQL into Elasticsearch queries
package translate

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

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Translate struct {
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

// NewTranslate type alias for index.
type NewTranslate func() *Translate

// NewTranslateFunc returns a new instance of Translate with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewTranslateFunc(tp elastictransport.Interface) NewTranslate {
	return func() *Translate {
		n := New(tp)

		return n
	}
}

// Translates SQL into Elasticsearch queries
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/sql-translate-api.html
func New(tp elastictransport.Interface) *Translate {
	r := &Translate{
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
func (r *Translate) Raw(raw io.Reader) *Translate {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Translate) Request(req *Request) *Translate {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Translate) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Translate: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_sql")
		path.WriteString("/")
		path.WriteString("translate")

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
func (r Translate) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Translate query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a translate.Response
func (r Translate) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the Translate headers map.
func (r *Translate) Header(key, value string) *Translate {
	r.headers.Set(key, value)

	return r
}

// FetchSize The maximum number of rows (or entries) to return in one response.
// API name: fetch_size
func (r *Translate) FetchSize(fetchsize int) *Translate {
	r.req.FetchSize = &fetchsize

	return r
}

// Filter Elasticsearch query DSL for additional filtering.
// API name: filter
func (r *Translate) Filter(filter *types.Query) *Translate {

	r.req.Filter = filter

	return r
}

// Query SQL query to run.
// API name: query
func (r *Translate) Query(query string) *Translate {

	r.req.Query = query

	return r
}

// TimeZone ISO-8601 time zone ID for the search.
// API name: time_zone
func (r *Translate) TimeZone(timezone string) *Translate {
	r.req.TimeZone = &timezone

	return r
}
