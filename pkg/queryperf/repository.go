package queryperf

import "context"

//go:generate mockgen -destination=../../mocks/mock_queryperf_repository.go -mock_names=Repository=QueryPerfRepository -package=mocks github.com/artback/queryperformance/pkg/queryperf Repository
type Repository interface {
	QueriesByMeanTime(ctx context.Context, options Options) (*Result, error)
}
