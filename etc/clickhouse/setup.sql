DROP DATABASE IF EXISTS iris;

CREATE DATABASE IF NOT EXISTS iris;

USE iris;

CREATE TABLE IF NOT EXISTS iris.source
(
    account String,
    timestamp DateTime,
    event_type String,
    visitor_id String,
    session_id String,
    event_data Nullable(String),
    document_location String,
    referrer_location Nullable(String),
    document_encoding Nullable(String),
    screen_resolution Nullable(String),
    view_port Nullable(String),
    color_depth Nullable(UInt16),
    document_title Nullable(String),
    browser_name Nullable(String),
    is_mobile_device Nullable(UInt8),
    user_agent Nullable(String),
    timezone_offset Nullable(Int16),
    utm Nullable(String),
    ip_address String,
    utm_source String MATERIALIZED if(visitParamHas(utm, 'utm_source'), visitParamExtractString(utm, 'utm_source'), '_'),
    utm_medium String MATERIALIZED if(visitParamHas(utm, 'utm_medium'), visitParamExtractString(utm, 'utm_medium'), '_'),
    utm_term String MATERIALIZED if(visitParamHas(utm, 'utm_term'), visitParamExtractString(utm, 'utm_term'), '_'),
    utm_content String MATERIALIZED if(visitParamHas(utm, 'utm_content'), visitParamExtractString(utm, 'utm_content'), '_'),
    utm_campaign String MATERIALIZED if(visitParamHas(utm, 'utm_campaign'), visitParamExtractString(utm, 'utm_campaign'), '_')
) ENGINE = ReplicatedMergeTree('/clickhouse/tables/iris/source/{shard}/', '{replica}')
PARTITION BY (account, toYYYYMM(timestamp))
ORDER BY (account, timestamp, event_type, session_id)
;

CREATE TABLE iris.source_buffer AS iris.source ENGINE = Buffer(iris, source, 16, 10, 20, 1000, 100000, 1000000, 10000000);
