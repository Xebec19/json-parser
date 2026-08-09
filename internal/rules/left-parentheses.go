package rules

import "github.com/Xebec19/json-parser/pkg/tokens"

type LPRule struct {
	Index  int
	Tokens []tokens.Token
}

func (r *LPRule) check() bool {
	return false
}
