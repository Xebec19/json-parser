package rules

import "github.com/Xebec19/json-parser/pkg/tokens"

type RPRule struct {
	Index  int
	Tokens []tokens.Token
}

func (r *RPRule) check() bool {
	return false
}
