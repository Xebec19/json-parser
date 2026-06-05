package internal

import "strings"

// re := regexp.MustCompile(`\d+`)

// re.MatchString("abc123")
// re.Match([]byte("abc123"))

// const T = map[string]string{
// 	"Number": "Number",

// }

var SYMBOLS = map[string]string{
	"NUMBER": "NUMBER",
	"PLUS":   "PLUS",
	"SUBST":  "SUBSTRACT",
	"MULT":   "MULTIPLY",
	"DIV":    "DIVIDE",
	"LPAREN": "LPAREN",
	"RPAREN": "RPAREN",
}

type TokenList struct {
	Key string
	Val string
}

// Lexer takes a string and convert it
// to TokenList struct
func Lexer(input string) ([]TokenList, error) {

	tokens := []TokenList{}
	inputList := strings.Split(input, "")

	for _, elem := range inputList {

		if isNumber(elem) {
			tokens = append(tokens, TokenList{
				Key: SYMBOLS["NUMBER"],
				Val: elem,
			})
		} else if elem == "+" {
			tokens = append(tokens, TokenList{
				Key: SYMBOLS["PLUS"],
				Val: elem,
			})
		} else if elem == "-" {
			tokens = append(tokens, TokenList{
				Key: SYMBOLS["SUBST"],
				Val: elem,
			})
		} else if elem == "*" {
			tokens = append(tokens, TokenList{
				Key: SYMBOLS["MULT"],
				Val: elem,
			})
		} else if elem == "/" {
			tokens = append(tokens, TokenList{
				Key: SYMBOLS["DIV"],
				Val: elem,
			})
		} else if elem == "(" {
			tokens = append(tokens, TokenList{
				Key: SYMBOLS["LPAREN"],
				Val: elem,
			})
		} else if elem == ")" {
			tokens = append(tokens, TokenList{
				Key: SYMBOLS["RPAREN"],
				Val: elem,
			})
		} else {
			return nil, ErrInvalidExpression
		}

	}

	return tokens, nil
}
