package main


import (
	"fmt"
)

func main(){
	ans:= multiplesOfThreeAndFive()
	fmt.Println(ans)

}

func multiplesOfThreeAndFive() int{
	sum := 0
	for i:=3; i<1000; i++ {
		if (i%3 ==0 || i % 5 == 0) {
			sum += i
		}
	}
	return sum
}