package search

func Contains(haystack []int, needle int) bool {
	for _, value := range haystack {
		if value == needle {
			return true
		}
	}
	return false
}
