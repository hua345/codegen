package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	testData := GenerateRand()
	sortData := BubbleSort(testData)
	mergeSortData := MergeSort(testData)
	for i, v := range sortData {
		if mergeSortData[i] != v {
			t.Error("BubbleSort Result Not Equal MergeSort Result")
		}
	}
}

//go test -bench=.
func BenchmarkBubbleSort(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	testData := GenerateRand()
	b.StartTimer()             //重新开始时间
	for i := 0; i < b.N; i++ { //use b.N for looping
		BubbleSort(testData)
	}
}
