package sort

import "testing"

func TestGenerateRand(t *testing.T) {
	testData := GenerateRand()
	if len(testData) != num {
		t.Error("Generate Random Array Length Error")
	}
	for _, v := range testData {
		if v > rangeNum {
			t.Error("Generate Random Number Range Error")
		}
	}
}
