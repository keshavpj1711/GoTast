package main

import "testing"

func TestPeriRect(t *testing.T) {
	// I know the concept of helper function has been there for a while
	// But i will be using it this time

	// This is another form of defining a function
	// checkArea := func(t testing.TB, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}
	// }

	// t.Run("rectangles", func(t *testing.T) {
	// 	rectangle := Rectangle{12, 6}
	// 	checkArea(t, rectangle, 72.0)
	// })

	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	checkArea(t, circle, 314)
	// })

	// All this above test can be simplified to 
	// To the following lines

	// This is a form of defining anonymous structs as suggested by the name it can't be used everywhere
	areaTests := []struct {
		shape Shape
		want float64
	}{
		{Rectangle{12, 3}, 36.0},
		{Circle{10}, 314},
	}
	// After defining the struct we put two values in it in following way
	for _, tt := range areaTests {
		got := tt.shape.Area()

		if got != tt.want {
			t.Errorf("Got %v want %v", got, tt.want)
		}
	}
}
