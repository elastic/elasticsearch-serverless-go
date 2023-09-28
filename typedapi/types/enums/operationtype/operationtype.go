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

// Package operationtype
package operationtype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/d70d15b514ca03d715b6eb83fe5183246ded8717/specification/_global/bulk/types.ts#L83-L88
type OperationType struct {
	Name string
}

var (
	Index = OperationType{"index"}

	Create = OperationType{"create"}

	Update = OperationType{"update"}

	Delete = OperationType{"delete"}
)

func (o OperationType) MarshalText() (text []byte, err error) {
	return []byte(o.String()), nil
}

func (o *OperationType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "index":
		*o = Index
	case "create":
		*o = Create
	case "update":
		*o = Update
	case "delete":
		*o = Delete
	default:
		*o = OperationType{string(text)}
	}

	return nil
}

func (o OperationType) String() string {
	return o.Name
}
