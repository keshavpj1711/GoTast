package main

import "testing"

func TestSearch(t *testing.T) {

	// Defining a map(dictionary) in GO
	dictionary := Dictionary{"test": "this is a test"}

	got,_ := dictionary.Search("test")
	want := "this is a test"

	checkStrings(t, got, want)
}

func checkStrings(t testing.TB, got string, want string)  {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}	