package main

func Sum(num_array [5]int) (sum int) {

	// Required for normal for loop
	// count := len(num_array)

	// This is a simple way to iterate over arrays 
	// for i := 0; i < count; i++ {
	// 	sum += num_array[i]
	// }

	// Introducing range to iterate over arrays
	for _, number := range num_array {
		sum += number
	}

	return
}
