package main

import (
	"fmt"
	"math"
)



func main(){
	// the nth prime number obeys the following formula:
	// n*ln(n) + n(ln(ln(n)) − 1) < P(n) < n* ln(n) +n *ln(ln(n)) for n >= 6
	// I use the upperbound to generate an array and filter all the primes within that range with the
	// Seive of Erastosthenes
	// Return the 10001th prime in the array

	var upperBound int64 = 10001 * (int64(math.Log(10001)) + int64(math.Log(math.Log(10001))))
	numbers := createSlice(2, upperBound)
	var index int64= 0
	for {
		numbers = cleanMultiplesOf(numbers[index], numbers)
		index = nextNonZero(index, numbers)
		if (index == -1){
			break
		}
	}
	fmt.Println(numbers[0])
	count :=0

	for i:= 0; ; i++ {
		if (numbers[i]!= 0){
			count ++
			if (count ==10001){
				fmt.Println(numbers[i], "is the ", count, "th prime number")
				break
			}
		}
	}

}

func nextNonZero(index int64, arr []int64) int64 {
	for i:= index+1; i < int64(len(arr)) ; i++ {
		if (arr[i]!=0){
			return i
		}
	}
	return -1
}

func cleanMultiplesOf(n int64, arr []int64) []int64{
	for i, v := range arr {
		if (v%n ==0 && v!=n){
			arr[i] = 0
		}
	}
	return arr
}

func createSlice(start int64, end int64) []int64 {
	result := []int64{}
	for i:=start; i<= end; i++ {
		result = append(result, i)
	}
	return result
}

func countNonZeros(arr []int64) int {
	count :=0
	for i:=0; i< len(arr); i++{
		if(arr[i]>0){
			count ++
		}
	}
	return count
}