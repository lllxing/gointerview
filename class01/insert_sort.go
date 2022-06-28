package class01

import (
	"fmt"
	"gointerview/utils"
	"sort"
)

func insertSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				utils.Swap(arr, j, j+1)
			}
		}
	}
}


func InsertSortTest()  {
	testTime := 50
	maxSize := 10
	maxValue := 10
	succeed := true

	for i := 0; i < testTime; i++ {
		arr1 := utils.GenerateRandomArray(maxSize, maxValue)
		arr2 := utils.CopyArray(arr1)

		insertSort(arr1)
		sort.Ints(arr2)

		if !utils.IsEqual(arr1, arr2) {
			succeed = false
			fmt.Println(arr1)
			fmt.Println(arr2)
			break
		}
		fmt.Println(succeed)

	}

	if succeed {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	arr := utils.GenerateRandomArray(maxSize, maxValue)
	fmt.Println(arr)
	insertSort(arr)
	fmt.Println(arr)

}