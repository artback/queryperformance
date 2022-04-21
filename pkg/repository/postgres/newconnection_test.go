//go:build integration
// +build integration

package postgres_test

import (
	"context"
	"github.com/artback/queryperformance/pkg/repository/postgres"
	"testing"
)

func Test_haveExtension(t *testing.T) {
	type args struct {
		extName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "extension that exist",
			args: args{
				extName: "pg_stat_statements",
			},
		},
		{
			name: "extension that don't exist",
			args: args{
				extName: "humble_bumble",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, teardown := startPG()
			if err := postgres.HaveExtension(db, context.Background(), tt.args.extName); (err != nil) != tt.wantErr {
				t.Errorf("haveExtension() error = %v, wantErr %v", err, tt.wantErr)
			}
			err := teardown()
			if err != nil {
				return
			}
		})
	}
}
