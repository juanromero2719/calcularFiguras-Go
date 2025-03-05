package funciones

type palindromo struct {
	palabra string
}

func ValidarPalindromo(palabra string) bool {
	palabraAlReves := reverse(palabra)
	return isPalindrome(palabra, palabraAlReves)
}

func reverse(palabra string) string {

	runes := []rune(palabra)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindrome(palabra, palabraAlReves string) bool {
	return palabra == palabraAlReves
}
