# Token

## Structure
* Let's assume we have `token` variable of `*Token` type
```go
// Get token info (FileName:line)
// For example:   file1.txt:13
token.Info

// Token value (As is)
// String values are enclosed inside ' ' symbols
token.Value

// Token type
// For example:  TOK_STRING
// You can check another TOK_* constants
token.Type
```


## Token Types
* Token type is `TokenType` (`uint8` value)
```go
const (
    // Number token
    // Example: 123  -398  12.44
	TOK_NUMBER

    // String token
    // Example: "Hello there" 'hi' `string`
	TOK_STRING

    // Symbols token
    // Example:  %%==> 
	TOK_SYMBOLS

    // Atom tokens
    // Example:  :key  :value  :reset
	TOK_ATOM

    // Word tokens
    // Example:  variable   step  abc
	TOK_WORD

    // Bracket tokens
    // Example:  (  )
	TOK_BRACKET

    // Path token 
    // Example:  test.abc.trigger
	TOK_PATH
)
```