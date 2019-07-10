package sort

func MergeSort(data []int) []int {

	if len(data) < 2 {
		return data
	}
	var middle = len(data) / 2
	var a = MergeSort(data[:middle])
	var b = MergeSort(data[middle:])
	return merge(a, b)
}

func merge(left []int, right []int) []int {

	var sortedData = make([]int, len(left)+len(right))
	var i = 0
	var j = 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			sortedData[i+j] = left[i]
			i++
		} else {
			sortedData[i+j] = right[j]
			j++
		}
	}
	for i < len(left) {
		sortedData[i+j] = left[i]
		i++
	}
	for j < len(right) {
		sortedData[i+j] = right[j]
		j++
	}
	return sortedData
}
