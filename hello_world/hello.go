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

	prefix := prefixHelloEnglish

	switch language {
	case spanish:
		prefix = prefixHelloSpanish
	case french:
		prefix = prefixHelloFrench
	}

	return prefix + name
}
