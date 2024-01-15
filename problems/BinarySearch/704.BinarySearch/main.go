package main

import "fmt"

func search(nums []int, target int) int {
	// Инициализация переменных l (левая граница) и r (правая граница).
	l, r := 0, len(nums)

	for r-l > 1 { // идем по циклу пока указатели не будут рядом

		m := (l + r) / 2 // Находим середину массива.

		if nums[m] <= target { // находим новую зону поиска, сдвигаем один из указателей
			l = m
		} else {
			r = m
		}
	}

	if nums[l] == target { // на выходе, когда указатели "соседи", сверяем левый с таргет
		return l
	}
	return -1
}

func main() {
	// Исходный массив и целевые значения для поиска.
	nums := []int{-1, 0, 3, 5, 9, 12}
	target1 := 9
	target2 := 2

	// Вызываем функцию search для каждого целевого значения.
	result1 := search(nums, target1)
	result2 := search(nums, target2)

	// Выводим результаты поиска.
	fmt.Printf("Input: nums = %v, target = %d\n", nums, target1)
	fmt.Printf("Output: %d\n", result1)

	fmt.Printf("Input: nums = %v, target = %d\n", nums, target2)
	fmt.Printf("Output: %d\n", result2)
}
