package main

import (
	"fmt"
)

func main(){
	var ans int64 = evenFibonacciNumbers()
	fmt.Println(ans)
}

func evenFibonacciNumbers() int64{
	var prev int64 = 1
	var curr int64 = 2
	var sum int64 = 0
	for curr < 4000000 {
		if curr % 2 == 0 {
			sum += curr
		}
		temp := curr
		curr += prev
		prev = temp
	}
	return sum
}