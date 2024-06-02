package main

func Repeat(strChar string, count int) string {

	output := ""

	for i := 0; i < count; i++ {
		output += strChar
	}

	return output
}
