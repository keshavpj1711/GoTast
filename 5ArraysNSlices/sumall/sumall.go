package main

func SumAll(arr1, arr2 []int) []int {

	sum := []int{0, 0}

	for _, v := range arr1 {
		sum[0] += v
	}
	for _, v := range arr2 {
		sum[1] += v
	}
	
	return sum
}