package main

import "testing"

func TestPeriRect(t *testing.T)  {
	got := PeriRect(Rectangle{19, 2})
	want := 42.0

	if got != want {
		t.Errorf("Got %v want %v", got, want)
	}
}