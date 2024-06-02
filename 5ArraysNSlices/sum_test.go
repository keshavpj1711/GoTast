package main

import "testing"

func TestSum(t *testing.T) {

	numbers := [5]int{1, 4, 5, 3, 7}

	got := Sum(numbers)
	want := 20

	if got != want {
		t.Errorf("Got %d want %d for arrays %v", got, want, numbers)
	}
}
