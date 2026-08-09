package rules

import "github.com/Xebec19/json-parser/pkg/tokens"

type LBRule struct {
	Index  int
	Tokens []tokens.Token
}

func (r *LBRule) check() bool {
	return false
}
