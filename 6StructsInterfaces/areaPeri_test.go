package main

import "testing"

func TestPeriRect(t *testing.T)  {
	t.Run("To calculate peri of rectangle", func(t *testing.T) {
		rectangle := Rectangle{19, 2}
		got := rectangle.PeriRect()
		want := 42.0

		if got != want {
			t.Errorf("Got %v want %v", got, want)
		}
	})

	t.Run("To calculate area of circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.AreaCircle()
		want := 314.0

		if got != want {
			t.Errorf("Got %v want %v", got, want)
		}
	})
}