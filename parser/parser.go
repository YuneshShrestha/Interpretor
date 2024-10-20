package parser

import (
	"github.com/YuneshShrestha/Interpretor/token"

	"github.com/YuneshShrestha/Interpretor/lexer"
)

type Parser struct {
	l *lexer.Lexer

	curToken token.Token // current token under examination
	peekToken  token.Token // we use this to look ahead to see what the next token is and make descision whether we are at the end of the line or if are at just the start of the arithmetic expression.
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens so curToken and peekToken are both set.
	p.nextToken()
	p.nextToken()

	return p
}
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}