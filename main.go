package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	c := table(a)
	for i := 0; i < len(c); i++ {
		fmt.Println(c[i])
	}
}

func table(a int) [][]int {
	a++
	arr := make([][]int, a)
	for i := range arr {
		arr[i] = make([]int, a)
	}
	for i := 1; i < len(arr); i++ {
		arr[i][0] = i
		arr[0][i] = i
	}

	for i := 1; i < len(arr); i++ {
		for o := 1; o < len(arr); o++ {
			arr[i][o] = i * o
			arr[o][i] = i * o
		}
	}

	return arr
}
