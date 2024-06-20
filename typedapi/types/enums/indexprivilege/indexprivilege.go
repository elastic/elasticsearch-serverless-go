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
// https://github.com/elastic/elasticsearch-specification/tree/ce9d54c65d2c588b2576ab24cfaaa74f7281b37d

// Package indexprivilege
package indexprivilege

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/ce9d54c65d2c588b2576ab24cfaaa74f7281b37d/specification/security/_types/Privileges.ts#L285-L327
type IndexPrivilege struct {
	Name string
}

var (
	All = IndexPrivilege{"all"}

	Autoconfigure = IndexPrivilege{"auto_configure"}

	Create = IndexPrivilege{"create"}

	Createdoc = IndexPrivilege{"create_doc"}

	Createindex = IndexPrivilege{"create_index"}

	Delete = IndexPrivilege{"delete"}

	Deleteindex = IndexPrivilege{"delete_index"}

	Index = IndexPrivilege{"index"}

	Maintenance = IndexPrivilege{"maintenance"}

	Manage = IndexPrivilege{"manage"}

	Managedatastreamlifecycle = IndexPrivilege{"manage_data_stream_lifecycle"}

	Monitor = IndexPrivilege{"monitor"}

	None = IndexPrivilege{"none"}

	Read = IndexPrivilege{"read"}

	Viewindexmetadata = IndexPrivilege{"view_index_metadata"}

	Write = IndexPrivilege{"write"}
)

func (i IndexPrivilege) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

func (i *IndexPrivilege) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "all":
		*i = All
	case "auto_configure":
		*i = Autoconfigure
	case "create":
		*i = Create
	case "create_doc":
		*i = Createdoc
	case "create_index":
		*i = Createindex
	case "delete":
		*i = Delete
	case "delete_index":
		*i = Deleteindex
	case "index":
		*i = Index
	case "maintenance":
		*i = Maintenance
	case "manage":
		*i = Manage
	case "manage_data_stream_lifecycle":
		*i = Managedatastreamlifecycle
	case "monitor":
		*i = Monitor
	case "none":
		*i = None
	case "read":
		*i = Read
	case "view_index_metadata":
		*i = Viewindexmetadata
	case "write":
		*i = Write
	default:
		*i = IndexPrivilege{string(text)}
	}

	return nil
}

func (i IndexPrivilege) String() string {
	return i.Name
}
