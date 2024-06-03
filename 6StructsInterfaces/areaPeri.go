package main

type Rectangle struct{
	Length float64
	Breadth float64
}

func PeriRect(rect Rectangle) float64 {
	return 2*(rect.Length + rect.Breadth)
}
