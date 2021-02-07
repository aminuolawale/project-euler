package main 


import (
	"fmt"
	"strconv"
)

func main(){
	d := largestPalindromeProduct()
	fmt.Println(d)
}

func largestPalindromeProduct() int {
	var product int
	var max int
	for num1:=1000; num1 >0; num1 -- {
		for num2:= 1000; num2 >0; num2 -- {
			product = num1 * num2
			if isPalindrome(product){
				if max < product{
					max = product
				}
			}
		}
	}
	return max
}

func isPalindrome(number int) bool {
	stringNumber := strconv.Itoa(number)
	return stringNumber == reverseString(stringNumber)
}

func reverseString(s string) string{
	out := ""
	for i:= len(s)-1; i >-1; i-- {
		out +=string( s[i])
	}
	return out
}