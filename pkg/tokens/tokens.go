package tokens

type Token string

const (
	LEFT_PARENTHESIS Token = "LEFT_PARENTHESIS"
	RIGHT_PARETHESIS Token = "RIGHT_PARENTHESIS"
	LEFT_BRACKETS    Token = "LEFT_BRACKETS"
	RIGHT_BRACKETS   Token = "RIGHT_BRACKETS"
	COLON            Token = "COLON"
	QUOTE            Token = "QUOTE"
	COMMA            Token = "COMMA"
)

var TokenMapping = map[rune]Token{
	'{': LEFT_PARENTHESIS,
	'}': RIGHT_PARETHESIS,
	'[': LEFT_BRACKETS,
	']': RIGHT_BRACKETS,
	':': COLON,
	'"': QUOTE,
	',': COMMA,
}
