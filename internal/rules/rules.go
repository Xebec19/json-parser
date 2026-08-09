package rules

import "github.com/Xebec19/json-parser/pkg/tokens"

type TokenRule interface {
	check() bool
}

func ApplyRule(index int, t []tokens.Token) bool {
	return getCheckerFunc(index, t).check()
}

func getCheckerFunc(index int, t []tokens.Token) TokenRule {

	switch t[index] {
	case tokens.LEFT_PARENTHESIS:
		return &LPRule{
			Index:  index,
			Tokens: t,
		}

	case tokens.RIGHT_PARETHESIS:
		return &RPRule{
			Index:  index,
			Tokens: t,
		}

	case tokens.LEFT_BRACKETS:
		return &LBRule{
			Index:  index,
			Tokens: t,
		}

	case tokens.RIGHT_BRACKETS:
		return &RBRule{
			Index:  index,
			Tokens: t,
		}

	case tokens.COLON:
		return &ColonRule{
			Index:  index,
			Tokens: t,
		}

	case tokens.QUOTE:
		return &QuoteRule{
			Index:  index,
			Tokens: t,
		}

	case tokens.COMMA:
		return &CommaRule{
			Index:  index,
			Tokens: t,
		}
	}

	panic("invalid token")
}
