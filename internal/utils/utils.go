package utils

import "errors"

func Split(source []int, chunkSize int) ([][]int, error) {
	if chunkSize < 1 {
		return nil, errors.New("chunkSize should be greater then 1")
	}
	if source == nil {
		return nil, errors.New("source slice is nil")
	}

	copyOfSource := make([]int, len(source))
	copy(copyOfSource, source)

	chunkCount := (len(copyOfSource) + chunkSize - 1) / chunkSize
	result := make([][]int, chunkCount)
	var rightPosition, index int

	for leftPosition := 0; leftPosition < len(source); leftPosition += chunkSize {
		rightPosition = min(leftPosition+chunkSize, len(copyOfSource))

		batch := copyOfSource[leftPosition:rightPosition]
		result[index] = batch
		index++
	}

	return result, nil
}

func SwapKeyValues(source map[string]string) (map[string]string, error) {
	if source == nil {
		return nil, errors.New("source is nil")
	}
	resultMap := make(map[string]string, len(source))

	for k, v := range source {
		resultMap[v] = k
	}
	return resultMap, nil
}

func Disjoin(source []string, filterCollection []string) ([]string, error) {
	if source == nil {
		return nil, errors.New("source slice is nil")
	}
	var result []string
	cachedFilter := getMapFromSlice(filterCollection)

	for i := 0; i < len(source); i++ {
		if _, ok := cachedFilter[source[i]]; ok {
			continue
		}
		result = append(result, source[i])
	}
	return result, nil
}

func getMapFromSlice(source []string) map[string]bool {
	if source == nil {
		return map[string]bool{}
	}
	result := make(map[string]bool, len(source))
	for i := 0; i < len(source); i++ {
		result[source[i]] = true
	}
	return result
}

func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}
