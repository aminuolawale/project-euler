package main

import (
	"fmt"
)


func main(){	
	result := sumSquareDifference(100)
	fmt.Println(result)
}	


func sumSquareDifference(n int) int {
	sum := 0
	for i:=1; i <=n; i++ {
		for j:=1; j<=n; j++ {
			if (i != j){
				sum +=   (i * j)
			}
		}
	}
	return sum
}