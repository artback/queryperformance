package postgres

import (
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/artback/hygh/pkg/queryperf"
	"strings"
)

type dbPerfRepository struct {
	DB *sql.DB
}

func (d dbPerfRepository) QueriesByMeanTime(ctx context.Context, options queryperf.Options) (*queryperf.Result, error) {
	stmt := sq.Select("query,calls,mean_exec_time,total_exec_time,stddev_exec_time,rows").From("pg_stat_statements").Offset(options.Offset)
	if options.Limit > 0 {
		stmt = stmt.Limit(options.Limit)
	}
	if options.MinCalls > 0 {
		stmt = stmt.Where(sq.GtOrEq{"calls": options.MinCalls})
	}
	if len(options.Statements) > 0 {
		var commands []string
		for _, c := range options.Statements {
			commands = append(commands, fmt.Sprintf(`Lower(query) LIKE Lower('%%%s%%')`, c))
		}
		stmt = stmt.Where(strings.Join(commands, " OR "))
	}
	if options.SortBy != "" {
		var orderBy string
		if options.ASC == false {
			orderBy = fmt.Sprintf("%s desc", options.SortBy)
		} else {
			orderBy = fmt.Sprintf("%s asc", options.SortBy)
		}
		stmt = stmt.OrderBy(orderBy)
	}
	rows, err := stmt.RunWith(d.DB).PlaceholderFormat(sq.Dollar).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("QueriesByMeanTime %w", err)
	}
	defer rows.Close()

	var result queryperf.Result
	for rows.Next() {
		var m queryperf.Measure
		if err := rows.Scan(&m.Query, &m.Calls, &m.MeanExecTime, &m.TotalExecTime, &m.StdDevExecTime, &m.NrOfRows); err != nil {
			return &result, fmt.Errorf("QueriesByMeanTime %w", err)
		}
		result = append(result, m)
	}

	return &result, rows.Err()
}

func NewDbPerfRepository(db *sql.DB) *dbPerfRepository {
	return &dbPerfRepository{DB: db}
}
