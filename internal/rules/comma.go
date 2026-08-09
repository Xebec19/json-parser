package rules

import "github.com/Xebec19/json-parser/pkg/tokens"

type CommaRule struct {
	Index  int
	Tokens []tokens.Token
}

func (r *CommaRule) check() bool {
	return false
}
