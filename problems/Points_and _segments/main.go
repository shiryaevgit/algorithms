package main

import (
	"fmt"
	"sort"
)

func solve(A [][]int) int {

	points := make([][]int, 0, len(A))

	// заполняем points
	for _, elem := range A {
		points = append(points, []int{elem[0], 1})
		points = append(points, []int{elem[1], -1})
	}

	// сортируем по времени
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	maxRoomNum := 0
	currRoomNum := 0

	//
	for _, point := range points {
		currRoomNum += point[1]
		if currRoomNum > maxRoomNum {
			maxRoomNum = currRoomNum
		}
	}

	return maxRoomNum
}

func main() {
	A := [][]int{{7, 10}, {4, 19}, {19, 26}, {14, 16}, {13, 18}, {16, 21}}
	result := solve(A)
	fmt.Println(result) // Вывод: 2

	//B := [][]int{{1, 18}, {18, 23}, {15, 29}, {4, 15}, {2, 11}, {5, 13}}
	//result = solve(B)
	//fmt.Println(result) // Вывод: 4

}
