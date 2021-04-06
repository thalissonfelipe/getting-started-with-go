package mocks

import (
	"bytes"
	"fmt"
)

func Count(output *bytes.Buffer) {
	fmt.Fprint(output, "3")
}
