package rules

import "github.com/Xebec19/json-parser/pkg/tokens"

type ColonRule struct {
	Index  int
	Tokens []tokens.Token
}

func (r *ColonRule) check() bool {
	return false
}
