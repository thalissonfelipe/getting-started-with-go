package mocks

import (
	"fmt"
	"io"
	"time"
)

const lastWord = "Go!"
const startCount = 3

func Count(output io.Writer) {
	for i := startCount; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(output, i)
	}

	time.Sleep(1 * time.Second)
	fmt.Fprint(output, lastWord)
}
