package main

import "fmt"

func main() {
	arr := twoSum([]int{3, 2, 3, 7, 8}, 10)

	fmt.Println(arr)

	array := twoSumWithOneLoop([]int{3, 2, 3, 7, 8}, 10)
	fmt.Println(array)
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func twoSumWithOneLoop(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		index, ok := m[complement]
		if ok {
			return []int{index, i}
		}
		m[nums[i]] = i
	}

	return nil
}
