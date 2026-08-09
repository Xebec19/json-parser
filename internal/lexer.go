package internal

import "github.com/Xebec19/json-parser/pkg/tokens"

// takes text and returns Token
func Lexer(text string) []tokens.Token {

	var token []tokens.Token

	for _, c := range text {

		if val, ok := tokens.TokenMapping[c]; ok {
			token = append(token, val)
		}

	}

	return token
}
