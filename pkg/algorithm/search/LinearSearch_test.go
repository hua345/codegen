package search

import (
	"codegen/pkg/algorithm/sort"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	testData := sort.GenerateRand()
	randomIndex := sort.GetRandomIndex()
	queryIndex := LinearSearch(testData, testData[randomIndex])
	if randomIndex != queryIndex {
		t.Error("Search Failed")
	}
}

//go test -bench=.
func BenchmarkLinearSearch(b *testing.B) {
	b.StopTimer() //停止压力测试的时间计数
	testData := sort.GenerateRand()
	b.StartTimer()             //重新开始时间
	for i := 0; i < b.N; i++ { //use b.N for looping
		randomIndex := sort.GetRandomIndex()
		LinearSearch(testData, testData[randomIndex])
	}
}
