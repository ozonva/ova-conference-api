package utils

import (
	"errors"
	"ova-conference-api/internal/domain"
)

func SplitToBulks(conferences []domain.Conference, chunkSize int) ([][]domain.Conference, error) {
	if chunkSize < 1 {
		return nil, errors.New("chunkSize should be greater then 1")
	}
	if conferences == nil {
		return nil, errors.New("source slice is nil")
	}

	copyOfSource := make([]domain.Conference, len(conferences))
	copy(copyOfSource, conferences)
	return SplitToBulksWithoutCopy(copyOfSource, chunkSize)
}

func SplitToBulksWithoutCopy(conferences []domain.Conference, chunkSize int) ([][]domain.Conference, error) {
	if chunkSize < 1 {
		return nil, errors.New("chunkSize should be greater then 1")
	}
	if conferences == nil {
		return nil, errors.New("source slice is nil")
	}

	chunkCount := (len(conferences) + chunkSize - 1) / chunkSize
	result := make([][]domain.Conference, chunkCount)
	var rightPosition, index int

	for leftPosition := 0; leftPosition < len(conferences); leftPosition += chunkSize {
		rightPosition = min(leftPosition+chunkSize, len(conferences))

		batch := conferences[leftPosition:rightPosition]
		result[index] = batch
		index++
	}
	return result, nil
}

func ToMapById(conferences []domain.Conference) (map[int64]domain.Conference, error) {
	if conferences == nil {
		return nil, errors.New("input parameter is nil")
	}
	result := make(map[int64]domain.Conference, len(conferences))
	for _, val := range conferences {
		if _, alreadyExist := result[val.Id]; alreadyExist {
			return nil, errors.New("not unique userId key")
		}

		result[val.Id] = val
	}
	return result, nil
}
