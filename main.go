package main

import "fmt"

func main() {
	fmt.Print(Sum(10))
}

func Sum(num int) int {
	sum := 0
	for i := 0; i < num; i++ {
		sum += i
	}
	return sum
}

func TuanNM() int {
	return 1
}
