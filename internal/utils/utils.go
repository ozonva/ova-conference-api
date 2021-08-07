package utils

func Split(source []int, chunkSize int) [][]int {
	chunkCount := (len(source) + chunkSize - 1) / chunkSize
	result := make([][]int, chunkCount)
	var rightPosition, index int

	for leftPosition := 0; leftPosition < len(source); leftPosition += chunkSize {
		rightPosition = min(leftPosition+chunkSize, len(source))

		batch := source[leftPosition:rightPosition]
		result[index] = batch
		index++
	}

	return result
}

func SwapKeyValues(source map[string]string) map[string]string {
	resultMap := make(map[string]string, len(source))

	for k, v := range source {
		resultMap[v] = k
	}
	return resultMap
}

func Disjoin(source []string, filterCollection []string) []string {
	var result []string

	cachedFilter := getMapFromSlice(filterCollection)
	for i := 0; i < len(source); i++ {

		if _, ok := cachedFilter[source[i]]; !ok {
			result = append(result, source[i])
		}
	}
	return result
}

func getMapFromSlice(source []string) map[string]bool {
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
