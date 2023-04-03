package utils

import (
	"math/rand"
	"time"
	"strconv"
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

// get mod of num/k
func StrMod(num string, k int) (int,error) {
	left := 0
	for m := 0; m < len(num); m++ {
		number, err := strconv.Atoi(num[m : m+1])
		if err!=nil {
			return 0,err
		}
		left = (left*10 + number) % k
	}
	return left,nil
}

// get random element from slice
func GetRandElem[T any](count int, data []T, isCopy bool) []T {
	if len(data) < count {
		return nil
	}
	nums := data
	if isCopy {
		nums = make([]T, len(data))
		copy(nums, data)
	}
	
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	return nums[0:count]
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
