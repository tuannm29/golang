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
<<<<<<< Updated upstream
=======

func TuanNM() int {
	return 1
}

func TuanNHu() int {
	return 2
}

func TuanNM3() int {
	return 1
}
>>>>>>> Stashed changes
