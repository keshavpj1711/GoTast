package main

import "testing"

func TestSearch(t *testing.T) {

	t.Run("known key", func(t *testing.T) {
		// Defining a map(dictionary) in GO
		dictionary := Dictionary{"test": "this is a test"}

		got,_ := dictionary.Search("test")
		want := "this is a test"

		checkStrings(t, got, want)
	})

	t.Run("unknown key", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is 2nd test"}

		_, err := dictionary.Search("test2")
		
		checkError(t, err, ErrKeyAbsent)		
	})
}

func checkError(t testing.TB, got, want error)  {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func checkStrings(t testing.TB, got string, want string)  {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}	
