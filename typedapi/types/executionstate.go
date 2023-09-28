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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// ExecutionState type.
//
// https://github.com/elastic/elasticsearch-specification/blob/d70d15b514ca03d715b6eb83fe5183246ded8717/specification/watcher/_types/Action.ts#L120-L124
type ExecutionState struct {
	Reason     *string  `json:"reason,omitempty"`
	Successful bool     `json:"successful"`
	Timestamp  DateTime `json:"timestamp"`
}

func (s *ExecutionState) UnmarshalJSON(data []byte) error {

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

		case "reason":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Reason = &o

		case "successful":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Successful = value
			case bool:
				s.Successful = v
			}

		case "timestamp":
			if err := dec.Decode(&s.Timestamp); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewExecutionState returns a ExecutionState.
func NewExecutionState() *ExecutionState {
	r := &ExecutionState{}

	return r
}
