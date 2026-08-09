package rules

import "github.com/Xebec19/json-parser/pkg/tokens"

type RBRule struct {
	Index  int
	Tokens []tokens.Token
}

func (r *RBRule) check() bool {
	return false
}
