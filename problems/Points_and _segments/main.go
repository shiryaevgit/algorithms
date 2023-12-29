package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := 1000000
	nCPU := runtime.NumCPU()
	slice := makeSlice(n)
	slices := splitSlice(slice, nCPU)
	channels := distribution(slices)

	sum := 0
	for _, ch := range channels {
		sum += <-ch
	}
	fmt.Println("Результат:", sum)

}

func calculation(slice []int, ch chan int) {
	defer close(ch)
	sum := 0
	for _, v := range slice {
		sum += v
	}
	ch <- sum
}

func distribution(slices [][]int) []chan int {
	channels := make([]chan int, 0, len(slices))

	for i := 0; i < len(slices); i++ {
		ch := make(chan int)
		go calculation(slices[i], ch)
		channels = append(channels, ch)
	}
	return channels
}

func makeSlice(n int) []int {
	res := make([]int, n)
	count := 0
	sum := 0
	for i := 0; i < n; i++ {
		res[i] = i + 1
		//res[i] = 1
		sum += res[i]
		count++
	}
	return res
}

func splitSlice(slice []int, nCPU int) [][]int {
	cutSize := len(slice) / nCPU
	res := make([][]int, 0, nCPU)

	for i := 1; i <= nCPU; i++ {
		if i == nCPU {
			res = append(res, slice)
			break
		}
		res = append(res, slice[:cutSize])
		slice = slice[cutSize:]
	}
	return res
}
