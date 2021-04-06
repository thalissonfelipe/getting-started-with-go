package mocks

import (
	"fmt"
	"io"
)

const lastWord = "Go!"
const startCount = 3

type Sleeper interface {
	Sleep()
}

type SleeperSpy struct {
	Calls int
}

func (s *SleeperSpy) Sleep() {
	s.Calls++
}

func Count(output io.Writer, sleeper Sleeper) {
	for i := startCount; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(output, i)
	}

	sleeper.Sleep()
	fmt.Fprint(output, lastWord)
}
