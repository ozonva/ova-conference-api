package repo

import "ova-conference-api/internal/domain"

type Repo interface {
	AddEntities(entities []domain.Conference) error
	ListEntities(limit, offset uint64) ([]domain.Conference, error)
	DescribeEntity(entityId uint64) (*domain.Conference, error)
}
