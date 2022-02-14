package inquiry

var clickhouseTableDataORM = map[int]string{
	TableTypeTimeString: `create table %s
(
	_timestamp_ DateTime,
	_trace_time_ DateTime64(9, 'Asia/Shanghai'),
	_source_ String,
	_cluster_ String,
	_log_agent_ String,
	_namespace_ String,
	_node_name_ String,
	_node_ip_ String,
	_container_name_ String,
	_pod_name_ String,
	_raw_log_ String
)
engine = MergeTree PARTITION BY toYYYYMMDD(_timestamp_)
ORDER BY _timestamp_
TTL toDateTime(_timestamp_) + INTERVAL %d DAY 
SETTINGS index_granularity = 8192;`,
	TableTypeTimeFloat: `create table %s
(
	_timestamp_ DateTime,
	_trace_time_ DateTime64(9, 'Asia/Shanghai'),
	_source_ String,
	_cluster_ String,
	_log_agent_ String,
	_namespace_ String,
	_node_name_ String,
	_node_ip_ String,
	_container_name_ String,
	_pod_name_ String,
	_raw_log_ String
)
engine = MergeTree PARTITION BY toYYYYMMDD(_timestamp_)
ORDER BY _timestamp_
TTL toDateTime(_timestamp_) + INTERVAL %d DAY
SETTINGS index_granularity = 8192;`,
}

var clickhouseTableStreamORM = map[int]string{
	TableTypeTimeString: `create table %s
(
	_source_ String,
	_pod_name_ String,
	_namespace_ String,
	_node_name_ String,
	_container_name_ String,
	_cluster_ String,
	_log_agent_ String,
	_node_ip_ String,
	_time_ String,
	log String
)
engine = Kafka SETTINGS kafka_broker_list = '%s', kafka_topic_list = '%s', kafka_group_name = '%s', kafka_format = 'JSONEachRow', kafka_num_consumers = 1;`,
	TableTypeTimeFloat: `create table %s
(
	_source_ String,
	_pod_name_ String,
	_namespace_ String,
	_node_name_ String,
	_container_name_ String,
	_cluster_ String,
	_log_agent_ String,
	_node_ip_ String,
	_time_ Float64,
	log String
)
engine = Kafka SETTINGS kafka_broker_list = '%s', kafka_topic_list = '%s', kafka_group_name = '%s', kafka_format = 'JSONEachRow', kafka_num_consumers = 1;`,
}

var clickhouseViewORM = map[int]string{
	TableTypeTimeString: `CREATE MATERIALIZED VIEW %s TO %s AS
SELECT
    %s,
    _source_,
    _cluster_,
    _log_agent_,
    _namespace_,
    _node_name_,
    _node_ip_,
    _container_name_,
    _pod_name_,
	log AS _raw_log_%s
	FROM %s where %s;`,
	TableTypeTimeFloat: `CREATE MATERIALIZED VIEW %s TO %s AS
SELECT
    %s,
	_pod_name_,
	_namespace_,
	_node_name_,
	_container_name_,
	_cluster_,
	_log_agent_,
	_node_ip_,
	_source_,
	log AS _raw_log_%s
	FROM %s where %s;`,
}
