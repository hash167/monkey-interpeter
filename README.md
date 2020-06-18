# Monkey Interpreter

Coomponents

- lexer
- parser
- Abstract Syntax Tree (AST)
- internal object system
- evaluator

## Monkey programming language

Monkey has the following features:

- C-like syntax
- variable bindings
- integers and booleans
- arithmetic expressions
- built-in functions
- first-class and higher-order functions
- closures
- a string data structure
- an array data structure
- a hash data structure

## Lexer

The lexer will take our source code as input and tokenize it.

Token data structure

```go
type struct Token {
    Type TokenType // Eg token.EQ
    Literal string // '=='
}
```

Lexer datastructure

```go
type Lexer struct {
    input        string
    position     int  // points to current char
    readPosition int  // after current char
    ch           byte // current char under examination
}
```

`NextToken()` is the most important function of the lexer. It returns the next identifed token from the input source

## Parser

Recursive descent parser - Top down operator precedence or Pratt parser

AST starts with two different types of nodes

- statement
- expression

For Monkey programming language, expressions produce values and statements do not

1. Root node consists of

    - slice of statements

    A let statements consists of two parts, an identifer and an expression

1. Node for variable bindings(let statement) consists of
    1. name of identifier
    1. value which will be an expression node
    1. Token

Overall structure of program so far

![](./static/ast_basic.png 'Title')

### recursive-descent parsing

