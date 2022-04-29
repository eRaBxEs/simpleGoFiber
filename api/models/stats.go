package models

import "strings"

type OID struct {
	ID   int
	Type string
}

func (o OID) Name() string {
	return strings.ToUpper(o.Type)
}

type PgStatStatement struct {
	UserId            int     `json:"userid,omitempty"`
	DbId              int     `json:"dbid,omitempty"`
	QueryId           int64   `json:"queryid,omitempty"`
	Query             string  `json:"query"`
	Calls             int64   `json:"calls"`
	TotalTime         float64 `json:"total_time"`
	MinTime           float64 `json:"min_time"`
	MaxTime           float64 `json:"max_time"`
	MeanTime          float64 `json:"mean_time"`
	StddevTime        float64 `json:"stddev_time"`
	Rows              int64   `json:"rows"`
	SharedBlksHit     int64   `json:"shared_blks_hit"`
	SharedBlksRead    int64   `json:"shared_blks_read,omitempty"`
	SharedBlksDirtied int64   `json:"shared_blks_dirtied,omitempty"`
	SharedBlksWritten int64   `json:"shared_blks_written,omitempty"`
	LocalBlksHit      int64   `json:"local_blks_hit,omitempty"`
	LocalBlksRead     int64   `json:"local_blks_read,omitempty"`
	LocalBlksDirtied  int64   `json:"local_blks_dirtied,omitempty"`
	LocalBlksWritten  int64   `json:"local_blks_written,omitempty"`
	TempBlksRead      int64   `json:"temp_blks_read,omitempty"`
	TempBlksWritten   int64   `json:"temp_blks_written,omitempty"`
	BlkReadTime       float64 `json:"blk_read_time,omitempty"`
	BlkWriteTime      float64 `json:"blk_write_time,omitempty"`
}
