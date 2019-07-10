package search

func LinearSearch(data []int, query int) int {
	for i, item := range data {
		if item == query {
			return i
		}
	}
	return -1
}
