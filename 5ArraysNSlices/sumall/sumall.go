package main

func SumAll(numArr ...[]int) []int {

	lengthOfNumArr := len(numArr)
	sumArr := make([]int, lengthOfNumArr)

	for i, number := range numArr {
		sumArr[i] += Sum(number)	
	}
	
	return sumArr
}