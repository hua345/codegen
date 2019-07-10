package sort

import (
	"math/rand"
	"time"
)

const (
	num      = 100000 // 测试数组的长度
	rangeNum = 100000 // 数组元素大小范围
)

func GenerateRand() []int {
	randSeed := rand.New(rand.NewSource(time.Now().Unix() + time.Now().UnixNano()))
	randomArr := make([]int, num)
	for i := 0; i < num; i++ {
		randomArr[i] = randSeed.Intn(rangeNum)
	}
	return randomArr
}
