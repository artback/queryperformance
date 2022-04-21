// go:build integration
//go:build integration
// +build integration

package postgres_test

import (
	"context"
	"github.com/artback/hygh/pkg/queryperf"
	"github.com/artback/hygh/pkg/repository/postgres"
	"reflect"
	"testing"
)

func TestDbPerfRepository_QueriesByMeanTime(t *testing.T) {
	t.Parallel()
	type args struct {
		options queryperf.Options
	}
	tests := []struct {
		name       string
		args       args
		wantLength int
		wantErr    bool
	}{
		{
			name:       "get all",
			args:       args{options: queryperf.Options{}},
			wantLength: 8,
		},
		{
			name:       "get all with sorting on mean_exec_time DESC",
			args:       args{options: queryperf.Options{SortBy: "mean_exec_time"}},
			wantLength: 8,
		},
		{
			name:       "get all with sorting on mean_exec_time ASC",
			args:       args{options: queryperf.Options{SortBy: "mean_exec_time", ASC: true}},
			wantLength: 8,
		},
		{
			name:       "get all with sorting on query DESC",
			args:       args{options: queryperf.Options{SortBy: "query"}},
			wantLength: 8,
		},
		{
			name:       "get all with sorting on query ASC",
			args:       args{options: queryperf.Options{SortBy: "query", ASC: true}},
			wantLength: 8,
		},
		{
			name:       "get all with offset",
			args:       args{options: queryperf.Options{Offset: 2}},
			wantLength: 6,
		},
		{
			name:       "get all with offset and limit",
			args:       args{options: queryperf.Options{Offset: 2, Limit: 3}},
			wantLength: 3,
		},
		{
			name:       "get all with minimum 2 calls",
			args:       args{options: queryperf.Options{MinCalls: 3}},
			wantLength: 1,
		},
		{
			name:       "get all SELECT",
			args:       args{options: queryperf.Options{Statements: []string{"SELECT"}}},
			wantLength: 2,
		},
		{
			name:       "get all SELECT and INSERT",
			args:       args{options: queryperf.Options{Statements: []string{"SELECT", "INSERT"}}},
			wantLength: 3,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, teardown := startPG()
			d := postgres.NewDbPerfRepository(db)
			got, err := d.QueriesByMeanTime(context.Background(), tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueriesByMeanTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(*got), tt.wantLength) {
				t.Errorf("QueriesByMeanTime() got = %v, want %v", len(*got), tt.wantLength)
			}
			err = teardown()
			if err != nil {
				return
			}
		})
	}
}
