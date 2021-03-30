package hello

const spanish = "spanish"
const french = "french"
const prefixHelloEnglish = "Hello, "
const prefixHelloSpanish = "Hola, "
const prefixHelloFrench = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = prefixHelloSpanish
	case french:
		prefix = prefixHelloFrench
	default:
		prefix = prefixHelloEnglish
	}
	return
}
