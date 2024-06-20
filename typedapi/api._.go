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

package typedapi

import (
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	async_search_delete "github.com/elastic/elasticsearch-serverless-go/typedapi/asyncsearch/delete"
	async_search_get "github.com/elastic/elasticsearch-serverless-go/typedapi/asyncsearch/get"
	async_search_status "github.com/elastic/elasticsearch-serverless-go/typedapi/asyncsearch/status"
	async_search_submit "github.com/elastic/elasticsearch-serverless-go/typedapi/asyncsearch/submit"
	cat_aliases "github.com/elastic/elasticsearch-serverless-go/typedapi/cat/aliases"
	cat_component_templates "github.com/elastic/elasticsearch-serverless-go/typedapi/cat/componenttemplates"
	cat_count "github.com/elastic/elasticsearch-serverless-go/typedapi/cat/count"
	cat_help "github.com/elastic/elasticsearch-serverless-go/typedapi/cat/help"
	cat_indices "github.com/elastic/elasticsearch-serverless-go/typedapi/cat/indices"
	cat_ml_datafeeds "github.com/elastic/elasticsearch-serverless-go/typedapi/cat/mldatafeeds"
	cat_ml_data_frame_analytics "github.com/elastic/elasticsearch-serverless-go/typedapi/cat/mldataframeanalytics"
	cat_ml_jobs "github.com/elastic/elasticsearch-serverless-go/typedapi/cat/mljobs"
	cat_ml_trained_models "github.com/elastic/elasticsearch-serverless-go/typedapi/cat/mltrainedmodels"
	cat_transforms "github.com/elastic/elasticsearch-serverless-go/typedapi/cat/transforms"
	cluster_delete_component_template "github.com/elastic/elasticsearch-serverless-go/typedapi/cluster/deletecomponenttemplate"
	cluster_exists_component_template "github.com/elastic/elasticsearch-serverless-go/typedapi/cluster/existscomponenttemplate"
	cluster_get_component_template "github.com/elastic/elasticsearch-serverless-go/typedapi/cluster/getcomponenttemplate"
	cluster_info "github.com/elastic/elasticsearch-serverless-go/typedapi/cluster/info"
	cluster_put_component_template "github.com/elastic/elasticsearch-serverless-go/typedapi/cluster/putcomponenttemplate"
	core_bulk "github.com/elastic/elasticsearch-serverless-go/typedapi/core/bulk"
	core_clear_scroll "github.com/elastic/elasticsearch-serverless-go/typedapi/core/clearscroll"
	core_close_point_in_time "github.com/elastic/elasticsearch-serverless-go/typedapi/core/closepointintime"
	core_count "github.com/elastic/elasticsearch-serverless-go/typedapi/core/count"
	core_create "github.com/elastic/elasticsearch-serverless-go/typedapi/core/create"
	core_delete "github.com/elastic/elasticsearch-serverless-go/typedapi/core/delete"
	core_delete_by_query "github.com/elastic/elasticsearch-serverless-go/typedapi/core/deletebyquery"
	core_delete_script "github.com/elastic/elasticsearch-serverless-go/typedapi/core/deletescript"
	core_exists "github.com/elastic/elasticsearch-serverless-go/typedapi/core/exists"
	core_exists_source "github.com/elastic/elasticsearch-serverless-go/typedapi/core/existssource"
	core_explain "github.com/elastic/elasticsearch-serverless-go/typedapi/core/explain"
	core_field_caps "github.com/elastic/elasticsearch-serverless-go/typedapi/core/fieldcaps"
	core_get "github.com/elastic/elasticsearch-serverless-go/typedapi/core/get"
	core_get_script "github.com/elastic/elasticsearch-serverless-go/typedapi/core/getscript"
	core_get_source "github.com/elastic/elasticsearch-serverless-go/typedapi/core/getsource"
	core_index "github.com/elastic/elasticsearch-serverless-go/typedapi/core/index"
	core_info "github.com/elastic/elasticsearch-serverless-go/typedapi/core/info"
	core_mget "github.com/elastic/elasticsearch-serverless-go/typedapi/core/mget"
	core_msearch "github.com/elastic/elasticsearch-serverless-go/typedapi/core/msearch"
	core_msearch_template "github.com/elastic/elasticsearch-serverless-go/typedapi/core/msearchtemplate"
	core_mtermvectors "github.com/elastic/elasticsearch-serverless-go/typedapi/core/mtermvectors"
	core_open_point_in_time "github.com/elastic/elasticsearch-serverless-go/typedapi/core/openpointintime"
	core_ping "github.com/elastic/elasticsearch-serverless-go/typedapi/core/ping"
	core_put_script "github.com/elastic/elasticsearch-serverless-go/typedapi/core/putscript"
	core_rank_eval "github.com/elastic/elasticsearch-serverless-go/typedapi/core/rankeval"
	core_reindex "github.com/elastic/elasticsearch-serverless-go/typedapi/core/reindex"
	core_render_search_template "github.com/elastic/elasticsearch-serverless-go/typedapi/core/rendersearchtemplate"
	core_scripts_painless_execute "github.com/elastic/elasticsearch-serverless-go/typedapi/core/scriptspainlessexecute"
	core_scroll "github.com/elastic/elasticsearch-serverless-go/typedapi/core/scroll"
	core_search "github.com/elastic/elasticsearch-serverless-go/typedapi/core/search"
	core_search_mvt "github.com/elastic/elasticsearch-serverless-go/typedapi/core/searchmvt"
	core_search_template "github.com/elastic/elasticsearch-serverless-go/typedapi/core/searchtemplate"
	core_terms_enum "github.com/elastic/elasticsearch-serverless-go/typedapi/core/termsenum"
	core_termvectors "github.com/elastic/elasticsearch-serverless-go/typedapi/core/termvectors"
	core_update "github.com/elastic/elasticsearch-serverless-go/typedapi/core/update"
	core_update_by_query "github.com/elastic/elasticsearch-serverless-go/typedapi/core/updatebyquery"
	enrich_delete_policy "github.com/elastic/elasticsearch-serverless-go/typedapi/enrich/deletepolicy"
	enrich_execute_policy "github.com/elastic/elasticsearch-serverless-go/typedapi/enrich/executepolicy"
	enrich_get_policy "github.com/elastic/elasticsearch-serverless-go/typedapi/enrich/getpolicy"
	enrich_put_policy "github.com/elastic/elasticsearch-serverless-go/typedapi/enrich/putpolicy"
	enrich_stats "github.com/elastic/elasticsearch-serverless-go/typedapi/enrich/stats"
	eql_delete "github.com/elastic/elasticsearch-serverless-go/typedapi/eql/delete"
	eql_get "github.com/elastic/elasticsearch-serverless-go/typedapi/eql/get"
	eql_get_status "github.com/elastic/elasticsearch-serverless-go/typedapi/eql/getstatus"
	eql_search "github.com/elastic/elasticsearch-serverless-go/typedapi/eql/search"
	esql_query "github.com/elastic/elasticsearch-serverless-go/typedapi/esql/query"
	graph_explore "github.com/elastic/elasticsearch-serverless-go/typedapi/graph/explore"
	indices_add_block "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/addblock"
	indices_analyze "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/analyze"
	indices_create "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/create"
	indices_create_data_stream "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/createdatastream"
	indices_data_streams_stats "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/datastreamsstats"
	indices_delete "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/delete"
	indices_delete_alias "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/deletealias"
	indices_delete_data_lifecycle "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/deletedatalifecycle"
	indices_delete_data_stream "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/deletedatastream"
	indices_delete_index_template "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/deleteindextemplate"
	indices_exists "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/exists"
	indices_exists_alias "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/existsalias"
	indices_exists_index_template "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/existsindextemplate"
	indices_explain_data_lifecycle "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/explaindatalifecycle"
	indices_get "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/get"
	indices_get_alias "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/getalias"
	indices_get_data_lifecycle "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/getdatalifecycle"
	indices_get_data_stream "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/getdatastream"
	indices_get_index_template "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/getindextemplate"
	indices_get_mapping "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/getmapping"
	indices_get_settings "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/getsettings"
	indices_migrate_to_data_stream "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/migratetodatastream"
	indices_modify_data_stream "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/modifydatastream"
	indices_put_alias "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/putalias"
	indices_put_data_lifecycle "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/putdatalifecycle"
	indices_put_index_template "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/putindextemplate"
	indices_put_mapping "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/putmapping"
	indices_put_settings "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/putsettings"
	indices_put_template "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/puttemplate"
	indices_refresh "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/refresh"
	indices_resolve_index "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/resolveindex"
	indices_rollover "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/rollover"
	indices_simulate_index_template "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/simulateindextemplate"
	indices_simulate_template "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/simulatetemplate"
	indices_update_aliases "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/updatealiases"
	indices_validate_query "github.com/elastic/elasticsearch-serverless-go/typedapi/indices/validatequery"
	inference_delete "github.com/elastic/elasticsearch-serverless-go/typedapi/inference/delete"
	inference_get "github.com/elastic/elasticsearch-serverless-go/typedapi/inference/get"
	inference_inference "github.com/elastic/elasticsearch-serverless-go/typedapi/inference/inference"
	inference_put "github.com/elastic/elasticsearch-serverless-go/typedapi/inference/put"
	ingest_delete_pipeline "github.com/elastic/elasticsearch-serverless-go/typedapi/ingest/deletepipeline"
	ingest_get_pipeline "github.com/elastic/elasticsearch-serverless-go/typedapi/ingest/getpipeline"
	ingest_processor_grok "github.com/elastic/elasticsearch-serverless-go/typedapi/ingest/processorgrok"
	ingest_put_pipeline "github.com/elastic/elasticsearch-serverless-go/typedapi/ingest/putpipeline"
	ingest_simulate "github.com/elastic/elasticsearch-serverless-go/typedapi/ingest/simulate"
	license_get "github.com/elastic/elasticsearch-serverless-go/typedapi/license/get"
	logstash_delete_pipeline "github.com/elastic/elasticsearch-serverless-go/typedapi/logstash/deletepipeline"
	logstash_get_pipeline "github.com/elastic/elasticsearch-serverless-go/typedapi/logstash/getpipeline"
	logstash_put_pipeline "github.com/elastic/elasticsearch-serverless-go/typedapi/logstash/putpipeline"
	ml_close_job "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/closejob"
	ml_delete_calendar "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/deletecalendar"
	ml_delete_calendar_event "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/deletecalendarevent"
	ml_delete_calendar_job "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/deletecalendarjob"
	ml_delete_datafeed "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/deletedatafeed"
	ml_delete_data_frame_analytics "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/deletedataframeanalytics"
	ml_delete_filter "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/deletefilter"
	ml_delete_job "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/deletejob"
	ml_delete_trained_model "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/deletetrainedmodel"
	ml_delete_trained_model_alias "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/deletetrainedmodelalias"
	ml_estimate_model_memory "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/estimatemodelmemory"
	ml_evaluate_data_frame "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/evaluatedataframe"
	ml_flush_job "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/flushjob"
	ml_get_calendar_events "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/getcalendarevents"
	ml_get_calendars "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/getcalendars"
	ml_get_datafeeds "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/getdatafeeds"
	ml_get_datafeed_stats "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/getdatafeedstats"
	ml_get_data_frame_analytics "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/getdataframeanalytics"
	ml_get_data_frame_analytics_stats "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/getdataframeanalyticsstats"
	ml_get_filters "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/getfilters"
	ml_get_jobs "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/getjobs"
	ml_get_job_stats "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/getjobstats"
	ml_get_overall_buckets "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/getoverallbuckets"
	ml_get_trained_models "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/gettrainedmodels"
	ml_get_trained_models_stats "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/gettrainedmodelsstats"
	ml_infer_trained_model "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/infertrainedmodel"
	ml_open_job "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/openjob"
	ml_post_calendar_events "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/postcalendarevents"
	ml_preview_datafeed "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/previewdatafeed"
	ml_preview_data_frame_analytics "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/previewdataframeanalytics"
	ml_put_calendar "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/putcalendar"
	ml_put_calendar_job "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/putcalendarjob"
	ml_put_datafeed "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/putdatafeed"
	ml_put_data_frame_analytics "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/putdataframeanalytics"
	ml_put_filter "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/putfilter"
	ml_put_job "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/putjob"
	ml_put_trained_model "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/puttrainedmodel"
	ml_put_trained_model_alias "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/puttrainedmodelalias"
	ml_put_trained_model_definition_part "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/puttrainedmodeldefinitionpart"
	ml_put_trained_model_vocabulary "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/puttrainedmodelvocabulary"
	ml_reset_job "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/resetjob"
	ml_start_datafeed "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/startdatafeed"
	ml_start_data_frame_analytics "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/startdataframeanalytics"
	ml_start_trained_model_deployment "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/starttrainedmodeldeployment"
	ml_stop_datafeed "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/stopdatafeed"
	ml_stop_data_frame_analytics "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/stopdataframeanalytics"
	ml_stop_trained_model_deployment "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/stoptrainedmodeldeployment"
	ml_update_datafeed "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/updatedatafeed"
	ml_update_data_frame_analytics "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/updatedataframeanalytics"
	ml_update_filter "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/updatefilter"
	ml_update_job "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/updatejob"
	ml_update_trained_model_deployment "github.com/elastic/elasticsearch-serverless-go/typedapi/ml/updatetrainedmodeldeployment"
	query_ruleset_delete "github.com/elastic/elasticsearch-serverless-go/typedapi/queryruleset/delete"
	query_ruleset_get "github.com/elastic/elasticsearch-serverless-go/typedapi/queryruleset/get"
	query_ruleset_list "github.com/elastic/elasticsearch-serverless-go/typedapi/queryruleset/list"
	query_ruleset_put "github.com/elastic/elasticsearch-serverless-go/typedapi/queryruleset/put"
	search_application_delete "github.com/elastic/elasticsearch-serverless-go/typedapi/searchapplication/delete"
	search_application_delete_behavioral_analytics "github.com/elastic/elasticsearch-serverless-go/typedapi/searchapplication/deletebehavioralanalytics"
	search_application_get "github.com/elastic/elasticsearch-serverless-go/typedapi/searchapplication/get"
	search_application_get_behavioral_analytics "github.com/elastic/elasticsearch-serverless-go/typedapi/searchapplication/getbehavioralanalytics"
	search_application_list "github.com/elastic/elasticsearch-serverless-go/typedapi/searchapplication/list"
	search_application_put "github.com/elastic/elasticsearch-serverless-go/typedapi/searchapplication/put"
	search_application_put_behavioral_analytics "github.com/elastic/elasticsearch-serverless-go/typedapi/searchapplication/putbehavioralanalytics"
	search_application_search "github.com/elastic/elasticsearch-serverless-go/typedapi/searchapplication/search"
	security_authenticate "github.com/elastic/elasticsearch-serverless-go/typedapi/security/authenticate"
	security_create_api_key "github.com/elastic/elasticsearch-serverless-go/typedapi/security/createapikey"
	security_get_api_key "github.com/elastic/elasticsearch-serverless-go/typedapi/security/getapikey"
	security_has_privileges "github.com/elastic/elasticsearch-serverless-go/typedapi/security/hasprivileges"
	security_invalidate_api_key "github.com/elastic/elasticsearch-serverless-go/typedapi/security/invalidateapikey"
	security_query_api_keys "github.com/elastic/elasticsearch-serverless-go/typedapi/security/queryapikeys"
	security_update_api_key "github.com/elastic/elasticsearch-serverless-go/typedapi/security/updateapikey"
	sql_clear_cursor "github.com/elastic/elasticsearch-serverless-go/typedapi/sql/clearcursor"
	sql_delete_async "github.com/elastic/elasticsearch-serverless-go/typedapi/sql/deleteasync"
	sql_get_async "github.com/elastic/elasticsearch-serverless-go/typedapi/sql/getasync"
	sql_get_async_status "github.com/elastic/elasticsearch-serverless-go/typedapi/sql/getasyncstatus"
	sql_query "github.com/elastic/elasticsearch-serverless-go/typedapi/sql/query"
	sql_translate "github.com/elastic/elasticsearch-serverless-go/typedapi/sql/translate"
	synonyms_delete_synonym "github.com/elastic/elasticsearch-serverless-go/typedapi/synonyms/deletesynonym"
	synonyms_delete_synonym_rule "github.com/elastic/elasticsearch-serverless-go/typedapi/synonyms/deletesynonymrule"
	synonyms_get_synonym "github.com/elastic/elasticsearch-serverless-go/typedapi/synonyms/getsynonym"
	synonyms_get_synonym_rule "github.com/elastic/elasticsearch-serverless-go/typedapi/synonyms/getsynonymrule"
	synonyms_get_synonyms_sets "github.com/elastic/elasticsearch-serverless-go/typedapi/synonyms/getsynonymssets"
	synonyms_put_synonym "github.com/elastic/elasticsearch-serverless-go/typedapi/synonyms/putsynonym"
	synonyms_put_synonym_rule "github.com/elastic/elasticsearch-serverless-go/typedapi/synonyms/putsynonymrule"
	tasks_get "github.com/elastic/elasticsearch-serverless-go/typedapi/tasks/get"
	transform_delete_transform "github.com/elastic/elasticsearch-serverless-go/typedapi/transform/deletetransform"
	transform_get_transform "github.com/elastic/elasticsearch-serverless-go/typedapi/transform/gettransform"
	transform_get_transform_stats "github.com/elastic/elasticsearch-serverless-go/typedapi/transform/gettransformstats"
	transform_preview_transform "github.com/elastic/elasticsearch-serverless-go/typedapi/transform/previewtransform"
	transform_put_transform "github.com/elastic/elasticsearch-serverless-go/typedapi/transform/puttransform"
	transform_reset_transform "github.com/elastic/elasticsearch-serverless-go/typedapi/transform/resettransform"
	transform_schedule_now_transform "github.com/elastic/elasticsearch-serverless-go/typedapi/transform/schedulenowtransform"
	transform_start_transform "github.com/elastic/elasticsearch-serverless-go/typedapi/transform/starttransform"
	transform_stop_transform "github.com/elastic/elasticsearch-serverless-go/typedapi/transform/stoptransform"
	transform_update_transform "github.com/elastic/elasticsearch-serverless-go/typedapi/transform/updatetransform"
)

type AsyncSearch struct {
	// Deletes an async search by identifier.
	// If the search is still running, the search request will be cancelled.
	// Otherwise, the saved search results are deleted.
	// If the Elasticsearch security features are enabled, the deletion of a
	// specific async search is restricted to: the authenticated user that submitted
	// the original search request; users that have the `cancel_task` cluster
	// privilege.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/async-search.html
	Delete async_search_delete.NewDelete
	// Retrieves the results of a previously submitted async search request given
	// its identifier.
	// If the Elasticsearch security features are enabled, access to the results of
	// a specific async search is restricted to the user or API key that submitted
	// it.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/async-search.html
	Get async_search_get.NewGet
	// Get async search status
	// Retrieves the status of a previously submitted async search request given its
	// identifier, without retrieving search results.
	// If the Elasticsearch security features are enabled, use of this API is
	// restricted to the `monitoring_user` role.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/async-search.html
	Status async_search_status.NewStatus
	// Runs a search request asynchronously.
	// When the primary sort of the results is an indexed field, shards get sorted
	// based on minimum and maximum value that they hold for that field, hence
	// partial results become available following the sort criteria that was
	// requested.
	// Warning: Async search does not support scroll nor search requests that only
	// include the suggest section.
	// By default, Elasticsearch doesn’t allow you to store an async search response
	// larger than 10Mb and an attempt to do this results in an error.
	// The maximum allowed size for a stored async search response can be set by
	// changing the `search.max_async_search_response_size` cluster level setting.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/async-search.html
	Submit async_search_submit.NewSubmit
}

type Cat struct {
	// Retrieves the cluster’s index aliases, including filter and routing
	// information.
	// The API does not return data stream aliases.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or the Kibana console. They are not intended for use by applications.
	// For application consumption, use the aliases API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-alias.html
	Aliases cat_aliases.NewAliases
	// Returns information about component templates in a cluster.
	// Component templates are building blocks for constructing index templates that
	// specify index mappings, settings, and aliases.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console.
	// They are not intended for use by applications. For application consumption,
	// use the get component template API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-component-templates.html
	ComponentTemplates cat_component_templates.NewComponentTemplates
	// Provides quick access to a document count for a data stream, an index, or an
	// entire cluster.
	// NOTE: The document count only includes live documents, not deleted documents
	// which have not yet been removed by the merge process.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console.
	// They are not intended for use by applications. For application consumption,
	// use the count API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-count.html
	Count cat_count.NewCount
	// Returns help for the Cat APIs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat.html
	Help cat_help.NewHelp
	// Returns high-level information about indices in a cluster, including backing
	// indices for data streams.
	// IMPORTANT: cat APIs are only intended for human consumption using the command
	// line or Kibana console.
	// They are not intended for use by applications. For application consumption,
	// use the get index API.
	// Use the cat indices API to get the following information for each index in a
	// cluster: shard count; document count; deleted document count; primary store
	// size; total store size of all shards, including shard replicas.
	// These metrics are retrieved directly from Lucene, which Elasticsearch uses
	// internally to power indexing and search. As a result, all document counts
	// include hidden nested documents.
	// To get an accurate count of Elasticsearch documents, use the cat count or
	// count APIs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-indices.html
	Indices cat_indices.NewIndices
	// Returns configuration and usage information about data frame analytics jobs.
	//
	// IMPORTANT: cat APIs are only intended for human consumption using the Kibana
	// console or command line. They are not intended for use by applications. For
	// application consumption, use the get data frame analytics jobs statistics
	// API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-dfanalytics.html
	MlDataFrameAnalytics cat_ml_data_frame_analytics.NewMlDataFrameAnalytics
	// Returns configuration and usage information about datafeeds.
	// This API returns a maximum of 10,000 datafeeds.
	// If the Elasticsearch security features are enabled, you must have
	// `monitor_ml`, `monitor`, `manage_ml`, or `manage`
	// cluster privileges to use this API.
	//
	// IMPORTANT: cat APIs are only intended for human consumption using the Kibana
	// console or command line. They are not intended for use by applications. For
	// application consumption, use the get datafeed statistics API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-datafeeds.html
	MlDatafeeds cat_ml_datafeeds.NewMlDatafeeds
	// Returns configuration and usage information for anomaly detection jobs.
	// This API returns a maximum of 10,000 jobs.
	// If the Elasticsearch security features are enabled, you must have
	// `monitor_ml`,
	// `monitor`, `manage_ml`, or `manage` cluster privileges to use this API.
	//
	// IMPORTANT: cat APIs are only intended for human consumption using the Kibana
	// console or command line. They are not intended for use by applications. For
	// application consumption, use the get anomaly detection job statistics API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-anomaly-detectors.html
	MlJobs cat_ml_jobs.NewMlJobs
	// Returns configuration and usage information about inference trained models.
	//
	// IMPORTANT: cat APIs are only intended for human consumption using the Kibana
	// console or command line. They are not intended for use by applications. For
	// application consumption, use the get trained models statistics API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-trained-model.html
	MlTrainedModels cat_ml_trained_models.NewMlTrainedModels
	// Returns configuration and usage information about transforms.
	//
	// IMPORTANT: cat APIs are only intended for human consumption using the Kibana
	// console or command line. They are not intended for use by applications. For
	// application consumption, use the get transform statistics API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-transforms.html
	Transforms cat_transforms.NewTransforms
}

type Cluster struct {
	// Deletes component templates.
	// Component templates are building blocks for constructing index templates that
	// specify index mappings, settings, and aliases.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html
	DeleteComponentTemplate cluster_delete_component_template.NewDeleteComponentTemplate
	// Returns information about whether a particular component template exist
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html
	ExistsComponentTemplate cluster_exists_component_template.NewExistsComponentTemplate
	// Retrieves information about component templates.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html
	GetComponentTemplate cluster_get_component_template.NewGetComponentTemplate
	// Returns different information about the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-info.html
	Info cluster_info.NewInfo
	// Creates or updates a component template.
	// Component templates are building blocks for constructing index templates that
	// specify index mappings, settings, and aliases.
	//
	// An index template can be composed of multiple component templates.
	// To use a component template, specify it in an index template’s `composed_of`
	// list.
	// Component templates are only applied to new data streams and indices as part
	// of a matching index template.
	//
	// Settings and mappings specified directly in the index template or the create
	// index request override any settings or mappings specified in a component
	// template.
	//
	// Component templates are only used during index creation.
	// For data streams, this includes data stream creation and the creation of a
	// stream’s backing indices.
	// Changes to component templates do not affect existing indices, including a
	// stream’s backing indices.
	//
	// You can use C-style `/* *\/` block comments in component templates.
	// You can include comments anywhere in the request body except before the
	// opening curly bracket.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html
	PutComponentTemplate cluster_put_component_template.NewPutComponentTemplate
}

type Core struct {
	// Performs multiple indexing or delete operations in a single API call.
	// This reduces overhead and can greatly increase indexing speed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-bulk.html
	Bulk core_bulk.NewBulk
	// Clears the search context and results for a scrolling search.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clear-scroll-api.html
	ClearScroll core_clear_scroll.NewClearScroll
	// Closes a point-in-time.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
	ClosePointInTime core_close_point_in_time.NewClosePointInTime
	// Returns number of documents matching a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-count.html
	Count core_count.NewCount
	// Adds a JSON document to the specified data stream or index and makes it
	// searchable.
	// If the target is an index and the document already exists, the request
	// updates the document and increments its version.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
	Create core_create.NewCreate
	// Removes a JSON document from the specified index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete.html
	Delete core_delete.NewDelete
	// Deletes documents that match the specified query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete-by-query.html
	DeleteByQuery core_delete_by_query.NewDeleteByQuery
	// Deletes a stored script or search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	DeleteScript core_delete_script.NewDeleteScript
	// Checks if a document in an index exists.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	Exists core_exists.NewExists
	// Checks if a document's `_source` is stored.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	ExistsSource core_exists_source.NewExistsSource
	// Returns information about why a specific document matches (or doesn’t match)
	// a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-explain.html
	Explain core_explain.NewExplain
	// The field capabilities API returns the information about the capabilities of
	// fields among multiple indices.
	// The field capabilities API returns runtime fields like any other field. For
	// example, a runtime field with a type
	// of keyword is returned as any other field that belongs to the `keyword`
	// family.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-field-caps.html
	FieldCaps core_field_caps.NewFieldCaps
	// Returns a document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	Get core_get.NewGet
	// Retrieves a stored script or search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	GetScript core_get_script.NewGetScript
	// Returns the source of a document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	GetSource core_get_source.NewGetSource
	// Adds a JSON document to the specified data stream or index and makes it
	// searchable.
	// If the target is an index and the document already exists, the request
	// updates the document and increments its version.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
	Index core_index.NewIndex
	// Returns basic information about the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html
	Info core_info.NewInfo
	// Allows to get multiple documents in one request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-multi-get.html
	Mget core_mget.NewMget
	// Allows to execute several search operations in one request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-multi-search.html
	Msearch core_msearch.NewMsearch
	// Runs multiple templated searches with a single request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-multi-search.html
	MsearchTemplate core_msearch_template.NewMsearchTemplate
	// Returns multiple termvectors in one request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-multi-termvectors.html
	Mtermvectors core_mtermvectors.NewMtermvectors
	// A search request by default executes against the most recent visible data of
	// the target indices,
	// which is called point in time. Elasticsearch pit (point in time) is a
	// lightweight view into the
	// state of the data as it existed when initiated. In some cases, it’s preferred
	// to perform multiple
	// search requests using the same point in time. For example, if refreshes
	// happen between
	// `search_after` requests, then the results of those requests might not be
	// consistent as changes happening
	// between searches are only visible to the more recent point in time.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
	OpenPointInTime core_open_point_in_time.NewOpenPointInTime
	// Returns whether the cluster is running.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html
	Ping core_ping.NewPing
	// Creates or updates a stored script or search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	PutScript core_put_script.NewPutScript
	// Enables you to evaluate the quality of ranked search results over a set of
	// typical search queries.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-rank-eval.html
	RankEval core_rank_eval.NewRankEval
	// Allows to copy documents from one index to another, optionally filtering the
	// source
	// documents by a query, changing the destination index settings, or fetching
	// the
	// documents from a remote cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-reindex.html
	Reindex core_reindex.NewReindex
	// Renders a search template as a search request body.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/render-search-template-api.html
	RenderSearchTemplate core_render_search_template.NewRenderSearchTemplate
	// Runs a script and returns a result.
	// https://www.elastic.co/guide/en/elasticsearch/painless/current/painless-execute-api.html
	ScriptsPainlessExecute core_scripts_painless_execute.NewScriptsPainlessExecute
	// Allows to retrieve a large numbers of results from a single search request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-body.html#request-body-search-scroll
	Scroll core_scroll.NewScroll
	// Returns search hits that match the query defined in the request.
	// You can provide search queries using the `q` query string parameter or the
	// request body.
	// If both are specified, only the query parameter is used.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-search.html
	Search core_search.NewSearch
	// Searches a vector tile for geospatial values. Returns results as a binary
	// Mapbox vector tile.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-vector-tile-api.html
	SearchMvt core_search_mvt.NewSearchMvt
	// Runs a search with a search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-template.html
	SearchTemplate core_search_template.NewSearchTemplate
	// The terms enum API  can be used to discover terms in the index that begin
	// with the provided string. It is designed for low-latency look-ups used in
	// auto-complete scenarios.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-terms-enum.html
	TermsEnum core_terms_enum.NewTermsEnum
	// Returns information and statistics about terms in the fields of a particular
	// document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-termvectors.html
	Termvectors core_termvectors.NewTermvectors
	// Updates a document with a script or partial document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update.html
	Update core_update.NewUpdate
	// Updates documents that match the specified query.
	// If no query is specified, performs an update on every document in the data
	// stream or index without modifying the source, which is useful for picking up
	// mapping changes.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update-by-query.html
	UpdateByQuery core_update_by_query.NewUpdateByQuery
}

type Enrich struct {
	// Deletes an existing enrich policy and its enrich index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-enrich-policy-api.html
	DeletePolicy enrich_delete_policy.NewDeletePolicy
	// Creates the enrich index for an existing enrich policy.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/execute-enrich-policy-api.html
	ExecutePolicy enrich_execute_policy.NewExecutePolicy
	// Returns information about an enrich policy.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-enrich-policy-api.html
	GetPolicy enrich_get_policy.NewGetPolicy
	// Creates an enrich policy.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-enrich-policy-api.html
	PutPolicy enrich_put_policy.NewPutPolicy
	// Returns enrich coordinator statistics and information about enrich policies
	// that are currently executing.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/enrich-stats-api.html
	Stats enrich_stats.NewStats
}

type Eql struct {
	// Deletes an async EQL search or a stored synchronous EQL search.
	// The API also deletes results for the search.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/eql-search-api.html
	Delete eql_delete.NewDelete
	// Returns the current status and available results for an async EQL search or a
	// stored synchronous EQL search.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-async-eql-search-api.html
	Get eql_get.NewGet
	// Returns the current status for an async EQL search or a stored synchronous
	// EQL search without returning results.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-async-eql-status-api.html
	GetStatus eql_get_status.NewGetStatus
	// Returns results matching a query expressed in Event Query Language (EQL)
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/eql-search-api.html
	Search eql_search.NewSearch
}

type Esql struct {
	// Executes an ES|QL request
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/esql-rest.html
	Query esql_query.NewQuery
}

type Graph struct {
	// Extracts and summarizes information about the documents and terms in an
	// Elasticsearch data stream or index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/graph-explore-api.html
	Explore graph_explore.NewExplore
}

type Indices struct {
	// Adds a block to an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/index-modules-blocks.html
	AddBlock indices_add_block.NewAddBlock
	// Performs analysis on a text string and returns the resulting tokens.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-analyze.html
	Analyze indices_analyze.NewAnalyze
	// Creates a new index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html
	Create indices_create.NewCreate
	// Creates a data stream.
	// You must have a matching index template with data stream enabled.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	CreateDataStream indices_create_data_stream.NewCreateDataStream
	// Retrieves statistics for one or more data streams.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	DataStreamsStats indices_data_streams_stats.NewDataStreamsStats
	// Deletes one or more indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-delete-index.html
	Delete indices_delete.NewDelete
	// Removes a data stream or index from an alias.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-aliases.html
	DeleteAlias indices_delete_alias.NewDeleteAlias
	// Removes the data lifecycle from a data stream rendering it not managed by the
	// data stream lifecycle
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams-delete-lifecycle.html
	DeleteDataLifecycle indices_delete_data_lifecycle.NewDeleteDataLifecycle
	// Deletes one or more data streams and their backing indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	DeleteDataStream indices_delete_data_stream.NewDeleteDataStream
	// The provided <index-template> may contain multiple template names separated
	// by a comma. If multiple template
	// names are specified then there is no wildcard support and the provided names
	// should match completely with
	// existing templates.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-delete-template.html
	DeleteIndexTemplate indices_delete_index_template.NewDeleteIndexTemplate
	// Checks if a data stream, index, or alias exists.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-exists.html
	Exists indices_exists.NewExists
	// Checks if an alias exists.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-aliases.html
	ExistsAlias indices_exists_alias.NewExistsAlias
	// Returns information about whether a particular index template exists.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/index-templates.html
	ExistsIndexTemplate indices_exists_index_template.NewExistsIndexTemplate
	// Retrieves information about the index's current data stream lifecycle, such
	// as any potential encountered error, time since creation etc.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams-explain-lifecycle.html
	ExplainDataLifecycle indices_explain_data_lifecycle.NewExplainDataLifecycle
	// Returns information about one or more indices. For data streams, the API
	// returns information about the
	// stream’s backing indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-index.html
	Get indices_get.NewGet
	// Retrieves information for one or more aliases.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-aliases.html
	GetAlias indices_get_alias.NewGetAlias
	// Retrieves the data stream lifecycle configuration of one or more data
	// streams.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams-get-lifecycle.html
	GetDataLifecycle indices_get_data_lifecycle.NewGetDataLifecycle
	// Retrieves information about one or more data streams.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	GetDataStream indices_get_data_stream.NewGetDataStream
	// Returns information about one or more index templates.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-template.html
	GetIndexTemplate indices_get_index_template.NewGetIndexTemplate
	// Retrieves mapping definitions for one or more indices.
	// For data streams, the API retrieves mappings for the stream’s backing
	// indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-mapping.html
	GetMapping indices_get_mapping.NewGetMapping
	// Returns setting information for one or more indices. For data streams,
	// returns setting information for the stream’s backing indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-settings.html
	GetSettings indices_get_settings.NewGetSettings
	// Converts an index alias to a data stream.
	// You must have a matching index template that is data stream enabled.
	// The alias must meet the following criteria:
	// The alias must have a write index;
	// All indices for the alias must have a `@timestamp` field mapping of a `date`
	// or `date_nanos` field type;
	// The alias must not have any filters;
	// The alias must not use custom routing.
	// If successful, the request removes the alias and creates a data stream with
	// the same name.
	// The indices for the alias become hidden backing indices for the stream.
	// The write index for the alias becomes the write index for the stream.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	MigrateToDataStream indices_migrate_to_data_stream.NewMigrateToDataStream
	// Performs one or more data stream modification actions in a single atomic
	// operation.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	ModifyDataStream indices_modify_data_stream.NewModifyDataStream
	// Adds a data stream or index to an alias.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-aliases.html
	PutAlias indices_put_alias.NewPutAlias
	// Update the data lifecycle of the specified data streams.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams-put-lifecycle.html
	PutDataLifecycle indices_put_data_lifecycle.NewPutDataLifecycle
	// Creates or updates an index template.
	// Index templates define settings, mappings, and aliases that can be applied
	// automatically to new indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-put-template.html
	PutIndexTemplate indices_put_index_template.NewPutIndexTemplate
	// Adds new fields to an existing data stream or index.
	// You can also use this API to change the search settings of existing fields.
	// For data streams, these changes are applied to all backing indices by
	// default.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-put-mapping.html
	PutMapping indices_put_mapping.NewPutMapping
	// Changes a dynamic index setting in real time. For data streams, index setting
	// changes are applied to all backing indices by default.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-update-settings.html
	PutSettings indices_put_settings.NewPutSettings
	// Creates or updates an index template.
	// Index templates define settings, mappings, and aliases that can be applied
	// automatically to new indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates-v1.html
	PutTemplate indices_put_template.NewPutTemplate
	// A refresh makes recent operations performed on one or more indices available
	// for search.
	// For data streams, the API runs the refresh operation on the stream’s backing
	// indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-refresh.html
	Refresh indices_refresh.NewRefresh
	// Resolves the specified name(s) and/or index patterns for indices, aliases,
	// and data streams.
	// Multiple patterns and remote clusters are supported.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-resolve-index-api.html
	ResolveIndex indices_resolve_index.NewResolveIndex
	// Creates a new index for a data stream or index alias.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-rollover-index.html
	Rollover indices_rollover.NewRollover
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-simulate-index.html
	SimulateIndexTemplate indices_simulate_index_template.NewSimulateIndexTemplate
	// Returns the index configuration that would be applied by a particular index
	// template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-simulate-template.html
	SimulateTemplate indices_simulate_template.NewSimulateTemplate
	// Adds a data stream or index to an alias.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-aliases.html
	UpdateAliases indices_update_aliases.NewUpdateAliases
	// Validates a potentially expensive query without executing it.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-validate.html
	ValidateQuery indices_validate_query.NewValidateQuery
}

type Inference struct {
	// Delete an inference endpoint
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-inference-api.html
	Delete inference_delete.NewDelete
	// Get an inference endpoint
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-inference-api.html
	Get inference_get.NewGet
	// Perform inference on the service
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/post-inference-api.html
	Inference inference_inference.NewInference
	// Create an inference endpoint
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-inference-api.html
	Put inference_put.NewPut
}

type Ingest struct {
	// Deletes one or more existing ingest pipeline.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-pipeline-api.html
	DeletePipeline ingest_delete_pipeline.NewDeletePipeline
	// Returns information about one or more ingest pipelines.
	// This API returns a local reference of the pipeline.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-pipeline-api.html
	GetPipeline ingest_get_pipeline.NewGetPipeline
	// Extracts structured fields out of a single text field within a document.
	// You choose which field to extract matched fields from, as well as the grok
	// pattern you expect will match.
	// A grok pattern is like a regular expression that supports aliased expressions
	// that can be reused.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/grok-processor.html
	ProcessorGrok ingest_processor_grok.NewProcessorGrok
	// Creates or updates an ingest pipeline.
	// Changes made using this API take effect immediately.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest.html
	PutPipeline ingest_put_pipeline.NewPutPipeline
	// Executes an ingest pipeline against a set of provided documents.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/simulate-pipeline-api.html
	Simulate ingest_simulate.NewSimulate
}

type License struct {
	// This API returns information about the type of license, when it was issued,
	// and when it expires, for example.
	// For more information about the different types of licenses, see
	// https://www.elastic.co/subscriptions.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-license.html
	Get license_get.NewGet
}

type Logstash struct {
	// Deletes a pipeline used for Logstash Central Management.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/logstash-api-delete-pipeline.html
	DeletePipeline logstash_delete_pipeline.NewDeletePipeline
	// Retrieves pipelines used for Logstash Central Management.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/logstash-api-get-pipeline.html
	GetPipeline logstash_get_pipeline.NewGetPipeline
	// Creates or updates a pipeline used for Logstash Central Management.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/logstash-api-put-pipeline.html
	PutPipeline logstash_put_pipeline.NewPutPipeline
}

type Ml struct {
	// Close anomaly detection jobs
	// A job can be opened and closed multiple times throughout its lifecycle. A
	// closed job cannot receive data or perform analysis operations, but you can
	// still explore and navigate results.
	// When you close a job, it runs housekeeping tasks such as pruning the model
	// history, flushing buffers, calculating final results and persisting the model
	// snapshots. Depending upon the size of the job, it could take several minutes
	// to close and the equivalent time to re-open. After it is closed, the job has
	// a minimal overhead on the cluster except for maintaining its meta data.
	// Therefore it is a best practice to close jobs that are no longer required to
	// process data.
	// If you close an anomaly detection job whose datafeed is running, the request
	// first tries to stop the datafeed. This behavior is equivalent to calling stop
	// datafeed API with the same timeout and force parameters as the close job
	// request.
	// When a datafeed that has a specified end date stops, it automatically closes
	// its associated job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-close-job.html
	CloseJob ml_close_job.NewCloseJob
	// Removes all scheduled events from a calendar, then deletes it.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-calendar.html
	DeleteCalendar ml_delete_calendar.NewDeleteCalendar
	// Deletes scheduled events from a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-calendar-event.html
	DeleteCalendarEvent ml_delete_calendar_event.NewDeleteCalendarEvent
	// Deletes anomaly detection jobs from a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-calendar-job.html
	DeleteCalendarJob ml_delete_calendar_job.NewDeleteCalendarJob
	// Deletes a data frame analytics job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-dfanalytics.html
	DeleteDataFrameAnalytics ml_delete_data_frame_analytics.NewDeleteDataFrameAnalytics
	// Deletes an existing datafeed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-datafeed.html
	DeleteDatafeed ml_delete_datafeed.NewDeleteDatafeed
	// Deletes a filter.
	// If an anomaly detection job references the filter, you cannot delete the
	// filter. You must update or delete the job before you can delete the filter.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-filter.html
	DeleteFilter ml_delete_filter.NewDeleteFilter
	// Deletes an anomaly detection job.
	//
	// All job configuration, model state and results are deleted.
	// It is not currently possible to delete multiple jobs using wildcards or a
	// comma separated list. If you delete a job that has a datafeed, the request
	// first tries to delete the datafeed. This behavior is equivalent to calling
	// the delete datafeed API with the same timeout and force parameters as the
	// delete job request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-job.html
	DeleteJob ml_delete_job.NewDeleteJob
	// Deletes an existing trained inference model that is currently not referenced
	// by an ingest pipeline.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-trained-models.html
	DeleteTrainedModel ml_delete_trained_model.NewDeleteTrainedModel
	// Deletes a trained model alias.
	// This API deletes an existing model alias that refers to a trained model. If
	// the model alias is missing or refers to a model other than the one identified
	// by the `model_id`, this API returns an error.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-trained-models-aliases.html
	DeleteTrainedModelAlias ml_delete_trained_model_alias.NewDeleteTrainedModelAlias
	// Makes an estimation of the memory usage for an anomaly detection job model.
	// It is based on analysis configuration details for the job and cardinality
	// estimates for the fields it references.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-apis.html
	EstimateModelMemory ml_estimate_model_memory.NewEstimateModelMemory
	// Evaluates the data frame analytics for an annotated index.
	// The API packages together commonly used evaluation metrics for various types
	// of machine learning features. This has been designed for use on indexes
	// created by data frame analytics. Evaluation requires both a ground truth
	// field and an analytics result field to be present.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/evaluate-dfanalytics.html
	EvaluateDataFrame ml_evaluate_data_frame.NewEvaluateDataFrame
	// Forces any buffered data to be processed by the job.
	// The flush jobs API is only applicable when sending data for analysis using
	// the post data API. Depending on the content of the buffer, then it might
	// additionally calculate new results. Both flush and close operations are
	// similar, however the flush is more efficient if you are expecting to send
	// more data for analysis. When flushing, the job remains open and is available
	// to continue analyzing data. A close operation additionally prunes and
	// persists the model state to disk and the job must be opened again before
	// analyzing further data.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-flush-job.html
	FlushJob ml_flush_job.NewFlushJob
	// Retrieves information about the scheduled events in calendars.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-calendar-event.html
	GetCalendarEvents ml_get_calendar_events.NewGetCalendarEvents
	// Retrieves configuration information for calendars.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-calendar.html
	GetCalendars ml_get_calendars.NewGetCalendars
	// Retrieves configuration information for data frame analytics jobs.
	// You can get information for multiple data frame analytics jobs in a single
	// API request by using a comma-separated list of data frame analytics jobs or a
	// wildcard expression.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-dfanalytics.html
	GetDataFrameAnalytics ml_get_data_frame_analytics.NewGetDataFrameAnalytics
	// Retrieves usage information for data frame analytics jobs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-dfanalytics-stats.html
	GetDataFrameAnalyticsStats ml_get_data_frame_analytics_stats.NewGetDataFrameAnalyticsStats
	// Retrieves usage information for datafeeds.
	// You can get statistics for multiple datafeeds in a single API request by
	// using a comma-separated list of datafeeds or a wildcard expression. You can
	// get statistics for all datafeeds by using `_all`, by specifying `*` as the
	// `<feed_id>`, or by omitting the `<feed_id>`. If the datafeed is stopped, the
	// only information you receive is the `datafeed_id` and the `state`.
	// This API returns a maximum of 10,000 datafeeds.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-datafeed-stats.html
	GetDatafeedStats ml_get_datafeed_stats.NewGetDatafeedStats
	// Retrieves configuration information for datafeeds.
	// You can get information for multiple datafeeds in a single API request by
	// using a comma-separated list of datafeeds or a wildcard expression. You can
	// get information for all datafeeds by using `_all`, by specifying `*` as the
	// `<feed_id>`, or by omitting the `<feed_id>`.
	// This API returns a maximum of 10,000 datafeeds.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-datafeed.html
	GetDatafeeds ml_get_datafeeds.NewGetDatafeeds
	// Retrieves filters.
	// You can get a single filter or all filters.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-filter.html
	GetFilters ml_get_filters.NewGetFilters
	// Retrieves usage information for anomaly detection jobs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-job-stats.html
	GetJobStats ml_get_job_stats.NewGetJobStats
	// Retrieves configuration information for anomaly detection jobs.
	// You can get information for multiple anomaly detection jobs in a single API
	// request by using a group name, a comma-separated list of jobs, or a wildcard
	// expression. You can get information for all anomaly detection jobs by using
	// `_all`, by specifying `*` as the `<job_id>`, or by omitting the `<job_id>`.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-job.html
	GetJobs ml_get_jobs.NewGetJobs
	// Retrieves overall bucket results that summarize the bucket results of
	// multiple anomaly detection jobs.
	//
	// The `overall_score` is calculated by combining the scores of all the
	// buckets within the overall bucket span. First, the maximum
	// `anomaly_score` per anomaly detection job in the overall bucket is
	// calculated. Then the `top_n` of those scores are averaged to result in
	// the `overall_score`. This means that you can fine-tune the
	// `overall_score` so that it is more or less sensitive to the number of
	// jobs that detect an anomaly at the same time. For example, if you set
	// `top_n` to `1`, the `overall_score` is the maximum bucket score in the
	// overall bucket. Alternatively, if you set `top_n` to the number of jobs,
	// the `overall_score` is high only when all jobs detect anomalies in that
	// overall bucket. If you set the `bucket_span` parameter (to a value
	// greater than its default), the `overall_score` is the maximum
	// `overall_score` of the overall buckets that have a span equal to the
	// jobs' largest bucket span.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-overall-buckets.html
	GetOverallBuckets ml_get_overall_buckets.NewGetOverallBuckets
	// Retrieves configuration information for a trained model.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-trained-models.html
	GetTrainedModels ml_get_trained_models.NewGetTrainedModels
	// Retrieves usage information for trained models. You can get usage information
	// for multiple trained
	// models in a single API request by using a comma-separated list of model IDs
	// or a wildcard expression.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-trained-models-stats.html
	GetTrainedModelsStats ml_get_trained_models_stats.NewGetTrainedModelsStats
	// Evaluates a trained model.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-trained-model.html
	InferTrainedModel ml_infer_trained_model.NewInferTrainedModel
	// Opens one or more anomaly detection jobs.
	// An anomaly detection job must be opened in order for it to be ready to
	// receive and analyze data. It can be opened and closed multiple times
	// throughout its lifecycle.
	// When you open a new job, it starts with an empty model.
	// When you open an existing job, the most recent model state is automatically
	// loaded. The job is ready to resume its analysis from where it left off, once
	// new data is received.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-open-job.html
	OpenJob ml_open_job.NewOpenJob
	// Adds scheduled events to a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-post-calendar-event.html
	PostCalendarEvents ml_post_calendar_events.NewPostCalendarEvents
	// Previews the extracted features used by a data frame analytics config.
	// http://www.elastic.co/guide/en/elasticsearch/reference/current/preview-dfanalytics.html
	PreviewDataFrameAnalytics ml_preview_data_frame_analytics.NewPreviewDataFrameAnalytics
	// Previews a datafeed.
	// This API returns the first "page" of search results from a datafeed.
	// You can preview an existing datafeed or provide configuration details for a
	// datafeed
	// and anomaly detection job in the API. The preview shows the structure of the
	// data
	// that will be passed to the anomaly detection engine.
	// IMPORTANT: When Elasticsearch security features are enabled, the preview uses
	// the credentials of the user that
	// called the API. However, when the datafeed starts it uses the roles of the
	// last user that created or updated the
	// datafeed. To get a preview that accurately reflects the behavior of the
	// datafeed, use the appropriate credentials.
	// You can also use secondary authorization headers to supply the credentials.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-preview-datafeed.html
	PreviewDatafeed ml_preview_datafeed.NewPreviewDatafeed
	// Creates a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-calendar.html
	PutCalendar ml_put_calendar.NewPutCalendar
	// Adds an anomaly detection job to a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-calendar-job.html
	PutCalendarJob ml_put_calendar_job.NewPutCalendarJob
	// Instantiates a data frame analytics job.
	// This API creates a data frame analytics job that performs an analysis on the
	// source indices and stores the outcome in a destination index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-dfanalytics.html
	PutDataFrameAnalytics ml_put_data_frame_analytics.NewPutDataFrameAnalytics
	// Instantiates a datafeed.
	// Datafeeds retrieve data from Elasticsearch for analysis by an anomaly
	// detection job.
	// You can associate only one datafeed with each anomaly detection job.
	// The datafeed contains a query that runs at a defined interval (`frequency`).
	// If you are concerned about delayed data, you can add a delay (`query_delay')
	// at each interval.
	// When Elasticsearch security features are enabled, your datafeed remembers
	// which roles the user who created it had
	// at the time of creation and runs the query using those same roles. If you
	// provide secondary authorization headers,
	// those credentials are used instead.
	// You must use Kibana, this API, or the create anomaly detection jobs API to
	// create a datafeed. Do not add a datafeed
	// directly to the `.ml-config` index. Do not give users `write` privileges on
	// the `.ml-config` index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-datafeed.html
	PutDatafeed ml_put_datafeed.NewPutDatafeed
	// Instantiates a filter.
	// A filter contains a list of strings. It can be used by one or more anomaly
	// detection jobs.
	// Specifically, filters are referenced in the `custom_rules` property of
	// detector configuration objects.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-filter.html
	PutFilter ml_put_filter.NewPutFilter
	// Instantiates an anomaly detection job. If you include a `datafeed_config`,
	// you must have read index privileges on the source index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-job.html
	PutJob ml_put_job.NewPutJob
	// Enables you to supply a trained model that is not created by data frame
	// analytics.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-models.html
	PutTrainedModel ml_put_trained_model.NewPutTrainedModel
	// Creates or updates a trained model alias. A trained model alias is a logical
	// name used to reference a single trained model.
	// You can use aliases instead of trained model identifiers to make it easier to
	// reference your models. For example, you can use aliases in inference
	// aggregations and processors.
	// An alias must be unique and refer to only a single trained model. However,
	// you can have multiple aliases for each trained model.
	// If you use this API to update an alias such that it references a different
	// trained model ID and the model uses a different type of data frame analytics,
	// an error occurs. For example, this situation occurs if you have a trained
	// model for regression analysis and a trained model for classification
	// analysis; you cannot reassign an alias from one type of trained model to
	// another.
	// If you use this API to update an alias and there are very few input fields in
	// common between the old and new trained models for the model alias, the API
	// returns a warning.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-models-aliases.html
	PutTrainedModelAlias ml_put_trained_model_alias.NewPutTrainedModelAlias
	// Creates part of a trained model definition.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-model-definition-part.html
	PutTrainedModelDefinitionPart ml_put_trained_model_definition_part.NewPutTrainedModelDefinitionPart
	// Creates a trained model vocabulary.
	// This API is supported only for natural language processing (NLP) models.
	// The vocabulary is stored in the index as described in
	// `inference_config.*.vocabulary` of the trained model definition.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-model-vocabulary.html
	PutTrainedModelVocabulary ml_put_trained_model_vocabulary.NewPutTrainedModelVocabulary
	// Resets an anomaly detection job.
	// All model state and results are deleted. The job is ready to start over as if
	// it had just been created.
	// It is not currently possible to reset multiple jobs using wildcards or a
	// comma separated list.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-reset-job.html
	ResetJob ml_reset_job.NewResetJob
	// Starts a data frame analytics job.
	// A data frame analytics job can be started and stopped multiple times
	// throughout its lifecycle.
	// If the destination index does not exist, it is created automatically the
	// first time you start the data frame analytics job. The
	// `index.number_of_shards` and `index.number_of_replicas` settings for the
	// destination index are copied from the source index. If there are multiple
	// source indices, the destination index copies the highest setting values. The
	// mappings for the destination index are also copied from the source indices.
	// If there are any mapping conflicts, the job fails to start.
	// If the destination index exists, it is used as is. You can therefore set up
	// the destination index in advance with custom settings and mappings.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/start-dfanalytics.html
	StartDataFrameAnalytics ml_start_data_frame_analytics.NewStartDataFrameAnalytics
	// Starts one or more datafeeds.
	//
	// A datafeed must be started in order to retrieve data from Elasticsearch. A
	// datafeed can be started and stopped
	// multiple times throughout its lifecycle.
	//
	// Before you can start a datafeed, the anomaly detection job must be open.
	// Otherwise, an error occurs.
	//
	// If you restart a stopped datafeed, it continues processing input data from
	// the next millisecond after it was stopped.
	// If new data was indexed for that exact millisecond between stopping and
	// starting, it will be ignored.
	//
	// When Elasticsearch security features are enabled, your datafeed remembers
	// which roles the last user to create or
	// update it had at the time of creation or update and runs the query using
	// those same roles. If you provided secondary
	// authorization headers when you created or updated the datafeed, those
	// credentials are used instead.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-start-datafeed.html
	StartDatafeed ml_start_datafeed.NewStartDatafeed
	// Starts a trained model deployment, which allocates the model to every machine
	// learning node.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/start-trained-model-deployment.html
	StartTrainedModelDeployment ml_start_trained_model_deployment.NewStartTrainedModelDeployment
	// Stops one or more data frame analytics jobs.
	// A data frame analytics job can be started and stopped multiple times
	// throughout its lifecycle.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/stop-dfanalytics.html
	StopDataFrameAnalytics ml_stop_data_frame_analytics.NewStopDataFrameAnalytics
	// Stops one or more datafeeds.
	// A datafeed that is stopped ceases to retrieve data from Elasticsearch. A
	// datafeed can be started and stopped
	// multiple times throughout its lifecycle.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-stop-datafeed.html
	StopDatafeed ml_stop_datafeed.NewStopDatafeed
	// Stops a trained model deployment.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/stop-trained-model-deployment.html
	StopTrainedModelDeployment ml_stop_trained_model_deployment.NewStopTrainedModelDeployment
	// Updates an existing data frame analytics job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-dfanalytics.html
	UpdateDataFrameAnalytics ml_update_data_frame_analytics.NewUpdateDataFrameAnalytics
	// Updates the properties of a datafeed.
	// You must stop and start the datafeed for the changes to be applied.
	// When Elasticsearch security features are enabled, your datafeed remembers
	// which roles the user who updated it had at
	// the time of the update and runs the query using those same roles. If you
	// provide secondary authorization headers,
	// those credentials are used instead.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-update-datafeed.html
	UpdateDatafeed ml_update_datafeed.NewUpdateDatafeed
	// Updates the description of a filter, adds items, or removes items from the
	// list.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-update-filter.html
	UpdateFilter ml_update_filter.NewUpdateFilter
	// Updates certain properties of an anomaly detection job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-update-job.html
	UpdateJob ml_update_job.NewUpdateJob
	// Starts a trained model deployment, which allocates the model to every machine
	// learning node.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-trained-model-deployment.html
	UpdateTrainedModelDeployment ml_update_trained_model_deployment.NewUpdateTrainedModelDeployment
}

type QueryRuleset struct {
	// Deletes a query ruleset.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-query-ruleset.html
	Delete query_ruleset_delete.NewDelete
	// Returns the details about a query ruleset
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-query-ruleset.html
	Get query_ruleset_get.NewGet
	// Returns summarized information about existing query rulesets.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/list-query-rulesets.html
	List query_ruleset_list.NewList
	// Creates or updates a query ruleset.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-query-ruleset.html
	Put query_ruleset_put.NewPut
}

type SearchApplication struct {
	// Deletes a search application.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-search-application.html
	Delete search_application_delete.NewDelete
	// Delete a behavioral analytics collection.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-analytics-collection.html
	DeleteBehavioralAnalytics search_application_delete_behavioral_analytics.NewDeleteBehavioralAnalytics
	// Returns the details about a search application
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-search-application.html
	Get search_application_get.NewGet
	// Returns the existing behavioral analytics collections.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/list-analytics-collection.html
	GetBehavioralAnalytics search_application_get_behavioral_analytics.NewGetBehavioralAnalytics
	// Returns the existing search applications.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/list-search-applications.html
	List search_application_list.NewList
	// Creates or updates a search application.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-search-application.html
	Put search_application_put.NewPut
	// Creates a behavioral analytics collection.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-analytics-collection.html
	PutBehavioralAnalytics search_application_put_behavioral_analytics.NewPutBehavioralAnalytics
	// Perform a search against a search application.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-application-search.html
	Search search_application_search.NewSearch
}

type Security struct {
	// Enables you to submit a request with a basic auth header to authenticate a
	// user and retrieve information about the authenticated user.
	// A successful call returns a JSON structure that shows user information such
	// as their username, the roles that are assigned to the user, any assigned
	// metadata, and information about the realms that authenticated and authorized
	// the user.
	// If the user cannot be authenticated, this API returns a 401 status code.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-authenticate.html
	Authenticate security_authenticate.NewAuthenticate
	// Creates an API key for access without requiring basic authentication.
	// A successful request returns a JSON structure that contains the API key, its
	// unique id, and its name.
	// If applicable, it also returns expiration information for the API key in
	// milliseconds.
	// NOTE: By default, API keys never expire. You can specify expiration
	// information when you create the API keys.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-create-api-key.html
	CreateApiKey security_create_api_key.NewCreateApiKey
	// Retrieves information for one or more API keys.
	// NOTE: If you have only the `manage_own_api_key` privilege, this API returns
	// only the API keys that you own.
	// If you have `read_security`, `manage_api_key` or greater privileges
	// (including `manage_security`), this API returns all API keys regardless of
	// ownership.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-api-key.html
	GetApiKey security_get_api_key.NewGetApiKey
	// Determines whether the specified user has a specified list of privileges.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-has-privileges.html
	HasPrivileges security_has_privileges.NewHasPrivileges
	// Invalidates one or more API keys.
	// The `manage_api_key` privilege allows deleting any API keys.
	// The `manage_own_api_key` only allows deleting API keys that are owned by the
	// user.
	// In addition, with the `manage_own_api_key` privilege, an invalidation request
	// must be issued in one of the three formats:
	// - Set the parameter `owner=true`.
	// - Or, set both `username` and `realm_name` to match the user’s identity.
	// - Or, if the request is issued by an API key, i.e. an API key invalidates
	// itself, specify its ID in the `ids` field.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-invalidate-api-key.html
	InvalidateApiKey security_invalidate_api_key.NewInvalidateApiKey
	// Retrieves information for API keys in a paginated manner. You can optionally
	// filter the results with a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-query-api-key.html
	QueryApiKeys security_query_api_keys.NewQueryApiKeys
	// Updates attributes of an existing API key.
	// Users can only update API keys that they created or that were granted to
	// them.
	// Use this API to update API keys created by the create API Key or grant API
	// Key APIs.
	// If you need to apply the same update to many API keys, you can use bulk
	// update API Keys to reduce overhead.
	// It’s not possible to update expired API keys, or API keys that have been
	// invalidated by invalidate API Key.
	// This API supports updates to an API key’s access scope and metadata.
	// The access scope of an API key is derived from the `role_descriptors` you
	// specify in the request, and a snapshot of the owner user’s permissions at the
	// time of the request.
	// The snapshot of the owner’s permissions is updated automatically on every
	// call.
	// If you don’t specify `role_descriptors` in the request, a call to this API
	// might still change the API key’s access scope.
	// This change can occur if the owner user’s permissions have changed since the
	// API key was created or last modified.
	// To update another user’s API key, use the `run_as` feature to submit a
	// request on behalf of another user.
	// IMPORTANT: It’s not possible to use an API key as the authentication
	// credential for this API.
	// To update an API key, the owner user’s credentials are required.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-update-api-key.html
	UpdateApiKey security_update_api_key.NewUpdateApiKey
}

type Sql struct {
	// Clears the SQL cursor
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clear-sql-cursor-api.html
	ClearCursor sql_clear_cursor.NewClearCursor
	// Deletes an async SQL search or a stored synchronous SQL search. If the search
	// is still running, the API cancels it.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-async-sql-search-api.html
	DeleteAsync sql_delete_async.NewDeleteAsync
	// Returns the current status and available results for an async SQL search or
	// stored synchronous SQL search
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-async-sql-search-api.html
	GetAsync sql_get_async.NewGetAsync
	// Returns the current status of an async SQL search or a stored synchronous SQL
	// search
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-async-sql-search-status-api.html
	GetAsyncStatus sql_get_async_status.NewGetAsyncStatus
	// Executes a SQL request
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/sql-search-api.html
	Query sql_query.NewQuery
	// Translates SQL into Elasticsearch queries
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/sql-translate-api.html
	Translate sql_translate.NewTranslate
}

type Synonyms struct {
	// Deletes a synonym set
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-synonyms-set.html
	DeleteSynonym synonyms_delete_synonym.NewDeleteSynonym
	// Deletes a synonym rule in a synonym set
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-synonym-rule.html
	DeleteSynonymRule synonyms_delete_synonym_rule.NewDeleteSynonymRule
	// Retrieves a synonym set
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-synonyms-set.html
	GetSynonym synonyms_get_synonym.NewGetSynonym
	// Retrieves a synonym rule from a synonym set
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-synonym-rule.html
	GetSynonymRule synonyms_get_synonym_rule.NewGetSynonymRule
	// Retrieves a summary of all defined synonym sets
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/list-synonyms-sets.html
	GetSynonymsSets synonyms_get_synonyms_sets.NewGetSynonymsSets
	// Creates or updates a synonym set.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-synonyms-set.html
	PutSynonym synonyms_put_synonym.NewPutSynonym
	// Creates or updates a synonym rule in a synonym set
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-synonym-rule.html
	PutSynonymRule synonyms_put_synonym_rule.NewPutSynonymRule
}

type Tasks struct {
	// Returns information about a task.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/tasks.html
	Get tasks_get.NewGet
}

type Transform struct {
	// Deletes a transform.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-transform.html
	DeleteTransform transform_delete_transform.NewDeleteTransform
	// Retrieves configuration information for transforms.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-transform.html
	GetTransform transform_get_transform.NewGetTransform
	// Retrieves usage information for transforms.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-transform-stats.html
	GetTransformStats transform_get_transform_stats.NewGetTransformStats
	// Previews a transform.
	//
	// It returns a maximum of 100 results. The calculations are based on all the
	// current data in the source index. It also
	// generates a list of mappings and settings for the destination index. These
	// values are determined based on the field
	// types of the source index and the transform aggregations.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/preview-transform.html
	PreviewTransform transform_preview_transform.NewPreviewTransform
	// Creates a transform.
	//
	// A transform copies data from source indices, transforms it, and persists it
	// into an entity-centric destination index. You can also think of the
	// destination index as a two-dimensional tabular data structure (known as
	// a data frame). The ID for each document in the data frame is generated from a
	// hash of the entity, so there is a
	// unique row per entity.
	//
	// You must choose either the latest or pivot method for your transform; you
	// cannot use both in a single transform. If
	// you choose to use the pivot method for your transform, the entities are
	// defined by the set of `group_by` fields in
	// the pivot object. If you choose to use the latest method, the entities are
	// defined by the `unique_key` field values
	// in the latest object.
	//
	// You must have `create_index`, `index`, and `read` privileges on the
	// destination index and `read` and
	// `view_index_metadata` privileges on the source indices. When Elasticsearch
	// security features are enabled, the
	// transform remembers which roles the user that created it had at the time of
	// creation and uses those same roles. If
	// those roles do not have the required privileges on the source and destination
	// indices, the transform fails when it
	// attempts unauthorized operations.
	//
	// NOTE: You must use Kibana or this API to create a transform. Do not add a
	// transform directly into any
	// `.transform-internal*` indices using the Elasticsearch index API. If
	// Elasticsearch security features are enabled, do
	// not give users any privileges on `.transform-internal*` indices. If you used
	// transforms prior to 7.5, also do not
	// give users any privileges on `.data-frame-internal*` indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-transform.html
	PutTransform transform_put_transform.NewPutTransform
	// Resets a transform.
	// Before you can reset it, you must stop it; alternatively, use the `force`
	// query parameter.
	// If the destination index was created by the transform, it is deleted.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/reset-transform.html
	ResetTransform transform_reset_transform.NewResetTransform
	// Schedules now a transform.
	//
	// If you _schedule_now a transform, it will process the new data instantly,
	// without waiting for the configured frequency interval. After _schedule_now
	// API is called,
	// the transform will be processed again at now + frequency unless _schedule_now
	// API
	// is called again in the meantime.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/schedule-now-transform.html
	ScheduleNowTransform transform_schedule_now_transform.NewScheduleNowTransform
	// Starts a transform.
	//
	// When you start a transform, it creates the destination index if it does not
	// already exist. The `number_of_shards` is
	// set to `1` and the `auto_expand_replicas` is set to `0-1`. If it is a pivot
	// transform, it deduces the mapping
	// definitions for the destination index from the source indices and the
	// transform aggregations. If fields in the
	// destination index are derived from scripts (as in the case of
	// `scripted_metric` or `bucket_script` aggregations),
	// the transform uses dynamic mappings unless an index template exists. If it is
	// a latest transform, it does not deduce
	// mapping definitions; it uses dynamic mappings. To use explicit mappings,
	// create the destination index before you
	// start the transform. Alternatively, you can create an index template, though
	// it does not affect the deduced mappings
	// in a pivot transform.
	//
	// When the transform starts, a series of validations occur to ensure its
	// success. If you deferred validation when you
	// created the transform, they occur when you start the transform—​with the
	// exception of privilege checks. When
	// Elasticsearch security features are enabled, the transform remembers which
	// roles the user that created it had at the
	// time of creation and uses those same roles. If those roles do not have the
	// required privileges on the source and
	// destination indices, the transform fails when it attempts unauthorized
	// operations.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/start-transform.html
	StartTransform transform_start_transform.NewStartTransform
	// Stops one or more transforms.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/stop-transform.html
	StopTransform transform_stop_transform.NewStopTransform
	// Updates certain properties of a transform.
	//
	// All updated properties except `description` do not take effect until after
	// the transform starts the next checkpoint,
	// thus there is data consistency in each checkpoint. To use this API, you must
	// have `read` and `view_index_metadata`
	// privileges for the source indices. You must also have `index` and `read`
	// privileges for the destination index. When
	// Elasticsearch security features are enabled, the transform remembers which
	// roles the user who updated it had at the
	// time of update and runs with those privileges.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-transform.html
	UpdateTransform transform_update_transform.NewUpdateTransform
}

type API struct {
	AsyncSearch       AsyncSearch
	Cat               Cat
	Cluster           Cluster
	Core              Core
	Enrich            Enrich
	Eql               Eql
	Esql              Esql
	Graph             Graph
	Indices           Indices
	Inference         Inference
	Ingest            Ingest
	License           License
	Logstash          Logstash
	Ml                Ml
	QueryRuleset      QueryRuleset
	SearchApplication SearchApplication
	Security          Security
	Sql               Sql
	Synonyms          Synonyms
	Tasks             Tasks
	Transform         Transform

	// Performs multiple indexing or delete operations in a single API call.
	// This reduces overhead and can greatly increase indexing speed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-bulk.html
	Bulk core_bulk.NewBulk
	// Clears the search context and results for a scrolling search.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clear-scroll-api.html
	ClearScroll core_clear_scroll.NewClearScroll
	// Closes a point-in-time.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
	ClosePointInTime core_close_point_in_time.NewClosePointInTime
	// Returns number of documents matching a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-count.html
	Count core_count.NewCount
	// Adds a JSON document to the specified data stream or index and makes it
	// searchable.
	// If the target is an index and the document already exists, the request
	// updates the document and increments its version.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
	Create core_create.NewCreate
	// Removes a JSON document from the specified index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete.html
	Delete core_delete.NewDelete
	// Deletes documents that match the specified query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete-by-query.html
	DeleteByQuery core_delete_by_query.NewDeleteByQuery
	// Deletes a stored script or search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	DeleteScript core_delete_script.NewDeleteScript
	// Checks if a document in an index exists.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	Exists core_exists.NewExists
	// Checks if a document's `_source` is stored.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	ExistsSource core_exists_source.NewExistsSource
	// Returns information about why a specific document matches (or doesn’t match)
	// a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-explain.html
	Explain core_explain.NewExplain
	// The field capabilities API returns the information about the capabilities of
	// fields among multiple indices.
	// The field capabilities API returns runtime fields like any other field. For
	// example, a runtime field with a type
	// of keyword is returned as any other field that belongs to the `keyword`
	// family.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-field-caps.html
	FieldCaps core_field_caps.NewFieldCaps
	// Returns a document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	Get core_get.NewGet
	// Retrieves a stored script or search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	GetScript core_get_script.NewGetScript
	// Returns the source of a document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	GetSource core_get_source.NewGetSource
	// Adds a JSON document to the specified data stream or index and makes it
	// searchable.
	// If the target is an index and the document already exists, the request
	// updates the document and increments its version.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
	Index core_index.NewIndex
	// Returns basic information about the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html
	Info core_info.NewInfo
	// Allows to get multiple documents in one request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-multi-get.html
	Mget core_mget.NewMget
	// Allows to execute several search operations in one request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-multi-search.html
	Msearch core_msearch.NewMsearch
	// Runs multiple templated searches with a single request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-multi-search.html
	MsearchTemplate core_msearch_template.NewMsearchTemplate
	// Returns multiple termvectors in one request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-multi-termvectors.html
	Mtermvectors core_mtermvectors.NewMtermvectors
	// A search request by default executes against the most recent visible data of
	// the target indices,
	// which is called point in time. Elasticsearch pit (point in time) is a
	// lightweight view into the
	// state of the data as it existed when initiated. In some cases, it’s preferred
	// to perform multiple
	// search requests using the same point in time. For example, if refreshes
	// happen between
	// `search_after` requests, then the results of those requests might not be
	// consistent as changes happening
	// between searches are only visible to the more recent point in time.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
	OpenPointInTime core_open_point_in_time.NewOpenPointInTime
	// Returns whether the cluster is running.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html
	Ping core_ping.NewPing
	// Creates or updates a stored script or search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	PutScript core_put_script.NewPutScript
	// Enables you to evaluate the quality of ranked search results over a set of
	// typical search queries.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-rank-eval.html
	RankEval core_rank_eval.NewRankEval
	// Allows to copy documents from one index to another, optionally filtering the
	// source
	// documents by a query, changing the destination index settings, or fetching
	// the
	// documents from a remote cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-reindex.html
	Reindex core_reindex.NewReindex
	// Renders a search template as a search request body.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/render-search-template-api.html
	RenderSearchTemplate core_render_search_template.NewRenderSearchTemplate
	// Runs a script and returns a result.
	// https://www.elastic.co/guide/en/elasticsearch/painless/current/painless-execute-api.html
	ScriptsPainlessExecute core_scripts_painless_execute.NewScriptsPainlessExecute
	// Allows to retrieve a large numbers of results from a single search request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-body.html#request-body-search-scroll
	Scroll core_scroll.NewScroll
	// Returns search hits that match the query defined in the request.
	// You can provide search queries using the `q` query string parameter or the
	// request body.
	// If both are specified, only the query parameter is used.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-search.html
	Search core_search.NewSearch
	// Searches a vector tile for geospatial values. Returns results as a binary
	// Mapbox vector tile.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-vector-tile-api.html
	SearchMvt core_search_mvt.NewSearchMvt
	// Runs a search with a search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-template.html
	SearchTemplate core_search_template.NewSearchTemplate
	// The terms enum API  can be used to discover terms in the index that begin
	// with the provided string. It is designed for low-latency look-ups used in
	// auto-complete scenarios.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-terms-enum.html
	TermsEnum core_terms_enum.NewTermsEnum
	// Returns information and statistics about terms in the fields of a particular
	// document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-termvectors.html
	Termvectors core_termvectors.NewTermvectors
	// Updates a document with a script or partial document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update.html
	Update core_update.NewUpdate
	// Updates documents that match the specified query.
	// If no query is specified, performs an update on every document in the data
	// stream or index without modifying the source, which is useful for picking up
	// mapping changes.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update-by-query.html
	UpdateByQuery core_update_by_query.NewUpdateByQuery
}

func New(tp elastictransport.Interface) *API {
	return &API{
		// AsyncSearch
		AsyncSearch: AsyncSearch{
			Delete: async_search_delete.NewDeleteFunc(tp),
			Get:    async_search_get.NewGetFunc(tp),
			Status: async_search_status.NewStatusFunc(tp),
			Submit: async_search_submit.NewSubmitFunc(tp),
		},

		// Cat
		Cat: Cat{
			Aliases:              cat_aliases.NewAliasesFunc(tp),
			ComponentTemplates:   cat_component_templates.NewComponentTemplatesFunc(tp),
			Count:                cat_count.NewCountFunc(tp),
			Help:                 cat_help.NewHelpFunc(tp),
			Indices:              cat_indices.NewIndicesFunc(tp),
			MlDataFrameAnalytics: cat_ml_data_frame_analytics.NewMlDataFrameAnalyticsFunc(tp),
			MlDatafeeds:          cat_ml_datafeeds.NewMlDatafeedsFunc(tp),
			MlJobs:               cat_ml_jobs.NewMlJobsFunc(tp),
			MlTrainedModels:      cat_ml_trained_models.NewMlTrainedModelsFunc(tp),
			Transforms:           cat_transforms.NewTransformsFunc(tp),
		},

		// Cluster
		Cluster: Cluster{
			DeleteComponentTemplate: cluster_delete_component_template.NewDeleteComponentTemplateFunc(tp),
			ExistsComponentTemplate: cluster_exists_component_template.NewExistsComponentTemplateFunc(tp),
			GetComponentTemplate:    cluster_get_component_template.NewGetComponentTemplateFunc(tp),
			Info:                    cluster_info.NewInfoFunc(tp),
			PutComponentTemplate:    cluster_put_component_template.NewPutComponentTemplateFunc(tp),
		},

		// Core
		Core: Core{
			Bulk:                   core_bulk.NewBulkFunc(tp),
			ClearScroll:            core_clear_scroll.NewClearScrollFunc(tp),
			ClosePointInTime:       core_close_point_in_time.NewClosePointInTimeFunc(tp),
			Count:                  core_count.NewCountFunc(tp),
			Create:                 core_create.NewCreateFunc(tp),
			Delete:                 core_delete.NewDeleteFunc(tp),
			DeleteByQuery:          core_delete_by_query.NewDeleteByQueryFunc(tp),
			DeleteScript:           core_delete_script.NewDeleteScriptFunc(tp),
			Exists:                 core_exists.NewExistsFunc(tp),
			ExistsSource:           core_exists_source.NewExistsSourceFunc(tp),
			Explain:                core_explain.NewExplainFunc(tp),
			FieldCaps:              core_field_caps.NewFieldCapsFunc(tp),
			Get:                    core_get.NewGetFunc(tp),
			GetScript:              core_get_script.NewGetScriptFunc(tp),
			GetSource:              core_get_source.NewGetSourceFunc(tp),
			Index:                  core_index.NewIndexFunc(tp),
			Info:                   core_info.NewInfoFunc(tp),
			Mget:                   core_mget.NewMgetFunc(tp),
			Msearch:                core_msearch.NewMsearchFunc(tp),
			MsearchTemplate:        core_msearch_template.NewMsearchTemplateFunc(tp),
			Mtermvectors:           core_mtermvectors.NewMtermvectorsFunc(tp),
			OpenPointInTime:        core_open_point_in_time.NewOpenPointInTimeFunc(tp),
			Ping:                   core_ping.NewPingFunc(tp),
			PutScript:              core_put_script.NewPutScriptFunc(tp),
			RankEval:               core_rank_eval.NewRankEvalFunc(tp),
			Reindex:                core_reindex.NewReindexFunc(tp),
			RenderSearchTemplate:   core_render_search_template.NewRenderSearchTemplateFunc(tp),
			ScriptsPainlessExecute: core_scripts_painless_execute.NewScriptsPainlessExecuteFunc(tp),
			Scroll:                 core_scroll.NewScrollFunc(tp),
			Search:                 core_search.NewSearchFunc(tp),
			SearchMvt:              core_search_mvt.NewSearchMvtFunc(tp),
			SearchTemplate:         core_search_template.NewSearchTemplateFunc(tp),
			TermsEnum:              core_terms_enum.NewTermsEnumFunc(tp),
			Termvectors:            core_termvectors.NewTermvectorsFunc(tp),
			Update:                 core_update.NewUpdateFunc(tp),
			UpdateByQuery:          core_update_by_query.NewUpdateByQueryFunc(tp),
		},

		// Enrich
		Enrich: Enrich{
			DeletePolicy:  enrich_delete_policy.NewDeletePolicyFunc(tp),
			ExecutePolicy: enrich_execute_policy.NewExecutePolicyFunc(tp),
			GetPolicy:     enrich_get_policy.NewGetPolicyFunc(tp),
			PutPolicy:     enrich_put_policy.NewPutPolicyFunc(tp),
			Stats:         enrich_stats.NewStatsFunc(tp),
		},

		// Eql
		Eql: Eql{
			Delete:    eql_delete.NewDeleteFunc(tp),
			Get:       eql_get.NewGetFunc(tp),
			GetStatus: eql_get_status.NewGetStatusFunc(tp),
			Search:    eql_search.NewSearchFunc(tp),
		},

		// Esql
		Esql: Esql{
			Query: esql_query.NewQueryFunc(tp),
		},

		// Graph
		Graph: Graph{
			Explore: graph_explore.NewExploreFunc(tp),
		},

		// Indices
		Indices: Indices{
			AddBlock:              indices_add_block.NewAddBlockFunc(tp),
			Analyze:               indices_analyze.NewAnalyzeFunc(tp),
			Create:                indices_create.NewCreateFunc(tp),
			CreateDataStream:      indices_create_data_stream.NewCreateDataStreamFunc(tp),
			DataStreamsStats:      indices_data_streams_stats.NewDataStreamsStatsFunc(tp),
			Delete:                indices_delete.NewDeleteFunc(tp),
			DeleteAlias:           indices_delete_alias.NewDeleteAliasFunc(tp),
			DeleteDataLifecycle:   indices_delete_data_lifecycle.NewDeleteDataLifecycleFunc(tp),
			DeleteDataStream:      indices_delete_data_stream.NewDeleteDataStreamFunc(tp),
			DeleteIndexTemplate:   indices_delete_index_template.NewDeleteIndexTemplateFunc(tp),
			Exists:                indices_exists.NewExistsFunc(tp),
			ExistsAlias:           indices_exists_alias.NewExistsAliasFunc(tp),
			ExistsIndexTemplate:   indices_exists_index_template.NewExistsIndexTemplateFunc(tp),
			ExplainDataLifecycle:  indices_explain_data_lifecycle.NewExplainDataLifecycleFunc(tp),
			Get:                   indices_get.NewGetFunc(tp),
			GetAlias:              indices_get_alias.NewGetAliasFunc(tp),
			GetDataLifecycle:      indices_get_data_lifecycle.NewGetDataLifecycleFunc(tp),
			GetDataStream:         indices_get_data_stream.NewGetDataStreamFunc(tp),
			GetIndexTemplate:      indices_get_index_template.NewGetIndexTemplateFunc(tp),
			GetMapping:            indices_get_mapping.NewGetMappingFunc(tp),
			GetSettings:           indices_get_settings.NewGetSettingsFunc(tp),
			MigrateToDataStream:   indices_migrate_to_data_stream.NewMigrateToDataStreamFunc(tp),
			ModifyDataStream:      indices_modify_data_stream.NewModifyDataStreamFunc(tp),
			PutAlias:              indices_put_alias.NewPutAliasFunc(tp),
			PutDataLifecycle:      indices_put_data_lifecycle.NewPutDataLifecycleFunc(tp),
			PutIndexTemplate:      indices_put_index_template.NewPutIndexTemplateFunc(tp),
			PutMapping:            indices_put_mapping.NewPutMappingFunc(tp),
			PutSettings:           indices_put_settings.NewPutSettingsFunc(tp),
			PutTemplate:           indices_put_template.NewPutTemplateFunc(tp),
			Refresh:               indices_refresh.NewRefreshFunc(tp),
			ResolveIndex:          indices_resolve_index.NewResolveIndexFunc(tp),
			Rollover:              indices_rollover.NewRolloverFunc(tp),
			SimulateIndexTemplate: indices_simulate_index_template.NewSimulateIndexTemplateFunc(tp),
			SimulateTemplate:      indices_simulate_template.NewSimulateTemplateFunc(tp),
			UpdateAliases:         indices_update_aliases.NewUpdateAliasesFunc(tp),
			ValidateQuery:         indices_validate_query.NewValidateQueryFunc(tp),
		},

		// Inference
		Inference: Inference{
			Delete:    inference_delete.NewDeleteFunc(tp),
			Get:       inference_get.NewGetFunc(tp),
			Inference: inference_inference.NewInferenceFunc(tp),
			Put:       inference_put.NewPutFunc(tp),
		},

		// Ingest
		Ingest: Ingest{
			DeletePipeline: ingest_delete_pipeline.NewDeletePipelineFunc(tp),
			GetPipeline:    ingest_get_pipeline.NewGetPipelineFunc(tp),
			ProcessorGrok:  ingest_processor_grok.NewProcessorGrokFunc(tp),
			PutPipeline:    ingest_put_pipeline.NewPutPipelineFunc(tp),
			Simulate:       ingest_simulate.NewSimulateFunc(tp),
		},

		// License
		License: License{
			Get: license_get.NewGetFunc(tp),
		},

		// Logstash
		Logstash: Logstash{
			DeletePipeline: logstash_delete_pipeline.NewDeletePipelineFunc(tp),
			GetPipeline:    logstash_get_pipeline.NewGetPipelineFunc(tp),
			PutPipeline:    logstash_put_pipeline.NewPutPipelineFunc(tp),
		},

		// Ml
		Ml: Ml{
			CloseJob:                      ml_close_job.NewCloseJobFunc(tp),
			DeleteCalendar:                ml_delete_calendar.NewDeleteCalendarFunc(tp),
			DeleteCalendarEvent:           ml_delete_calendar_event.NewDeleteCalendarEventFunc(tp),
			DeleteCalendarJob:             ml_delete_calendar_job.NewDeleteCalendarJobFunc(tp),
			DeleteDataFrameAnalytics:      ml_delete_data_frame_analytics.NewDeleteDataFrameAnalyticsFunc(tp),
			DeleteDatafeed:                ml_delete_datafeed.NewDeleteDatafeedFunc(tp),
			DeleteFilter:                  ml_delete_filter.NewDeleteFilterFunc(tp),
			DeleteJob:                     ml_delete_job.NewDeleteJobFunc(tp),
			DeleteTrainedModel:            ml_delete_trained_model.NewDeleteTrainedModelFunc(tp),
			DeleteTrainedModelAlias:       ml_delete_trained_model_alias.NewDeleteTrainedModelAliasFunc(tp),
			EstimateModelMemory:           ml_estimate_model_memory.NewEstimateModelMemoryFunc(tp),
			EvaluateDataFrame:             ml_evaluate_data_frame.NewEvaluateDataFrameFunc(tp),
			FlushJob:                      ml_flush_job.NewFlushJobFunc(tp),
			GetCalendarEvents:             ml_get_calendar_events.NewGetCalendarEventsFunc(tp),
			GetCalendars:                  ml_get_calendars.NewGetCalendarsFunc(tp),
			GetDataFrameAnalytics:         ml_get_data_frame_analytics.NewGetDataFrameAnalyticsFunc(tp),
			GetDataFrameAnalyticsStats:    ml_get_data_frame_analytics_stats.NewGetDataFrameAnalyticsStatsFunc(tp),
			GetDatafeedStats:              ml_get_datafeed_stats.NewGetDatafeedStatsFunc(tp),
			GetDatafeeds:                  ml_get_datafeeds.NewGetDatafeedsFunc(tp),
			GetFilters:                    ml_get_filters.NewGetFiltersFunc(tp),
			GetJobStats:                   ml_get_job_stats.NewGetJobStatsFunc(tp),
			GetJobs:                       ml_get_jobs.NewGetJobsFunc(tp),
			GetOverallBuckets:             ml_get_overall_buckets.NewGetOverallBucketsFunc(tp),
			GetTrainedModels:              ml_get_trained_models.NewGetTrainedModelsFunc(tp),
			GetTrainedModelsStats:         ml_get_trained_models_stats.NewGetTrainedModelsStatsFunc(tp),
			InferTrainedModel:             ml_infer_trained_model.NewInferTrainedModelFunc(tp),
			OpenJob:                       ml_open_job.NewOpenJobFunc(tp),
			PostCalendarEvents:            ml_post_calendar_events.NewPostCalendarEventsFunc(tp),
			PreviewDataFrameAnalytics:     ml_preview_data_frame_analytics.NewPreviewDataFrameAnalyticsFunc(tp),
			PreviewDatafeed:               ml_preview_datafeed.NewPreviewDatafeedFunc(tp),
			PutCalendar:                   ml_put_calendar.NewPutCalendarFunc(tp),
			PutCalendarJob:                ml_put_calendar_job.NewPutCalendarJobFunc(tp),
			PutDataFrameAnalytics:         ml_put_data_frame_analytics.NewPutDataFrameAnalyticsFunc(tp),
			PutDatafeed:                   ml_put_datafeed.NewPutDatafeedFunc(tp),
			PutFilter:                     ml_put_filter.NewPutFilterFunc(tp),
			PutJob:                        ml_put_job.NewPutJobFunc(tp),
			PutTrainedModel:               ml_put_trained_model.NewPutTrainedModelFunc(tp),
			PutTrainedModelAlias:          ml_put_trained_model_alias.NewPutTrainedModelAliasFunc(tp),
			PutTrainedModelDefinitionPart: ml_put_trained_model_definition_part.NewPutTrainedModelDefinitionPartFunc(tp),
			PutTrainedModelVocabulary:     ml_put_trained_model_vocabulary.NewPutTrainedModelVocabularyFunc(tp),
			ResetJob:                      ml_reset_job.NewResetJobFunc(tp),
			StartDataFrameAnalytics:       ml_start_data_frame_analytics.NewStartDataFrameAnalyticsFunc(tp),
			StartDatafeed:                 ml_start_datafeed.NewStartDatafeedFunc(tp),
			StartTrainedModelDeployment:   ml_start_trained_model_deployment.NewStartTrainedModelDeploymentFunc(tp),
			StopDataFrameAnalytics:        ml_stop_data_frame_analytics.NewStopDataFrameAnalyticsFunc(tp),
			StopDatafeed:                  ml_stop_datafeed.NewStopDatafeedFunc(tp),
			StopTrainedModelDeployment:    ml_stop_trained_model_deployment.NewStopTrainedModelDeploymentFunc(tp),
			UpdateDataFrameAnalytics:      ml_update_data_frame_analytics.NewUpdateDataFrameAnalyticsFunc(tp),
			UpdateDatafeed:                ml_update_datafeed.NewUpdateDatafeedFunc(tp),
			UpdateFilter:                  ml_update_filter.NewUpdateFilterFunc(tp),
			UpdateJob:                     ml_update_job.NewUpdateJobFunc(tp),
			UpdateTrainedModelDeployment:  ml_update_trained_model_deployment.NewUpdateTrainedModelDeploymentFunc(tp),
		},

		// QueryRuleset
		QueryRuleset: QueryRuleset{
			Delete: query_ruleset_delete.NewDeleteFunc(tp),
			Get:    query_ruleset_get.NewGetFunc(tp),
			List:   query_ruleset_list.NewListFunc(tp),
			Put:    query_ruleset_put.NewPutFunc(tp),
		},

		// SearchApplication
		SearchApplication: SearchApplication{
			Delete:                    search_application_delete.NewDeleteFunc(tp),
			DeleteBehavioralAnalytics: search_application_delete_behavioral_analytics.NewDeleteBehavioralAnalyticsFunc(tp),
			Get:                       search_application_get.NewGetFunc(tp),
			GetBehavioralAnalytics:    search_application_get_behavioral_analytics.NewGetBehavioralAnalyticsFunc(tp),
			List:                      search_application_list.NewListFunc(tp),
			Put:                       search_application_put.NewPutFunc(tp),
			PutBehavioralAnalytics:    search_application_put_behavioral_analytics.NewPutBehavioralAnalyticsFunc(tp),
			Search:                    search_application_search.NewSearchFunc(tp),
		},

		// Security
		Security: Security{
			Authenticate:     security_authenticate.NewAuthenticateFunc(tp),
			CreateApiKey:     security_create_api_key.NewCreateApiKeyFunc(tp),
			GetApiKey:        security_get_api_key.NewGetApiKeyFunc(tp),
			HasPrivileges:    security_has_privileges.NewHasPrivilegesFunc(tp),
			InvalidateApiKey: security_invalidate_api_key.NewInvalidateApiKeyFunc(tp),
			QueryApiKeys:     security_query_api_keys.NewQueryApiKeysFunc(tp),
			UpdateApiKey:     security_update_api_key.NewUpdateApiKeyFunc(tp),
		},

		// Sql
		Sql: Sql{
			ClearCursor:    sql_clear_cursor.NewClearCursorFunc(tp),
			DeleteAsync:    sql_delete_async.NewDeleteAsyncFunc(tp),
			GetAsync:       sql_get_async.NewGetAsyncFunc(tp),
			GetAsyncStatus: sql_get_async_status.NewGetAsyncStatusFunc(tp),
			Query:          sql_query.NewQueryFunc(tp),
			Translate:      sql_translate.NewTranslateFunc(tp),
		},

		// Synonyms
		Synonyms: Synonyms{
			DeleteSynonym:     synonyms_delete_synonym.NewDeleteSynonymFunc(tp),
			DeleteSynonymRule: synonyms_delete_synonym_rule.NewDeleteSynonymRuleFunc(tp),
			GetSynonym:        synonyms_get_synonym.NewGetSynonymFunc(tp),
			GetSynonymRule:    synonyms_get_synonym_rule.NewGetSynonymRuleFunc(tp),
			GetSynonymsSets:   synonyms_get_synonyms_sets.NewGetSynonymsSetsFunc(tp),
			PutSynonym:        synonyms_put_synonym.NewPutSynonymFunc(tp),
			PutSynonymRule:    synonyms_put_synonym_rule.NewPutSynonymRuleFunc(tp),
		},

		// Tasks
		Tasks: Tasks{
			Get: tasks_get.NewGetFunc(tp),
		},

		// Transform
		Transform: Transform{
			DeleteTransform:      transform_delete_transform.NewDeleteTransformFunc(tp),
			GetTransform:         transform_get_transform.NewGetTransformFunc(tp),
			GetTransformStats:    transform_get_transform_stats.NewGetTransformStatsFunc(tp),
			PreviewTransform:     transform_preview_transform.NewPreviewTransformFunc(tp),
			PutTransform:         transform_put_transform.NewPutTransformFunc(tp),
			ResetTransform:       transform_reset_transform.NewResetTransformFunc(tp),
			ScheduleNowTransform: transform_schedule_now_transform.NewScheduleNowTransformFunc(tp),
			StartTransform:       transform_start_transform.NewStartTransformFunc(tp),
			StopTransform:        transform_stop_transform.NewStopTransformFunc(tp),
			UpdateTransform:      transform_update_transform.NewUpdateTransformFunc(tp),
		},

		Bulk:                   core_bulk.NewBulkFunc(tp),
		ClearScroll:            core_clear_scroll.NewClearScrollFunc(tp),
		ClosePointInTime:       core_close_point_in_time.NewClosePointInTimeFunc(tp),
		Count:                  core_count.NewCountFunc(tp),
		Create:                 core_create.NewCreateFunc(tp),
		Delete:                 core_delete.NewDeleteFunc(tp),
		DeleteByQuery:          core_delete_by_query.NewDeleteByQueryFunc(tp),
		DeleteScript:           core_delete_script.NewDeleteScriptFunc(tp),
		Exists:                 core_exists.NewExistsFunc(tp),
		ExistsSource:           core_exists_source.NewExistsSourceFunc(tp),
		Explain:                core_explain.NewExplainFunc(tp),
		FieldCaps:              core_field_caps.NewFieldCapsFunc(tp),
		Get:                    core_get.NewGetFunc(tp),
		GetScript:              core_get_script.NewGetScriptFunc(tp),
		GetSource:              core_get_source.NewGetSourceFunc(tp),
		Index:                  core_index.NewIndexFunc(tp),
		Info:                   core_info.NewInfoFunc(tp),
		Mget:                   core_mget.NewMgetFunc(tp),
		Msearch:                core_msearch.NewMsearchFunc(tp),
		MsearchTemplate:        core_msearch_template.NewMsearchTemplateFunc(tp),
		Mtermvectors:           core_mtermvectors.NewMtermvectorsFunc(tp),
		OpenPointInTime:        core_open_point_in_time.NewOpenPointInTimeFunc(tp),
		Ping:                   core_ping.NewPingFunc(tp),
		PutScript:              core_put_script.NewPutScriptFunc(tp),
		RankEval:               core_rank_eval.NewRankEvalFunc(tp),
		Reindex:                core_reindex.NewReindexFunc(tp),
		RenderSearchTemplate:   core_render_search_template.NewRenderSearchTemplateFunc(tp),
		ScriptsPainlessExecute: core_scripts_painless_execute.NewScriptsPainlessExecuteFunc(tp),
		Scroll:                 core_scroll.NewScrollFunc(tp),
		Search:                 core_search.NewSearchFunc(tp),
		SearchMvt:              core_search_mvt.NewSearchMvtFunc(tp),
		SearchTemplate:         core_search_template.NewSearchTemplateFunc(tp),
		TermsEnum:              core_terms_enum.NewTermsEnumFunc(tp),
		Termvectors:            core_termvectors.NewTermvectorsFunc(tp),
		Update:                 core_update.NewUpdateFunc(tp),
		UpdateByQuery:          core_update_by_query.NewUpdateByQueryFunc(tp),
	}
}
