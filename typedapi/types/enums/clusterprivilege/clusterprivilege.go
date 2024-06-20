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

// Package clusterprivilege
package clusterprivilege

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/ce9d54c65d2c588b2576ab24cfaaa74f7281b37d/specification/security/_types/Privileges.ts#L41-L198
type ClusterPrivilege struct {
	Name string
}

var (
	All = ClusterPrivilege{"all"}

	Canceltask = ClusterPrivilege{"cancel_task"}

	Manage = ClusterPrivilege{"manage"}

	Manageapikey = ClusterPrivilege{"manage_api_key"}

	Managebehavioralanalytics = ClusterPrivilege{"manage_behavioral_analytics"}

	Manageenrich = ClusterPrivilege{"manage_enrich"}

	Manageindextemplates = ClusterPrivilege{"manage_index_templates"}

	Manageinference = ClusterPrivilege{"manage_inference"}

	Manageingestpipelines = ClusterPrivilege{"manage_ingest_pipelines"}

	Managelogstashpipelines = ClusterPrivilege{"manage_logstash_pipelines"}

	Manageml = ClusterPrivilege{"manage_ml"}

	Manageownapikey = ClusterPrivilege{"manage_own_api_key"}

	Managepipeline = ClusterPrivilege{"manage_pipeline"}

	Managesearchapplication = ClusterPrivilege{"manage_search_application"}

	Managesearchqueryrules = ClusterPrivilege{"manage_search_query_rules"}

	Managesearchsynonyms = ClusterPrivilege{"manage_search_synonyms"}

	Managesecurity = ClusterPrivilege{"manage_security"}

	Managetransform = ClusterPrivilege{"manage_transform"}

	Monitor = ClusterPrivilege{"monitor"}

	Monitorenrich = ClusterPrivilege{"monitor_enrich"}

	Monitorinference = ClusterPrivilege{"monitor_inference"}

	Monitorml = ClusterPrivilege{"monitor_ml"}

	Monitortransform = ClusterPrivilege{"monitor_transform"}

	None = ClusterPrivilege{"none"}

	Postbehavioralanalyticsevent = ClusterPrivilege{"post_behavioral_analytics_event"}

	Readpipeline = ClusterPrivilege{"read_pipeline"}

	Readsecurity = ClusterPrivilege{"read_security"}
)

func (c ClusterPrivilege) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *ClusterPrivilege) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "all":
		*c = All
	case "cancel_task":
		*c = Canceltask
	case "manage":
		*c = Manage
	case "manage_api_key":
		*c = Manageapikey
	case "manage_behavioral_analytics":
		*c = Managebehavioralanalytics
	case "manage_enrich":
		*c = Manageenrich
	case "manage_index_templates":
		*c = Manageindextemplates
	case "manage_inference":
		*c = Manageinference
	case "manage_ingest_pipelines":
		*c = Manageingestpipelines
	case "manage_logstash_pipelines":
		*c = Managelogstashpipelines
	case "manage_ml":
		*c = Manageml
	case "manage_own_api_key":
		*c = Manageownapikey
	case "manage_pipeline":
		*c = Managepipeline
	case "manage_search_application":
		*c = Managesearchapplication
	case "manage_search_query_rules":
		*c = Managesearchqueryrules
	case "manage_search_synonyms":
		*c = Managesearchsynonyms
	case "manage_security":
		*c = Managesecurity
	case "manage_transform":
		*c = Managetransform
	case "monitor":
		*c = Monitor
	case "monitor_enrich":
		*c = Monitorenrich
	case "monitor_inference":
		*c = Monitorinference
	case "monitor_ml":
		*c = Monitorml
	case "monitor_transform":
		*c = Monitortransform
	case "none":
		*c = None
	case "post_behavioral_analytics_event":
		*c = Postbehavioralanalyticsevent
	case "read_pipeline":
		*c = Readpipeline
	case "read_security":
		*c = Readsecurity
	default:
		*c = ClusterPrivilege{string(text)}
	}

	return nil
}

func (c ClusterPrivilege) String() string {
	return c.Name
}
