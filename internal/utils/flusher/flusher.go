package flusher

import (
	"context"
	"fmt"
	"ova-conference-api/internal/domain"
	"ova-conference-api/internal/utils"
	"ova-conference-api/internal/utils/repo"
)

type Flusher interface {
	Flush(ctx context.Context, entities []domain.Conference) []domain.Conference
}

func NewFlusher(chunkSize int, entityRepo repo.Repo) Flusher {
	return &flusher{
		chunkSize:  chunkSize,
		entityRepo: entityRepo,
	}
}

type flusher struct {
	chunkSize  int
	entityRepo repo.Repo
}

func (flush flusher) Flush(ctx context.Context, entities []domain.Conference) []domain.Conference {
	bulks, err := utils.SplitToBulksWithoutCopy(entities, flush.chunkSize)
	if err != nil {
		fmt.Println(fmt.Errorf("splitToBulksWithoutCopy failed + %w", err))
		return entities
	}
	var failedEntities []domain.Conference

	for _, bulk := range bulks {
		err = flush.entityRepo.AddEntities(ctx, bulk)
		if err != nil {
			fmt.Println(fmt.Errorf("repo.AddEntities failed + %w", err))
			failedEntities = addFailedEntities(failedEntities, bulk)
		}
	}
	return failedEntities
}

func addFailedEntities(collection []domain.Conference, failed []domain.Conference) []domain.Conference {
	copyOfFailed := make([]domain.Conference, len(failed))
	copy(copyOfFailed, failed)
	return append(collection, copyOfFailed...)
}
