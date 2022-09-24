package timescaledb

import (
	"time"
	
	"github.com/cnosdb/tsdb-comparisons/cmd/generate_queries/uses/iot"
	"github.com/cnosdb/tsdb-comparisons/cmd/generate_queries/utils"
	"github.com/cnosdb/tsdb-comparisons/pkg/query"
)

const goTimeFmt = "2006-01-02 15:04:05.999999 -0700"

// BaseGenerator contains settings specific for TimescaleDB
type BaseGenerator struct {
	UseJSON       bool
	UseTags       bool
	UseTimeBucket bool
}

// GenerateEmptyQuery returns an empty query.TimescaleDB.
func (g *BaseGenerator) GenerateEmptyQuery() query.Query {
	return query.NewTimescaleDB()
}

// fillInQuery fills the query struct with data.
func (g *BaseGenerator) fillInQuery(qi query.Query, humanLabel, humanDesc, table, sql string) {
	q := qi.(*query.TimescaleDB)
	q.HumanLabel = []byte(humanLabel)
	q.HumanDescription = []byte(humanDesc)
	q.Hypertable = []byte(table)
	q.SqlQuery = []byte(sql)
}

// NewIoT creates a new iot use case query generator.
func (g *BaseGenerator) NewIoT(start, end time.Time, scale int) (utils.QueryGenerator, error) {
	core, err := iot.NewCore(start, end, scale)
	
	if err != nil {
		return nil, err
	}
	
	iot := &IoT{
		BaseGenerator: g,
		Core:          core,
	}
	
	return iot, nil
}
