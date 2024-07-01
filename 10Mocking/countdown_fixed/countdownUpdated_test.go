package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("printing 3 to GO!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if spySleeper.Calls != 3 {
			t.Errorf("not enough calls to sleeper, got %v want %v", spySleeper.Calls, countdownStart)
		}
	})

	t.Run("checking sleep runs after every count", func(t *testing.T) {
		spySleeperPrinter := &SpySleeperOperations{}
		Countdown(spySleeperPrinter, spySleeperPrinter)

		want := []string{write, sleep, write, sleep, write, sleep, write}
		got := spySleeperPrinter.Calls

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
		
	})
}
