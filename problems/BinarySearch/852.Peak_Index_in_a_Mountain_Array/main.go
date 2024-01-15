package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {

	good := func(m int) bool { // может ли элемент быть вершиной
		if m == 0 {
			return true
		}
		return arr[m-1] < arr[m]
	}

	//Так как arr гарантированно mountain array, поэтому r=len(arr)-1
	// ответ будет в диапазоне [0, len(arr)-2] включительно.
	l, r := 0, len(arr)-1

	//TODO Используется стандартный цикл бинарного поиска.
	for r-l > 1 {
		m := (l + r) / 2
		if good(m) { // используем доп функцию проверки
			l = m
		} else {
			r = m
		}
	}

	return l
}

func main() {
	// Примеры ввода
	arr1 := []int{0, 1, 0}
	arr2 := []int{0, 2, 1, 0}
	arr3 := []int{0, 3, 4, 10, 5, 2}

	// Получаем результаты для каждого примера
	result1 := peakIndexInMountainArray(arr1)
	result2 := peakIndexInMountainArray(arr2)
	result3 := peakIndexInMountainArray(arr3)

	// Выводим результаты
	fmt.Printf("Input: arr = %v\n", arr1)
	fmt.Printf("Output: %d\n", result1)

	fmt.Printf("Input: arr = %v\n", arr2)
	fmt.Printf("Output: %d\n", result2)

	fmt.Printf("Input: arr = %v\n", arr3)
	fmt.Printf("Output: %d\n", result3)
}
