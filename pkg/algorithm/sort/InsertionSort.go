package sort

func InsertionSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	for i := 1; i < len(data); i++ {
		for j := 0; j < i; j++ {
			if data[i] < data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	return data
}
