package internal

// takes text and returns Token
func Lexer(text string) []Token {

	var token []Token

	for _, c := range text {

		if val, ok := tokenMapping[c]; ok {
			token = append(token, val)
		}

	}

	return token
}
