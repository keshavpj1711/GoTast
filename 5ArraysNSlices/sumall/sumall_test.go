package main

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {

	numArr1 := []int{1, 4, 3}
	numArr2 := []int{8, 3}

	got := SumAll(numArr1, numArr2)
	want := []int{8, 11}

	if reflect.DeepEqual(got, want) == false {
		t.Errorf("Got %v want %v", got, want)
	}
}
