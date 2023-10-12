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

// IcuTokenizer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/0f7969a4e10ecb4423057d4ad29744c4a7c3c67b/specification/_types/analysis/icu-plugin.ts#L30-L33
type IcuTokenizer struct {
	RuleFiles string  `json:"rule_files"`
	Type      string  `json:"type,omitempty"`
	Version   *string `json:"version,omitempty"`
}

func (s *IcuTokenizer) UnmarshalJSON(data []byte) error {

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

		case "rule_files":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RuleFiles = o

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s IcuTokenizer) MarshalJSON() ([]byte, error) {
	type innerIcuTokenizer IcuTokenizer
	tmp := innerIcuTokenizer{
		RuleFiles: s.RuleFiles,
		Type:      s.Type,
		Version:   s.Version,
	}

	tmp.Type = "icu_tokenizer"

	return json.Marshal(tmp)
}

// NewIcuTokenizer returns a IcuTokenizer.
func NewIcuTokenizer() *IcuTokenizer {
	r := &IcuTokenizer{}

	return r
}
