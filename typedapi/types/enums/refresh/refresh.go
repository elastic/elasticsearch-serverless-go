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

// Package refresh
package refresh

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/d70d15b514ca03d715b6eb83fe5183246ded8717/specification/_types/common.ts#L266-L273
type Refresh struct {
	Name string
}

var (
	True = Refresh{"true"}

	False = Refresh{"false"}

	Waitfor = Refresh{"wait_for"}
)

func (r *Refresh) UnmarshalJSON(data []byte) error {
	return r.UnmarshalText(data)
}

func (r Refresh) MarshalText() (text []byte, err error) {
	return []byte(r.String()), nil
}

func (r *Refresh) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "true":
		*r = True
	case "false":
		*r = False
	case "wait_for":
		*r = Waitfor
	default:
		*r = Refresh{string(text)}
	}

	return nil
}

func (r Refresh) String() string {
	return r.Name
}
