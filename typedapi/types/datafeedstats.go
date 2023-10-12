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

	"github.com/elastic/elasticsearch-serverless-go/typedapi/types/enums/datafeedstate"
)

// DatafeedStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/0f7969a4e10ecb4423057d4ad29744c4a7c3c67b/specification/ml/_types/Datafeed.ts#L140-L169
type DatafeedStats struct {
	// AssignmentExplanation For started datafeeds only, contains messages relating to the selection of a
	// node.
	AssignmentExplanation *string `json:"assignment_explanation,omitempty"`
	// DatafeedId A numerical character string that uniquely identifies the datafeed.
	// This identifier can contain lowercase alphanumeric characters (a-z and 0-9),
	// hyphens, and underscores.
	// It must start and end with alphanumeric characters.
	DatafeedId string `json:"datafeed_id"`
	// RunningState An object containing the running state for this datafeed.
	// It is only provided if the datafeed is started.
	RunningState *DatafeedRunningState `json:"running_state,omitempty"`
	// State The status of the datafeed, which can be one of the following values:
	// `starting`, `started`, `stopping`, `stopped`.
	State datafeedstate.DatafeedState `json:"state"`
	// TimingStats An object that provides statistical information about timing aspect of this
	// datafeed.
	TimingStats DatafeedTimingStats `json:"timing_stats"`
}

func (s *DatafeedStats) UnmarshalJSON(data []byte) error {

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

		case "assignment_explanation":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.AssignmentExplanation = &o

		case "datafeed_id":
			if err := dec.Decode(&s.DatafeedId); err != nil {
				return err
			}

		case "running_state":
			if err := dec.Decode(&s.RunningState); err != nil {
				return err
			}

		case "state":
			if err := dec.Decode(&s.State); err != nil {
				return err
			}

		case "timing_stats":
			if err := dec.Decode(&s.TimingStats); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDatafeedStats returns a DatafeedStats.
func NewDatafeedStats() *DatafeedStats {
	r := &DatafeedStats{}

	return r
}
