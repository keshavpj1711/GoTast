package main

import "math"

const Pi float64 = 3.14

type Rectangle struct{
	Length float64
	Breadth float64
}

type Circle struct{
	Radius float64
}

func PeriRect(rect Rectangle) float64 {
	return 2*(rect.Length + rect.Breadth)
}

func AreaCircle(circle Circle) float64 {
	return Pi*(math.Pow(circle.Radius, 2))
}
