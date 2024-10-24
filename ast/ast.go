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

import (
	"bytes"

	"github.com/YuneshShrestha/Interpretor/token"
)

type Node interface {
	// TokenLiteral() returns the literal value of the token it's associated with. Eg. for a let statement, it would return the literal value of the 'let' token.
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type LetStatement struct {
	Token token.Token // token.LET token
	Name  *Identifier
	Value Expression // Identifiers like x are treated as expressions because in different parts of the program, they can produce values (like in x * 10).

}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// statementNode() lets the AST know this is a statement node.
func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

type Identifier struct {
	Token token.Token // token.IDENT token
	Value string      // name of the identifier

}

// expressionNode() lets the AST know this is an expression node.
func (i *Identifier) expressionNode() {}

// TokenLiteral() allows you to get the literal value of the token (in this case, the string "let").
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) String() string { return i.Value }

type ReturnStatement struct {
	Token       token.Token // token.RETURN token
	ReturnValue Expression  // 6*7, 10, x, etc.
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token // The first token of the expression
	Expression Expression
}
func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
