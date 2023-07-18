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
// https://github.com/elastic/elasticsearch-specification/tree/f16d22a4e5e7786419e984239713298b06183ba9

package queryapikeys

import (
	"github.com/elastic/elasticsearch-serverless-go/typedapi/types"
)

// Response holds the response body struct for the package queryapikeys
//
// https://github.com/elastic/elasticsearch-specification/blob/f16d22a4e5e7786419e984239713298b06183ba9/specification/security/query_api_keys/QueryApiKeysResponse.ts#L23-L38

type Response struct {

	// ApiKeys A list of API key information.
	ApiKeys []types.ApiKey `json:"api_keys"`
	// Count The number of API keys returned in the response.
	Count int `json:"count"`
	// Total The total number of API keys found.
	Total int `json:"total"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}