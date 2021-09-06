package repo

import (
	"context"
	"ova-conference-api/internal/domain"
)

type Repo interface {
	AddEntities(ctx context.Context, entities []domain.Conference, batchCompletedCallback func([]domain.Conference)) ([]domain.Conference, error)
	AddEntity(ctx context.Context, entity domain.Conference) (*domain.Conference, error)
	ListEntities(ctx context.Context, limit, offset int64) ([]domain.Conference, error)
	DescribeEntity(ctx context.Context, entityId int64) (*domain.Conference, error)
	UpdateEntity(ctx context.Context, entity domain.Conference) error
	DeleteEntity(ctx context.Context, entityId int64) error
	Open() error
	Close() error
}

func NewRepo(connectionString string, chunkSize int) Repo {
	return &repository{connection: connectionString, chunkSize: chunkSize}
}
