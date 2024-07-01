package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// Let's define our dependency as an interface.
// This lets us then use a real Sleeper in main and a spy sleeper in our tests.
type Sleeper interface {
	Sleep()
}

// So we created two Sleeper one to be called in main function that works as normal Sleeper
// Other one to be used in tests as it's just calling sleep with out any time passing and increasing call counts
type SpySleeper struct {
	Calls int
}
type DefaultSleeper struct{}

// A function to implement default sleeping when called in main
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)

}

// A function to call spySleeper when used in tests and everytime sleep is called it inc calls
func (s *SpySleeper) Sleep() {
	s.Calls++
}

func Countdown(out io.Writer, sleeper Sleeper) {

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
