package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomArray(maxSize int, maxValue int) []int {

	rand.Seed(time.Now().UnixNano())
	len := rand.Intn(maxValue) * (maxSize + 1)

	var array []int

	for i := 0; i < len; i++ {
		array = append(array, rand.Intn(maxValue)*(maxSize+1)-rand.Intn(maxValue)*maxSize)
	}

	return array
}

func CopyArray(arr []int) []int {
	if arr == nil {
		return nil
	}

	len := len(arr)

	var array []int

	for i := 0; i < len; i++ {
		array = append(array, arr[i])
	}

	return array
}

func IsEqual(arr1 []int, arr2 []int) bool {
	if (arr1 == nil && arr2 != nil) || (arr2 == nil && arr1 != nil) {
		return false
	}

	if arr1 == nil && arr2 == nil {
		return true
	}

	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func Swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
