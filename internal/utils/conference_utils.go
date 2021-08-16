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

	chunkCount := (len(copyOfSource) + chunkSize - 1) / chunkSize
	result := make([][]domain.Conference, chunkCount)
	var rightPosition, index int

	for leftPosition := 0; leftPosition < len(conferences); leftPosition += chunkSize {
		rightPosition = min(leftPosition+chunkSize, len(copyOfSource))

		batch := copyOfSource[leftPosition:rightPosition]
		result[index] = batch
		index++
	}
	return result, nil
}

func ToMapByUserId(conferences []domain.Conference) (map[uint64]domain.Conference, error) {
	if conferences == nil {
		return nil, errors.New("input parameter is nil")
	}
	result := make(map[uint64]domain.Conference, len(conferences))
	for _, val := range conferences {
		if _, alreadyExist := result[val.UserId]; alreadyExist {
			return nil, errors.New("not unique userId key")
		}

		result[val.UserId] = val
	}
	return result, nil
}
