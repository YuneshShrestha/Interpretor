// AST stands for Abstract Syntax Tree. We are following top-down parsing to build the AST.
/*
      *ast.Program
        |
    +----------+
    | Statements|
    +----------+
        |
  *ast.LetStatement
    +---------+
    |  Name   |
    |  Value  |
    +---------+
    /         \
*ast.Identifier   *ast.Expression


*/
package ast

import "github.com/YuneshShrestha/Interpretor/token"

type Node interface {
	// TokenLiteral() returns the literal value of the token it's associated with. Eg. for a let statement, it would return the literal value of the 'let' token.
	TokenLiteral() string
}

// statementNode() and expressionNode() are marker methods that are used to distinguish between statements and expressions for GO compilers.
type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // token.LET token
	Name  *Identifier
	Value Expression // Identifiers like x are treated as expressions because in different parts of the program, they can produce values (like in x * 10).

}

// statementNode() lets the AST know this is a statement node.
func (ls *LetStatement) statementNode() {}

type Identifier struct {
	Token token.Token // token.IDENT token
	Value string      // name of the identifier
}

// expressionNode() lets the AST know this is an expression node.
func (i *Identifier) expressionNode() {}

// TokenLiteral() allows you to get the literal value of the token (in this case, the string "let").
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
