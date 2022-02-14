package view

import (
	"github.com/shimohq/mogo/api/pkg/model/db"
)

type ReqQuery struct {
	Database      string `form:"database"`
	Table         string `form:"table"`
	DatabaseTable string `form:"databaseTable"`
	Field         string `form:"field"`
	Query         string `form:"query"`
	ST            int64  `form:"st"`
	ET            int64  `form:"et"`
	Page          uint32 `form:"page"`
	PageSize      uint32 `form:"pageSize"`
}

type RespQuery struct {
	Limited            uint32                   `json:"limited"`
	Keys               []*db.Index              `json:"keys"`
	ElapsedMillisecond int                      `json:"elapsedMillisecond"`
	Count              uint64                   `json:"count"`
	HasSQL             bool                     `json:"hasSQL"`
	WhereQuery         string                   `json:"whereQuery"`
	ProcessedRows      int                      `json:"processedRows"`
	Terms              [][]string               `json:"terms"`
	Marker             string                   `json:"marker"`
	Progress           string                   `json:"progress"`
	CpuSec             int                      `json:"cpuSec"`
	AggQueryd          string                   `json:"aggQueryd"`
	Logs               []map[string]interface{} `json:"logs"`
	CpuCore            int                      `json:"cpuCore"`
}

type HighCharts struct {
	Histograms []HighChart `json:"histograms"`
	Count      int         `json:"count"`
	Progress   string      `json:"progress"`
}

type HighChart struct {
	Count    uint64 `json:"count"`
	Progress string `json:"progress"`
	From     int64  `json:"from"`
	To       int64  `json:"to"`
}

type RespDatabase struct {
	DatabaseName   string `json:"databaseName"`
	InstanceName   string `json:"instanceName"`
	DatasourceType string `json:"datasourceType"`
	InstanceId     int    `json:"instanceId"`
}

type RespIndexItem struct {
	IndexName string  `json:"indexName"`
	Count     uint64  `json:"count"`
	Percent   float64 `json:"percent"`
}
