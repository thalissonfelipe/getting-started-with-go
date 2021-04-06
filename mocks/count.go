package mocks

import (
	"fmt"
	"io"
)

const lastWord = "Go!"
const startCount = 3

func Count(output io.Writer) {
	for i := startCount; i > 0; i-- {
		fmt.Fprintln(output, i)
	}
	fmt.Fprint(output, lastWord)
}
