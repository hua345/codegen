package sort

import "math/rand"

func QuickSort(data []int) []int {

	if len(data) <= 1 {
		return data
	}

	median := data[rand.Intn(len(data))]

	low_part := make([]int, 0, len(data))
	high_part := make([]int, 0, len(data))
	middle_part := make([]int, 0, len(data))

	for _, item := range data {
		switch {
		case item < median:
			low_part = append(low_part, item)
		case item == median:
			middle_part = append(middle_part, item)
		case item > median:
			high_part = append(high_part, item)
		}
	}

	low_part = QuickSort(low_part)
	high_part = QuickSort(high_part)

	low_part = append(low_part, middle_part...)
	low_part = append(low_part, high_part...)

	return low_part
}
