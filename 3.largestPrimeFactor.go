package main

import (
	"fmt"
	"math"
)

func main(){
	result := largestPrimeFactor(600851475143)
	fmt.Println(result)
}


func largestPrimeFactor( number int64 ) int64 {
	var maxPrime int64 = -1
	for (number % 2 == 0){
		maxPrime = 2
		number /=2
	}

	for i := 3; i <= int(math.Sqrt(float64(number))); i += 2 { 
        for number % int64(i) == 0 { 
            maxPrime = int64(i); 
            number = number / int64(i); 
        } 
	} 
	if (number > 2) {
        maxPrime = number; 
	}
  
    return maxPrime; 
}