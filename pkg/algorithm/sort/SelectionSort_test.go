package sort

import "testing"

func TestSelectionSort(t *testing.T) {
	testData := GenerateRand()
	sortData := SelectionSort(testData)
	bubbleSortData := BubbleSort(testData)
	for i, v := range sortData {
		if bubbleSortData[i] != v {
			t.Error("SelectionSort Result Not Equal BubbleSort Result")
		}
	}
}

//go test -bench=.
func BenchmarkSelectionSort(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	testData := GenerateRand()
	b.StartTimer()             //重新开始时间
	for i := 0; i < b.N; i++ { //use b.N for looping
		SelectionSort(testData)
	}
}
