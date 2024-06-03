package main

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {

	numArr1 := []int{1, 4, 3}
	numArr2 := []int{0, 2}

	got := SumAllTails(numArr1, numArr2)
	want := []int{7, 2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v want %v", got, want)
	}
}
