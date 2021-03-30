package iteration

// Iteration receives an character and number as parameters and returns
// a string containing "n" character
func Iteration(character string, number int) string {
	var repetitions string
	for i := 0; i < number; i++ {
		repetitions += character
	}
	return repetitions
}
