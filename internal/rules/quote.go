package rules

import "github.com/Xebec19/json-parser/pkg/tokens"

type QuoteRule struct {
	Index  int
	Tokens []tokens.Token
}

func (r *QuoteRule) check() bool {
	return false
}
