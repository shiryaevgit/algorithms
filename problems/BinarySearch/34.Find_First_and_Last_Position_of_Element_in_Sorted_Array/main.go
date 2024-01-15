package main

import "fmt"

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)
	res := []int{-1, -1}

	if len(nums) == 0 {
		return res
	}

	for r-l > 1 {
		m := (l + r) / 2

		if nums[m] <= target {
			l = m
		} else {
			r = m
		}
	}
	if nums[l] == target {
		res[0] = l - 1
	} else {
		return res
	}

	for r-l > 1 {
		m := (l + r) / 2
		if nums[m] >= target {
			r = m
		} else {
			l = m
		}
	}
	res[1] = r - 1

	return res
}

func main() {
	// Исходный массив и целевые значения для поиска.
	nums := []int{5, 7, 7, 8, 8, 10}
	nums2 := []int{5, 7, 7, 8, 8, 10}
	target1 := 8
	target2 := 6

	// Вызываем функцию search для каждого целевого значения.
	result1 := searchRange(nums, target1)
	result2 := searchRange(nums2, target2)

	// Выводим результаты поиска.
	fmt.Printf("Input: nums = %v, target = %d\n", nums, target1)
	fmt.Printf("Output: %d\n", result1)

	fmt.Printf("Input: nums = %v, target = %d\n", nums, target2)
	fmt.Printf("Output: %d\n", result2)
}
