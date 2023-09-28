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

	"github.com/elastic/elasticsearch-serverless-go/typedapi/types/enums/kuromojitokenizationmode"
)

// KuromojiTokenizer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/d70d15b514ca03d715b6eb83fe5183246ded8717/specification/_types/analysis/kuromoji-plugin.ts#L58-L67
type KuromojiTokenizer struct {
	DiscardCompoundToken *bool                                             `json:"discard_compound_token,omitempty"`
	DiscardPunctuation   *bool                                             `json:"discard_punctuation,omitempty"`
	Mode                 kuromojitokenizationmode.KuromojiTokenizationMode `json:"mode"`
	NbestCost            *int                                              `json:"nbest_cost,omitempty"`
	NbestExamples        *string                                           `json:"nbest_examples,omitempty"`
	Type                 string                                            `json:"type,omitempty"`
	UserDictionary       *string                                           `json:"user_dictionary,omitempty"`
	UserDictionaryRules  []string                                          `json:"user_dictionary_rules,omitempty"`
	Version              *string                                           `json:"version,omitempty"`
}

func (s *KuromojiTokenizer) UnmarshalJSON(data []byte) error {

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

		case "discard_compound_token":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.DiscardCompoundToken = &value
			case bool:
				s.DiscardCompoundToken = &v
			}

		case "discard_punctuation":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.DiscardPunctuation = &value
			case bool:
				s.DiscardPunctuation = &v
			}

		case "mode":
			if err := dec.Decode(&s.Mode); err != nil {
				return err
			}

		case "nbest_cost":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.NbestCost = &value
			case float64:
				f := int(v)
				s.NbestCost = &f
			}

		case "nbest_examples":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.NbestExamples = &o

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		case "user_dictionary":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.UserDictionary = &o

		case "user_dictionary_rules":
			if err := dec.Decode(&s.UserDictionaryRules); err != nil {
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
func (s KuromojiTokenizer) MarshalJSON() ([]byte, error) {
	type innerKuromojiTokenizer KuromojiTokenizer
	tmp := innerKuromojiTokenizer{
		DiscardCompoundToken: s.DiscardCompoundToken,
		DiscardPunctuation:   s.DiscardPunctuation,
		Mode:                 s.Mode,
		NbestCost:            s.NbestCost,
		NbestExamples:        s.NbestExamples,
		Type:                 s.Type,
		UserDictionary:       s.UserDictionary,
		UserDictionaryRules:  s.UserDictionaryRules,
		Version:              s.Version,
	}

	tmp.Type = "kuromoji_tokenizer"

	return json.Marshal(tmp)
}

// NewKuromojiTokenizer returns a KuromojiTokenizer.
func NewKuromojiTokenizer() *KuromojiTokenizer {
	r := &KuromojiTokenizer{}

	return r
}
