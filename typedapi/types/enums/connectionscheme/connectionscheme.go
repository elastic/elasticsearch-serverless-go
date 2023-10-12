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
// https://github.com/elastic/elasticsearch-specification/tree/0f7969a4e10ecb4423057d4ad29744c4a7c3c67b

// Package connectionscheme
package connectionscheme

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/0f7969a4e10ecb4423057d4ad29744c4a7c3c67b/specification/watcher/_types/Input.ts#L39-L42
type ConnectionScheme struct {
	Name string
}

var (
	Http = ConnectionScheme{"http"}

	Https = ConnectionScheme{"https"}
)

func (c ConnectionScheme) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *ConnectionScheme) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "http":
		*c = Http
	case "https":
		*c = Https
	default:
		*c = ConnectionScheme{string(text)}
	}

	return nil
}

func (c ConnectionScheme) String() string {
	return c.Name
}
