# Path words

* Some tokens like `TOK_PATH` has value like `abc.def.ghi` which need to be parsed to retreive the path

## Sample
* Let's assume we have `node` of type `*Node`
```go
if node.IsValue() {
    if node.Token.Type == TOK_PATH {

        // Use this function
        // It will return []string
        path := ReadPath(node.Token.Value)

        // You will have something like this:
        // [ "abc", "def", "ghi" ]
        fmt.Println(path)

    }
}
```