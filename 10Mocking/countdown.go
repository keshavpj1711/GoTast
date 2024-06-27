package main

import (
	"fmt"
	"time"
)

func countdown(n int)  {
	second := time.Second

	for i := 0; i < n; i++ {
		fmt.Println(n-i)
		time.Sleep(second)
	}
	fmt.Println("Go!")
}

func main()  {
	countdown(3)
}