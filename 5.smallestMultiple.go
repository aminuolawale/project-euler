package main

import (
	"fmt"
)


func main(){
	result := largestMultiple(20)
	fmt.Println(result)
}


func largestMultiple( n int)int {
	numbers := buildArray(1, n)
	ans := numbers[0]
	for i:=1; i< len(numbers); i++{
		ans = (numbers[i] * ans)/(gcd(numbers[i], ans))
	}
	return ans
}


func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func buildArray( start int, stop int)[]int{
	result :=[]int{}
	for i:=start; i<=stop; i++{
		result = append(result, i)
	}
	return result
}