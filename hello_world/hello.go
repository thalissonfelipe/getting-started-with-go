package hello

const spanish = "spanish"
const prefixHelloEnglish = "Hello, "
const prefixHelloSpanish = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == spanish {
		return prefixHelloSpanish + name
	}

	return prefixHelloEnglish + name
}
