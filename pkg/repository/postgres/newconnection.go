package postgres

import (
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

func HaveExtension(db *sql.DB, ctx context.Context, extName string) error {
	stmt := sq.Select("extname").From("pg_extension").Where(sq.Eq{"extname": extName})
	var extname string
	if err := stmt.RunWith(db).PlaceholderFormat(sq.Dollar).QueryRowContext(ctx).Scan(&extname); err != nil {
		return fmt.Errorf("haveExtension %w", err)
	}
	return nil
}
func NewConnection(ctx context.Context, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("NewConnection %w", err)
	}
	if err := HaveExtension(db, ctx, "pg_stat_statements"); err != nil {
		return nil, fmt.Errorf("NewConnection %w", err)
	}
	return db, nil
}
