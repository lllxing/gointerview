package class01

import (
	"fmt"
	"gointerview/utils"
	"sort"
)

func selectSort(arr []int) {

	if arr == nil || len(arr) < 2 {
		return
	}

	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		utils.Swap(arr, i, minIndex)
	}

}

func SelectSortTest() {
	testTime := 500
	maxSize := 100
	maxValue := 100
	succeed := true

	for i := 0; i < testTime; i++ {
		arr1 := utils.GenerateRandomArray(maxSize, maxValue)
		arr2 := utils.CopyArray(arr1)

		selectSort(arr1)
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
	selectSort(arr)
	fmt.Println(arr)

}
