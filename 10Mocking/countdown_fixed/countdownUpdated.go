package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Const to tell whether it slept after call or wrote
const sleep = "sleep"
const write = "write"

// Other consts
const countdownStart = 3
const word = "Go!"

// Interface to implement Sleep
type Sleeper interface {
	Sleep()
}

// SpySleeper
type SpySleeper struct {
	Calls int
}

// Default Sleeper
type DefaultSleeper struct{}

// Methods to define what each Sleeper does 
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

// To store info about number of calls and for what the call was made for
type SpySleeperOperations struct {
	Calls []string
}

func (s *SpySleeperOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpySleeperOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	
	fmt.Fprint(out, word)
}

func main()  {
	Countdown(os.Stdout, &DefaultSleeper{})
}
