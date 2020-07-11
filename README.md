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

## golang

### Testing

We are using a table driven testing approach because

- Setup code for the tests are the same, so we are avoiding a lot of duplication
- Logic behind the test assertions are the same

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

![Basic Structure](./static/ast_basic.png 'Title')

### recursive-descent parsing

Top down operator precedence or Pratt Parsing

Invented as an alternate to context free grammers and the Backus_Nauer_Form

Instead of associating parsing functions with grammer rules, Pratt associated parsing functions(sematic code) with single token types. Each token type can have two parsing functions associated with it depending on the token's position(infix or prefix)


### Expressions in Monkey

Everything in monkey is an expression besides the  `let` and `return` statements.

Expressions have the following operators

1. Prefix operators(Eg -5, !5)
1. Infix or binary operators(Eg 5 + 5)
1. Comparison operators(Eg 5 == 5)

- We can also use paranthesis to group expressions and influence the order of evaluation.

Eg (5 + 5) or ((5 + 5) * 5)

- call expressions Eg max(5 + add(5 + 5))
- Functions in monkey are first class citizens
- function literals are expressions
- We can use the the let statement to bind functions to names
- `let add = fn(x, y) { return x + y };` The function literal is the just the expression in the statement

Implementation of the Pratt Parser

- Associtation of parsing functions with token types
- When this token type is encountered, the parsing functions are called to parse the appropriate expression and return an AST node that represents it
- Each token type can have upto 2 parsing functions associated with it, depending on whether the token was found in a prefix or an infix position

```go
// parser/parser.go
type (
prefixParseFn func() ast.Expression
infixParseFn func(ast.Expression) ast.Expression
)
```

Identifiers are the first expression we will try to parse

