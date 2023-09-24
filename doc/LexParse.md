# Lex & Parse

## Lex
* Lexer is used to turn text `string` or file contents (whatever) into `[]*Token` array
* After that - `[]*Token` array is used for parser
* `Lex` function __NOT LOADING__ files, `filename` is used to indicate which file you are giving
    * In case you want to take them from network (For example)

```go
// filename   - Name of the file for token info
// src        - File contents (You need to load it by yourself)
//
// Returns:
// toks     - []*Tokens array after lexing
// err      - Possible errors if something went wrong
toks, err := Lex("somefile.txt", "(a 1 2 3)")
```




## Parse
```go
// Let's assume we parsed some file
toks, err := Lex("somefile.txt", fileContents)

// Now we want to parse it and take the []*Node array
nodes, err := Parse(toks)
```