package main

import (
	"fmt"
)

func main() {

	a := []int{2, 7, 11, 15}
	b := twoSum(a, 9)
	fmt.Println(b)

}

// Brute Force
// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			sum := nums[i] + nums[j]
// 			if sum == target {
// 				return []int{i, j}

// 			}

// 		}

// 	}
// 	return nil

// }

// One paas hash table
func twoSum(nums []int, target int) []int {
	numToIndexMap := make(map[int]int)

	for i, value := range nums {
		difference := target - value
		if index, found := numToIndexMap[difference]; found {
			return []int{index, i}

		}
		numToIndexMap[value] = i

	}
	return nil

}
