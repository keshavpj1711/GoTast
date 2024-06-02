package main

func Repeat(strChar string) string {

	count := 4
	output := ""

	for i := 0; i < count; i++ {
		output = output + strChar
	}

	return output
}
