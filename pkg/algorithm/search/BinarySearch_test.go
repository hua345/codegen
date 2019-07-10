package search

import (
	"codegen/pkg/algorithm/sort"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	testData := sort.GenerateRand()
	testData = sort.QuickSort(testData)
	randomIndex := sort.GetRandomIndex()
	queryIndex := BinarySearch(testData, testData[randomIndex])
	if randomIndex != queryIndex {
		t.Log("randomIndex: {} queryIndex: {}", randomIndex, queryIndex)
		t.Error("Search Failed")
	}
}

//go test -bench=.
func BenchmarkBinarySearch(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	testData := sort.GenerateRand()
	testData = sort.QuickSort(testData)
	b.StartTimer()             //重新开始时间
	for i := 0; i < b.N; i++ { //use b.N for looping
		randomIndex := sort.GetRandomIndex()
		LinearSearch(testData, testData[randomIndex])
	}
}
