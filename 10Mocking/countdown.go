package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(out io.Writer) {
	count := 3

	for i := 0; i < count; i++ {
		fmt.Fprintln(out, count-i)
	}

	fmt.Fprint(out, "Go!")
}

func main()  {
	Countdown(os.Stdout)
}