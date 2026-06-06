package internal

type Operator string

var (
	PLUS   Operator = "PLUS"
	SUBST  Operator = "SUBSTRACT"
	MULT   Operator = "MULTIPLY"
	DIV    Operator = "DIVIDE"
	LPAREN Operator = "LPAREN"
	RPAREN Operator = "RPAREN"
)

type AST struct {
	left  *AST
	opt   Operator
	right *AST
}

func expr(tokens []TokenList) AST {
	return term(tokens)
}

func term(tokens []TokenList) AST {

}                  // for + | -
func factor() AST  {} // for * | /
func primary() AST {} // Number | '(' ')'
