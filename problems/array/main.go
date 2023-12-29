package main

import (
	"fmt"
)

func main() {

}

// 304. Range Sum Query 2D - Immutable
type NumMatrix struct {
	ps [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	n := len(matrix)    //количество массивов
	m := len(matrix[0]) // длинна массива

	ps := make([][]int, n+1) // создаем 6 пустых массивов
	for i := range ps {      // добавляем в каждый по len(matrix[0]+1
		ps[i] = make([]int, m+1)
	}
	fmt.Println(ps)

	for i := 1; i <= n; i++ { // кол-во массивов

		for j := 1; j <= m; j++ { //кол-во эл-ов

			otherSum := ps[i-1][j] + ps[i][j-1] - ps[i-1][j-1]

			ps[i][j] = matrix[i-1][j-1] + otherSum
		}
	}

	return NumMatrix{ps: ps}
}

func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	row2++
	col2++
	return nm.ps[row2][col2] - nm.ps[row1][col2] - nm.ps[row2][col1] + nm.ps[row1][col1]
}

// 560. Subarray Sum Equals K
func subarraySum(nums []int, targetSum int) int {
	prefixSums := make(map[int]int)
	prefixSums[0] = 1

	currentPrefixSum := 0
	count := 0

	for _, el := range nums { //
		currentPrefixSum += el //

		if val, ok := prefixSums[currentPrefixSum-targetSum]; ok {
			count += val //
		}

		prefixSums[currentPrefixSum]++ //
	}

	return count
}

// 724. Find Pivot Index

func pivotIndex(nums []int) int {
	numsSum := 0
	rightSum := 0
	leftSum := 0

	for _, num := range nums {
		numsSum += num
	}

	for i, num := range nums {
		rightSum = numsSum - leftSum - num
		if leftSum == rightSum {
			return i
		}
		rightSum = 0
		leftSum += num
	}
	return -1
}

// 303. Range Sum Query - Immutable
type NumArray struct {
	prefixSum []int
}

func Constructor2(nums []int) NumArray {
	n := len(nums)
	prefixSum := make([]int, n+1)
	fmt.Println(prefixSum)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}
	return NumArray{prefixSum: prefixSum}
}

func (this *NumArray) SumRange(left int, right int) int {

	return this.prefixSum[right+1] - this.prefixSum[left]
}

//724. Find Pivot Index
