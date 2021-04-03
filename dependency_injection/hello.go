package dependency_injection

import (
	"fmt"
	"io"
)

func Hello(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
