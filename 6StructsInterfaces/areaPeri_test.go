package main

import "testing"

func TestPeriRect(t *testing.T)  {
	t.Run("To calculate peri of rectangle", func(t *testing.T) {
		got := PeriRect(Rectangle{19, 2})
		want := 42.0

		if got != want {
			t.Errorf("Got %v want %v", got, want)
		}
	})

	t.Run("To calculate area of circle", func(t *testing.T) {
		got := AreaCircle(Circle{10})
		want := 314.0

		if got != want {
			t.Errorf("Got %v want %v", got, want)
		}
	})
}