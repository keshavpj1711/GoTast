package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 4, 5, 3, 7}

		got := Sum(numbers)
		want := 20

		if got != want {
			t.Errorf("Got %d want %d for arrays %v", got, want, numbers)
	}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 4, 5}

		got := Sum(numbers)
		want := 10

		if got != want {
			t.Errorf("Got %d want %d for arrays %v", got, want, numbers)
	}
	})
}
