package ast

import "monkey/token"

// Node is the basic interface of the AST
type Node interface {
	TokenLiteral() string
}

// Statement is the interface that inherits from Node and will form the basis of all statements of our AST
type Statement interface {
	Node
	statementNode()
}

// Expression is TODO
type Expression interface {
	Node
	expressionNode()
}

// Program This is going to be the root node of the AST
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the Literal value of the Token struct
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// ReturnStatement returns and expression
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

// LetStatement has a name and expression as value
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral returns Token.Literal value
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral returns Token.Literal value
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// Identifier object stores the token object and string value of Identifer
type Identifier struct {
	Token token.Token // token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns Token.Literal value
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
