package main

func SumAllTails(numArrays ...[]int) (sum []int) {
	
	for _, numArr := range numArrays {
		// Theres a small problem btw 
		// When we pass empty slice there will be a runtime error stating that numArr[1:] IS trying to acess something that is out of bounds
		// So to fix this 
		if len(numArr) == 0 {
			sum = append(sum, 0)
		}else {
			sum = append(sum, Sum(numArr[1:]))
		}
	}

	return
}