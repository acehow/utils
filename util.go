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
			return -1,err
		}
		left = (left*10 + number) % k
	}
	return left,nil
}

// get num1*num2 as string
func StrMul(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	sumArr := make([]int, len(num1)+len(num2))

	for i := len(num2) - 1; i >= 0; i-- {
		n2 := int(num2[i] - '0')
		for j := len(num1) - 1; j >= 0; j-- {
			n1 := int(num1[j] - '0')
			sum := n2*n1 + sumArr[i+j+1]
			sumArr[i+j+1] = sum % 10
			sumArr[i+j] += sum / 10
		}
	}

	res := ""
	for k, v := range sumArr {
		if k == 0 && v == 0 {
			continue
		}
		res += string(v + '0')
	}
	return res
}

func StrAdd(str1,str2 string) string {
	str2Count := len(str2)
	str1Count := len(str1)
	add:=0
	ans:=""
	for(str1Count>0 || str2Count>0) {
		x:=0
		if str1Count>0 {
			x=int(str1[str1Count-1]-'0')
		}
		y:=0
		if str2Count>0 {
			y=int(str2[str2Count-1]-'0')
		}
		result:= x+y+add
		ans+=strconv.Itoa(result%10)
		add=result/10
		str1Count--
		str2Count--
	}
	return StrReverse(ans)
}

func StrReverse(s string) string {
    rns := []rune(s) 
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
        rns[i], rns[j] = rns[j], rns[i]
    }
    return string(rns)
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
