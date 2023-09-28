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
	// Deletes an async search by ID. If the search is still running, the search
	// request will be cancelled. Otherwise, the saved search results are deleted.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/async-search.html
	Delete async_search_delete.NewDelete
	// Retrieves the results of a previously submitted async search request given
	// its ID.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/async-search.html
	Get async_search_get.NewGet
	// Retrieves the status of a previously submitted async search request given its
	// ID.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/async-search.html
	Status async_search_status.NewStatus
	// Executes a search request asynchronously.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/async-search.html
	Submit async_search_submit.NewSubmit
}

type Cat struct {
	// Shows information about currently configured aliases to indices including
	// filter and routing infos.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-alias.html
	Aliases cat_aliases.NewAliases
	// Returns information about existing component_templates templates.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-component-templates.html
	ComponentTemplates cat_component_templates.NewComponentTemplates
	// Provides quick access to the document count of the entire cluster, or
	// individual indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-count.html
	Count cat_count.NewCount
	// Returns help for the Cat APIs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat.html
	Help cat_help.NewHelp
	// Returns information about indices: number of primaries and replicas, document
	// counts, disk size, ...
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-indices.html
	Indices cat_indices.NewIndices
	// Gets configuration and usage information about data frame analytics jobs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-dfanalytics.html
	MlDataFrameAnalytics cat_ml_data_frame_analytics.NewMlDataFrameAnalytics
	// Gets configuration and usage information about datafeeds.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-datafeeds.html
	MlDatafeeds cat_ml_datafeeds.NewMlDatafeeds
	// Gets configuration and usage information about anomaly detection jobs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-anomaly-detectors.html
	MlJobs cat_ml_jobs.NewMlJobs
	// Gets configuration and usage information about inference trained models.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-trained-model.html
	MlTrainedModels cat_ml_trained_models.NewMlTrainedModels
	// Gets configuration and usage information about transforms.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-transforms.html
	Transforms cat_transforms.NewTransforms
}

type Cluster struct {
	// Deletes a component template
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html
	DeleteComponentTemplate cluster_delete_component_template.NewDeleteComponentTemplate
	// Returns information about whether a particular component template exist
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html
	ExistsComponentTemplate cluster_exists_component_template.NewExistsComponentTemplate
	// Returns one or more component templates
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html
	GetComponentTemplate cluster_get_component_template.NewGetComponentTemplate
	// Returns different information about the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-info.html
	Info cluster_info.NewInfo
	// Creates or updates a component template
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html
	PutComponentTemplate cluster_put_component_template.NewPutComponentTemplate
}

type Core struct {
	// Allows to perform multiple index/update/delete operations in a single
	// request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-bulk.html
	Bulk core_bulk.NewBulk
	// Explicitly clears the search context for a scroll.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clear-scroll-api.html
	ClearScroll core_clear_scroll.NewClearScroll
	// Close a point in time
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
	ClosePointInTime core_close_point_in_time.NewClosePointInTime
	// Returns number of documents matching a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-count.html
	Count core_count.NewCount
	// Creates a new document in the index.
	//
	// Returns a 409 response when a document with a same ID already exists in the
	// index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
	Create core_create.NewCreate
	// Removes a document from the index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete.html
	Delete core_delete.NewDelete
	// Deletes documents matching the provided query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete-by-query.html
	DeleteByQuery core_delete_by_query.NewDeleteByQuery
	// Deletes a script.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	DeleteScript core_delete_script.NewDeleteScript
	// Returns information about whether a document exists in an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	Exists core_exists.NewExists
	// Returns information about whether a document source exists in an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	ExistsSource core_exists_source.NewExistsSource
	// Returns information about why a specific matches (or doesn't match) a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-explain.html
	Explain core_explain.NewExplain
	// Returns the information about the capabilities of fields among multiple
	// indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-field-caps.html
	FieldCaps core_field_caps.NewFieldCaps
	// Returns a document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	Get core_get.NewGet
	// Returns a script.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	GetScript core_get_script.NewGetScript
	// Returns the source of a document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	GetSource core_get_source.NewGetSource
	// Creates or updates a document in an index.
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
	// Allows to execute several search template operations in one request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-multi-search.html
	MsearchTemplate core_msearch_template.NewMsearchTemplate
	// Returns multiple termvectors in one request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-multi-termvectors.html
	Mtermvectors core_mtermvectors.NewMtermvectors
	// Open a point in time that can be used in subsequent searches
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
	OpenPointInTime core_open_point_in_time.NewOpenPointInTime
	// Returns whether the cluster is running.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html
	Ping core_ping.NewPing
	// Creates or updates a script.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	PutScript core_put_script.NewPutScript
	// Allows to evaluate the quality of ranked search results over a set of typical
	// search queries
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-rank-eval.html
	RankEval core_rank_eval.NewRankEval
	// Allows to copy documents from one index to another, optionally filtering the
	// source
	// documents by a query, changing the destination index settings, or fetching
	// the
	// documents from a remote cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-reindex.html
	Reindex core_reindex.NewReindex
	// Allows to use the Mustache language to pre-render a search definition.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/render-search-template-api.html
	RenderSearchTemplate core_render_search_template.NewRenderSearchTemplate
	// Allows an arbitrary script to be executed and a result to be returned
	// https://www.elastic.co/guide/en/elasticsearch/painless/current/painless-execute-api.html
	ScriptsPainlessExecute core_scripts_painless_execute.NewScriptsPainlessExecute
	// Allows to retrieve a large numbers of results from a single search request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-body.html#request-body-search-scroll
	Scroll core_scroll.NewScroll
	// Returns results matching a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-search.html
	Search core_search.NewSearch
	// Searches a vector tile for geospatial values. Returns results as a binary
	// Mapbox vector tile.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-vector-tile-api.html
	SearchMvt core_search_mvt.NewSearchMvt
	// Allows to use the Mustache language to pre-render a search definition.
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
	// Performs an update on every document in the index without changing the
	// source,
	// for example to pick up a mapping change.
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
	// Gets information about an enrich policy.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-enrich-policy-api.html
	GetPolicy enrich_get_policy.NewGetPolicy
	// Creates a new enrich policy.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-enrich-policy-api.html
	PutPolicy enrich_put_policy.NewPutPolicy
	// Gets enrich coordinator statistics and information about enrich policies that
	// are currently executing.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/enrich-stats-api.html
	Stats enrich_stats.NewStats
}

type Eql struct {
	// Deletes an async EQL search by ID. If the search is still running, the search
	// request will be cancelled. Otherwise, the saved search results are deleted.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/eql-search-api.html
	Delete eql_delete.NewDelete
	// Returns async results from previously executed Event Query Language (EQL)
	// search
	//  https://www.elastic.co/guide/en/elasticsearch/reference/current/get-async-eql-search-api.html
	Get eql_get.NewGet
	// Returns the status of a previously submitted async or stored Event Query
	// Language (EQL) search
	//  https://www.elastic.co/guide/en/elasticsearch/reference/current/get-async-eql-status-api.html
	GetStatus eql_get_status.NewGetStatus
	// Returns results matching a query expressed in Event Query Language (EQL)
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/eql-search-api.html
	Search eql_search.NewSearch
}

type Graph struct {
	// Explore extracted and summarized information about the documents and terms in
	// an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/graph-explore-api.html
	Explore graph_explore.NewExplore
}

type Indices struct {
	// Adds a block to an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/index-modules-blocks.html
	AddBlock indices_add_block.NewAddBlock
	// Performs the analysis process on a text and return the tokens breakdown of
	// the text.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-analyze.html
	Analyze indices_analyze.NewAnalyze
	// Creates an index with optional settings and mappings.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html
	Create indices_create.NewCreate
	// Creates a data stream
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	CreateDataStream indices_create_data_stream.NewCreateDataStream
	// Provides statistics on operations happening in a data stream.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	DataStreamsStats indices_data_streams_stats.NewDataStreamsStats
	// Deletes an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-delete-index.html
	Delete indices_delete.NewDelete
	// Deletes an alias.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-aliases.html
	DeleteAlias indices_delete_alias.NewDeleteAlias
	// Deletes the data stream lifecycle of the selected data streams.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams-delete-lifecycle.html
	DeleteDataLifecycle indices_delete_data_lifecycle.NewDeleteDataLifecycle
	// Deletes a data stream.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	DeleteDataStream indices_delete_data_stream.NewDeleteDataStream
	// Deletes an index template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates.html
	DeleteIndexTemplate indices_delete_index_template.NewDeleteIndexTemplate
	// Returns information about whether a particular index exists.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-exists.html
	Exists indices_exists.NewExists
	// Returns information about whether a particular alias exists.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-aliases.html
	ExistsAlias indices_exists_alias.NewExistsAlias
	// Returns information about whether a particular index template exists.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates.html
	ExistsIndexTemplate indices_exists_index_template.NewExistsIndexTemplate
	// Retrieves information about the index's current data stream lifecycle, such
	// as any potential encountered error, time since creation etc.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams-explain-lifecycle.html
	ExplainDataLifecycle indices_explain_data_lifecycle.NewExplainDataLifecycle
	// Returns information about one or more indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-index.html
	Get indices_get.NewGet
	// Returns an alias.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-aliases.html
	GetAlias indices_get_alias.NewGetAlias
	// Returns the data stream lifecycle of the selected data streams.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams-get-lifecycle.html
	GetDataLifecycle indices_get_data_lifecycle.NewGetDataLifecycle
	// Returns data streams.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	GetDataStream indices_get_data_stream.NewGetDataStream
	// Returns an index template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates.html
	GetIndexTemplate indices_get_index_template.NewGetIndexTemplate
	// Returns mappings for one or more indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-mapping.html
	GetMapping indices_get_mapping.NewGetMapping
	// Returns settings for one or more indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-settings.html
	GetSettings indices_get_settings.NewGetSettings
	// Migrates an alias to a data stream
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	MigrateToDataStream indices_migrate_to_data_stream.NewMigrateToDataStream
	// Modifies a data stream
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams.html
	ModifyDataStream indices_modify_data_stream.NewModifyDataStream
	// Creates or updates an alias.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-aliases.html
	PutAlias indices_put_alias.NewPutAlias
	// Updates the data stream lifecycle of the selected data streams.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams-put-lifecycle.html
	PutDataLifecycle indices_put_data_lifecycle.NewPutDataLifecycle
	// Creates or updates an index template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates.html
	PutIndexTemplate indices_put_index_template.NewPutIndexTemplate
	// Updates the index mappings.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-put-mapping.html
	PutMapping indices_put_mapping.NewPutMapping
	// Updates the index settings.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-update-settings.html
	PutSettings indices_put_settings.NewPutSettings
	// Creates or updates an index template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates.html
	PutTemplate indices_put_template.NewPutTemplate
	// Performs the refresh operation in one or more indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-refresh.html
	Refresh indices_refresh.NewRefresh
	// Returns information about any matching indices, aliases, and data streams
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-resolve-index-api.html
	ResolveIndex indices_resolve_index.NewResolveIndex
	// Updates an alias to point to a new index when the existing index
	// is considered to be too large or too old.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-rollover-index.html
	Rollover indices_rollover.NewRollover
	// Simulate matching the given index name against the index templates in the
	// system
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates.html
	SimulateIndexTemplate indices_simulate_index_template.NewSimulateIndexTemplate
	// Simulate resolving the given template name or body
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates.html
	SimulateTemplate indices_simulate_template.NewSimulateTemplate
	// Updates index aliases.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-aliases.html
	UpdateAliases indices_update_aliases.NewUpdateAliases
	// Allows a user to validate a potentially expensive query without executing it.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-validate.html
	ValidateQuery indices_validate_query.NewValidateQuery
}

type Ingest struct {
	// Deletes a pipeline.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-pipeline-api.html
	DeletePipeline ingest_delete_pipeline.NewDeletePipeline
	// Returns a pipeline.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-pipeline-api.html
	GetPipeline ingest_get_pipeline.NewGetPipeline
	// Returns a list of the built-in patterns.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/grok-processor.html
	ProcessorGrok ingest_processor_grok.NewProcessorGrok
	// Creates or updates a pipeline.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest.html
	PutPipeline ingest_put_pipeline.NewPutPipeline
	// Allows to simulate a pipeline with example documents.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/simulate-pipeline-api.html
	Simulate ingest_simulate.NewSimulate
}

type License struct {
	// Retrieves licensing information for the cluster
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-license.html
	Get license_get.NewGet
}

type Logstash struct {
	// Deletes Logstash Pipelines used by Central Management
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/logstash-api-delete-pipeline.html
	DeletePipeline logstash_delete_pipeline.NewDeletePipeline
	// Retrieves Logstash Pipelines used by Central Management
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/logstash-api-get-pipeline.html
	GetPipeline logstash_get_pipeline.NewGetPipeline
	// Adds and updates Logstash Pipelines used for Central Management
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/logstash-api-put-pipeline.html
	PutPipeline logstash_put_pipeline.NewPutPipeline
}

type Ml struct {
	// Closes one or more anomaly detection jobs. A job can be opened and closed
	// multiple times throughout its lifecycle.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-close-job.html
	CloseJob ml_close_job.NewCloseJob
	// Deletes a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-calendar.html
	DeleteCalendar ml_delete_calendar.NewDeleteCalendar
	// Deletes scheduled events from a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-calendar-event.html
	DeleteCalendarEvent ml_delete_calendar_event.NewDeleteCalendarEvent
	// Deletes anomaly detection jobs from a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-calendar-job.html
	DeleteCalendarJob ml_delete_calendar_job.NewDeleteCalendarJob
	// Deletes an existing data frame analytics job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-dfanalytics.html
	DeleteDataFrameAnalytics ml_delete_data_frame_analytics.NewDeleteDataFrameAnalytics
	// Deletes an existing datafeed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-datafeed.html
	DeleteDatafeed ml_delete_datafeed.NewDeleteDatafeed
	// Deletes a filter.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-filter.html
	DeleteFilter ml_delete_filter.NewDeleteFilter
	// Deletes an existing anomaly detection job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-job.html
	DeleteJob ml_delete_job.NewDeleteJob
	// Deletes an existing trained inference model that is currently not referenced
	// by an ingest pipeline.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-trained-models.html
	DeleteTrainedModel ml_delete_trained_model.NewDeleteTrainedModel
	// Deletes a model alias that refers to the trained model
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-trained-models-aliases.html
	DeleteTrainedModelAlias ml_delete_trained_model_alias.NewDeleteTrainedModelAlias
	// Estimates the model memory
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-apis.html
	EstimateModelMemory ml_estimate_model_memory.NewEstimateModelMemory
	// Evaluates the data frame analytics for an annotated index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/evaluate-dfanalytics.html
	EvaluateDataFrame ml_evaluate_data_frame.NewEvaluateDataFrame
	// Forces any buffered data to be processed by the job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-flush-job.html
	FlushJob ml_flush_job.NewFlushJob
	// Retrieves information about the scheduled events in calendars.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-calendar-event.html
	GetCalendarEvents ml_get_calendar_events.NewGetCalendarEvents
	// Retrieves configuration information for calendars.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-calendar.html
	GetCalendars ml_get_calendars.NewGetCalendars
	// Retrieves configuration information for data frame analytics jobs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-dfanalytics.html
	GetDataFrameAnalytics ml_get_data_frame_analytics.NewGetDataFrameAnalytics
	// Retrieves usage information for data frame analytics jobs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-dfanalytics-stats.html
	GetDataFrameAnalyticsStats ml_get_data_frame_analytics_stats.NewGetDataFrameAnalyticsStats
	// Retrieves usage information for datafeeds.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-datafeed-stats.html
	GetDatafeedStats ml_get_datafeed_stats.NewGetDatafeedStats
	// Retrieves configuration information for datafeeds.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-datafeed.html
	GetDatafeeds ml_get_datafeeds.NewGetDatafeeds
	// Retrieves filters.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-filter.html
	GetFilters ml_get_filters.NewGetFilters
	// Retrieves usage information for anomaly detection jobs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-job-stats.html
	GetJobStats ml_get_job_stats.NewGetJobStats
	// Retrieves configuration information for anomaly detection jobs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-job.html
	GetJobs ml_get_jobs.NewGetJobs
	// Retrieves overall bucket results that summarize the bucket results of
	// multiple anomaly detection jobs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-overall-buckets.html
	GetOverallBuckets ml_get_overall_buckets.NewGetOverallBuckets
	// Retrieves configuration information for a trained inference model.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-trained-models.html
	GetTrainedModels ml_get_trained_models.NewGetTrainedModels
	// Retrieves usage information for trained inference models.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-trained-models-stats.html
	GetTrainedModelsStats ml_get_trained_models_stats.NewGetTrainedModelsStats
	// Evaluate a trained model.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-trained-model.html
	InferTrainedModel ml_infer_trained_model.NewInferTrainedModel
	// Opens one or more anomaly detection jobs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-open-job.html
	OpenJob ml_open_job.NewOpenJob
	// Posts scheduled events in a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-post-calendar-event.html
	PostCalendarEvents ml_post_calendar_events.NewPostCalendarEvents
	// Previews that will be analyzed given a data frame analytics config.
	// http://www.elastic.co/guide/en/elasticsearch/reference/current/preview-dfanalytics.html
	PreviewDataFrameAnalytics ml_preview_data_frame_analytics.NewPreviewDataFrameAnalytics
	// Previews a datafeed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-preview-datafeed.html
	PreviewDatafeed ml_preview_datafeed.NewPreviewDatafeed
	// Instantiates a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-calendar.html
	PutCalendar ml_put_calendar.NewPutCalendar
	// Adds an anomaly detection job to a calendar.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-calendar-job.html
	PutCalendarJob ml_put_calendar_job.NewPutCalendarJob
	// Instantiates a data frame analytics job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-dfanalytics.html
	PutDataFrameAnalytics ml_put_data_frame_analytics.NewPutDataFrameAnalytics
	// Instantiates a datafeed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-datafeed.html
	PutDatafeed ml_put_datafeed.NewPutDatafeed
	// Instantiates a filter.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-filter.html
	PutFilter ml_put_filter.NewPutFilter
	// Instantiates an anomaly detection job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-put-job.html
	PutJob ml_put_job.NewPutJob
	// Creates an inference trained model.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-models.html
	PutTrainedModel ml_put_trained_model.NewPutTrainedModel
	// Creates a new model alias (or reassigns an existing one) to refer to the
	// trained model
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-models-aliases.html
	PutTrainedModelAlias ml_put_trained_model_alias.NewPutTrainedModelAlias
	// Creates part of a trained model definition
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-model-definition-part.html
	PutTrainedModelDefinitionPart ml_put_trained_model_definition_part.NewPutTrainedModelDefinitionPart
	// Creates a trained model vocabulary
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-trained-model-vocabulary.html
	PutTrainedModelVocabulary ml_put_trained_model_vocabulary.NewPutTrainedModelVocabulary
	// Resets an existing anomaly detection job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-reset-job.html
	ResetJob ml_reset_job.NewResetJob
	// Starts a data frame analytics job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/start-dfanalytics.html
	StartDataFrameAnalytics ml_start_data_frame_analytics.NewStartDataFrameAnalytics
	// Starts one or more datafeeds.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-start-datafeed.html
	StartDatafeed ml_start_datafeed.NewStartDatafeed
	// Start a trained model deployment.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/start-trained-model-deployment.html
	StartTrainedModelDeployment ml_start_trained_model_deployment.NewStartTrainedModelDeployment
	// Stops one or more data frame analytics jobs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/stop-dfanalytics.html
	StopDataFrameAnalytics ml_stop_data_frame_analytics.NewStopDataFrameAnalytics
	// Stops one or more datafeeds.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-stop-datafeed.html
	StopDatafeed ml_stop_datafeed.NewStopDatafeed
	// Stop a trained model deployment.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/stop-trained-model-deployment.html
	StopTrainedModelDeployment ml_stop_trained_model_deployment.NewStopTrainedModelDeployment
	// Updates certain properties of a data frame analytics job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-dfanalytics.html
	UpdateDataFrameAnalytics ml_update_data_frame_analytics.NewUpdateDataFrameAnalytics
	// Updates certain properties of a datafeed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-update-datafeed.html
	UpdateDatafeed ml_update_datafeed.NewUpdateDatafeed
	// Updates the description of a filter, adds items, or removes items.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-update-filter.html
	UpdateFilter ml_update_filter.NewUpdateFilter
	// Updates certain properties of an anomaly detection job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-update-job.html
	UpdateJob ml_update_job.NewUpdateJob
}

type QueryRuleset struct {
	// Deletes a query ruleset.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-query-ruleset.html
	Delete query_ruleset_delete.NewDelete
	// Returns the details about a query ruleset.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-query-ruleset.html
	Get query_ruleset_get.NewGet
	// Lists query rulesets.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/list-query-rulesets.html
	List query_ruleset_list.NewList
	// Creates or updates a query ruleset.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-query-ruleset.html
	Put query_ruleset_put.NewPut
}

type SearchApplication struct {
	// Deletes a search application.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-search-application.html
	Delete search_application_delete.NewDelete
	// Delete a behavioral analytics collection.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-analytics-collection.html
	DeleteBehavioralAnalytics search_application_delete_behavioral_analytics.NewDeleteBehavioralAnalytics
	// Returns the details about a search application.
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
	// Perform a search against a search application
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-application-search.html
	Search search_application_search.NewSearch
}

type Security struct {
	// Enables authentication as a user and retrieve information about the
	// authenticated user.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-authenticate.html
	Authenticate security_authenticate.NewAuthenticate
	// Creates an API key for access without requiring basic authentication.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-create-api-key.html
	CreateApiKey security_create_api_key.NewCreateApiKey
	// Retrieves information for one or more API keys.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-api-key.html
	GetApiKey security_get_api_key.NewGetApiKey
	// Determines whether the specified user has a specified list of privileges.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-has-privileges.html
	HasPrivileges security_has_privileges.NewHasPrivileges
	// Invalidates one or more API keys.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-invalidate-api-key.html
	InvalidateApiKey security_invalidate_api_key.NewInvalidateApiKey
	// Retrieves information for API keys using a subset of query DSL
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-query-api-key.html
	QueryApiKeys security_query_api_keys.NewQueryApiKeys
	// Updates attributes of an existing API key.
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
	// Creates or updates a synonyms set
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
	// Deletes an existing transform.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-transform.html
	DeleteTransform transform_delete_transform.NewDeleteTransform
	// Retrieves configuration information for transforms.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-transform.html
	GetTransform transform_get_transform.NewGetTransform
	// Retrieves usage information for transforms.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-transform-stats.html
	GetTransformStats transform_get_transform_stats.NewGetTransformStats
	// Previews a transform.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/preview-transform.html
	PreviewTransform transform_preview_transform.NewPreviewTransform
	// Instantiates a transform.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-transform.html
	PutTransform transform_put_transform.NewPutTransform
	// Resets an existing transform.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/reset-transform.html
	ResetTransform transform_reset_transform.NewResetTransform
	// Schedules now a transform.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/schedule-now-transform.html
	ScheduleNowTransform transform_schedule_now_transform.NewScheduleNowTransform
	// Starts one or more transforms.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/start-transform.html
	StartTransform transform_start_transform.NewStartTransform
	// Stops one or more transforms.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/stop-transform.html
	StopTransform transform_stop_transform.NewStopTransform
	// Updates certain properties of a transform.
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
	Graph             Graph
	Indices           Indices
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

	// Allows to perform multiple index/update/delete operations in a single
	// request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-bulk.html
	Bulk core_bulk.NewBulk
	// Explicitly clears the search context for a scroll.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clear-scroll-api.html
	ClearScroll core_clear_scroll.NewClearScroll
	// Close a point in time
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
	ClosePointInTime core_close_point_in_time.NewClosePointInTime
	// Returns number of documents matching a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-count.html
	Count core_count.NewCount
	// Creates a new document in the index.
	//
	// Returns a 409 response when a document with a same ID already exists in the
	// index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
	Create core_create.NewCreate
	// Removes a document from the index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete.html
	Delete core_delete.NewDelete
	// Deletes documents matching the provided query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete-by-query.html
	DeleteByQuery core_delete_by_query.NewDeleteByQuery
	// Deletes a script.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	DeleteScript core_delete_script.NewDeleteScript
	// Returns information about whether a document exists in an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	Exists core_exists.NewExists
	// Returns information about whether a document source exists in an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	ExistsSource core_exists_source.NewExistsSource
	// Returns information about why a specific matches (or doesn't match) a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-explain.html
	Explain core_explain.NewExplain
	// Returns the information about the capabilities of fields among multiple
	// indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-field-caps.html
	FieldCaps core_field_caps.NewFieldCaps
	// Returns a document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	Get core_get.NewGet
	// Returns a script.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	GetScript core_get_script.NewGetScript
	// Returns the source of a document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	GetSource core_get_source.NewGetSource
	// Creates or updates a document in an index.
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
	// Allows to execute several search template operations in one request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-multi-search.html
	MsearchTemplate core_msearch_template.NewMsearchTemplate
	// Returns multiple termvectors in one request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-multi-termvectors.html
	Mtermvectors core_mtermvectors.NewMtermvectors
	// Open a point in time that can be used in subsequent searches
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
	OpenPointInTime core_open_point_in_time.NewOpenPointInTime
	// Returns whether the cluster is running.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html
	Ping core_ping.NewPing
	// Creates or updates a script.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	PutScript core_put_script.NewPutScript
	// Allows to evaluate the quality of ranked search results over a set of typical
	// search queries
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-rank-eval.html
	RankEval core_rank_eval.NewRankEval
	// Allows to copy documents from one index to another, optionally filtering the
	// source
	// documents by a query, changing the destination index settings, or fetching
	// the
	// documents from a remote cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-reindex.html
	Reindex core_reindex.NewReindex
	// Allows to use the Mustache language to pre-render a search definition.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/render-search-template-api.html
	RenderSearchTemplate core_render_search_template.NewRenderSearchTemplate
	// Allows an arbitrary script to be executed and a result to be returned
	// https://www.elastic.co/guide/en/elasticsearch/painless/current/painless-execute-api.html
	ScriptsPainlessExecute core_scripts_painless_execute.NewScriptsPainlessExecute
	// Allows to retrieve a large numbers of results from a single search request.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-body.html#request-body-search-scroll
	Scroll core_scroll.NewScroll
	// Returns results matching a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-search.html
	Search core_search.NewSearch
	// Searches a vector tile for geospatial values. Returns results as a binary
	// Mapbox vector tile.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-vector-tile-api.html
	SearchMvt core_search_mvt.NewSearchMvt
	// Allows to use the Mustache language to pre-render a search definition.
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
	// Performs an update on every document in the index without changing the
	// source,
	// for example to pick up a mapping change.
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
