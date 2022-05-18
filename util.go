package utils

import (
	"math/rand"
	"time"
)

// get pointer of object
func NewOf[T any](obj T) *T {
	return &obj
}

// check object in the array or not
func IsIn[T comparable](target T, array []T) bool {
	for _, element := range array {
		if target == element {
			return true
		}
	}
	return false
}

// get distinct random numbers from [start,end)
func GetRandomNumber(start int, end int, count int) []int {
	if end < start || (end-start) < count {
		return nil
	}
	nums := make([]int, count)
	i := 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i < count {
		num := r.Intn((end - start)) + start
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums[i] = num
			i++
		}
	}
	return nums
}
