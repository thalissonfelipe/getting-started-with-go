package mocks

import (
	"fmt"
	"io"
)

func Count(output io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(output, i)
	}
	fmt.Fprint(output, "Go!")
}
