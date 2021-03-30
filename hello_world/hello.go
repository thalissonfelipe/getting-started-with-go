package hello

const spanish = "spanish"
const french = "french"
const prefixHelloEnglish = "Hello, "
const prefixHelloSpanish = "Hola, "
const prefixHelloFrench = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == spanish {
		return prefixHelloSpanish + name
	}

	if language == french {
		return prefixHelloFrench + name
	}

	return prefixHelloEnglish + name
}
