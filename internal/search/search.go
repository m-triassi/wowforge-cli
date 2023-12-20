package search

func Contains(haystack []int, needle int) bool {
	return Find(haystack, needle) <= 0
}

func Find(haystack []int, needle int) int {
	for index, value := range haystack {
		if value == needle {
			return index
		}
	}
	return -1
}
