package hello

import "testing"

func TestHello(t *testing.T) {
	verifyCorrectMessage := func(t *testing.T, result, expect string) {
		t.Helper()
		if result != expect {
			t.Errorf("result '%s', expect '%s'", result, expect)
		}
	}

	t.Run("should say hello to people", func(t *testing.T) {
		result := Hello("Felipe", "")
		expect := "Hello, Felipe"
		verifyCorrectMessage(t, result, expect)
	})

	t.Run("should say 'Hello, world' when an empty string is provided", func(t *testing.T) {
		result := Hello("", "")
		expect := "Hello, world"
		verifyCorrectMessage(t, result, expect)
	})

	t.Run("should say a hello message in spanish", func(t *testing.T) {
		result := Hello("Felipe", "spanish")
		expect := "Hola, Felipe"
		verifyCorrectMessage(t, result, expect)
	})

	t.Run("should say a hello message in french", func(t *testing.T) {
		result := Hello("Felipe", "french")
		expect := "Bonjour, Felipe"
		verifyCorrectMessage(t, result, expect)
	})
}
