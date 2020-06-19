package ast

import (
	"bytes"
	"monkey/token"
)

// Node is the basic interface of the AST
type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// ReturnStatement returns an expression
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

// ExpressionStatement is a wrapper around Expression Struct
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral returns Token.Literal value
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Value.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral returns Token.Literal value
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral returns Token.Literal value
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// Identifier object stores the token object and string value of Identifer
type Identifier struct {
	Token token.Token // token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns Token.Literal value
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// IntegerLiteral is of the type Expression
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

// TokenLiteral satisfies the Expression Node interface
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}
func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) String() string  { return il.Token.Literal }
