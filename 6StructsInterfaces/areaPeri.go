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

func (rect Rectangle) PeriRect() float64 {
	return 2*(rect.Length + rect.Breadth)
}

func (circle Circle) AreaCircle() float64 {
	return Pi*(math.Pow(circle.Radius, 2))
}
