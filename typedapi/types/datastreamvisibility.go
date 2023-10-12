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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// DataStreamVisibility type.
//
// https://github.com/elastic/elasticsearch-specification/blob/0f7969a4e10ecb4423057d4ad29744c4a7c3c67b/specification/indices/_types/DataStream.ts#L116-L118
type DataStreamVisibility struct {
	Hidden *bool `json:"hidden,omitempty"`
}

func (s *DataStreamVisibility) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "hidden":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Hidden = &value
			case bool:
				s.Hidden = &v
			}

		}
	}
	return nil
}

// NewDataStreamVisibility returns a DataStreamVisibility.
func NewDataStreamVisibility() *DataStreamVisibility {
	r := &DataStreamVisibility{}

	return r
}
