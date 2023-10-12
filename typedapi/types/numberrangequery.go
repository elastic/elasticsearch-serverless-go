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

	"github.com/elastic/elasticsearch-serverless-go/typedapi/types/enums/rangerelation"
)

// NumberRangeQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/0f7969a4e10ecb4423057d4ad29744c4a7c3c67b/specification/_types/query_dsl/term.ts#L145-L164
type NumberRangeQuery struct {
	// Boost Floating point number used to decrease or increase the relevance scores of
	// the query.
	// Boost values are relative to the default value of 1.0.
	// A boost value between 0 and 1.0 decreases the relevance score.
	// A value greater than 1.0 increases the relevance score.
	Boost *float32 `json:"boost,omitempty"`
	From  Float64  `json:"from,omitempty"`
	// Gt Greater than.
	Gt *Float64 `json:"gt,omitempty"`
	// Gte Greater than or equal to.
	Gte *Float64 `json:"gte,omitempty"`
	// Lt Less than.
	Lt *Float64 `json:"lt,omitempty"`
	// Lte Less than or equal to.
	Lte        *Float64 `json:"lte,omitempty"`
	QueryName_ *string  `json:"_name,omitempty"`
	// Relation Indicates how the range query matches values for `range` fields.
	Relation *rangerelation.RangeRelation `json:"relation,omitempty"`
	To       Float64                      `json:"to,omitempty"`
}

func (s *NumberRangeQuery) UnmarshalJSON(data []byte) error {

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

		case "boost":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				f := float32(value)
				s.Boost = &f
			case float64:
				f := float32(v)
				s.Boost = &f
			}

		case "from":
			if err := dec.Decode(&s.From); err != nil {
				return err
			}

		case "gt":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Gt = &f
			case float64:
				f := Float64(v)
				s.Gt = &f
			}

		case "gte":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Gte = &f
			case float64:
				f := Float64(v)
				s.Gte = &f
			}

		case "lt":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Lt = &f
			case float64:
				f := Float64(v)
				s.Lt = &f
			}

		case "lte":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Lte = &f
			case float64:
				f := Float64(v)
				s.Lte = &f
			}

		case "_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.QueryName_ = &o

		case "relation":
			if err := dec.Decode(&s.Relation); err != nil {
				return err
			}

		case "to":
			if err := dec.Decode(&s.To); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewNumberRangeQuery returns a NumberRangeQuery.
func NewNumberRangeQuery() *NumberRangeQuery {
	r := &NumberRangeQuery{}

	return r
}
