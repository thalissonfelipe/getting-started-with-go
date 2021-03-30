package iteration

func Iteration(character string, number int) string {
	var repetitions string
	for i := 0; i < number; i++ {
		repetitions += character
	}
	return repetitions
}
