package sort

func SelectionSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	for i := 0; i < len(data)-1; i++ {
		minIndex := 0
		for j := i + 1; j < len(data); j++ {
			if data[minIndex] > data[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			data[i], data[minIndex] = data[minIndex], data[i]
		}
	}
	return data
}
