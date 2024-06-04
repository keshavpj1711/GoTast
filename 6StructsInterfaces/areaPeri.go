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

// This interface basically allows us to call Area() method for different shapes 
// from a single piece of code depending upon what we call it for 
// Like in our case we do it for depending on what shape.Area() is called
// Is it for rectangle or circle
type Shape interface{
	Area() float64
}

func (rect Rectangle) Area() float64 {
	return rect.Length*rect.Breadth
}

func (circle Circle) Area() float64 {
	return Pi*(math.Pow(circle.Radius, 2))
}
