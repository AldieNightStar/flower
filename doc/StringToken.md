# String Token

## Notes
* String tokens are tokens of type `TOK_STRING` or `TOK_ATOM`
* They are enclosed inside `"` or `'` or even `` ` `` symbols, some starts with `:`

## Read `String` and `Atom` tokens
* Let's assume we have `node` of type `*Node`
```go
if node.IsValue() {
    node.Token.Type == TOK_STRING {

        // Use this function to read the string
        // Returns: string
        value := ReadString(node.Token.Value)

        // Then you can print it for example
        fmt.Println(value)
    }
}

```