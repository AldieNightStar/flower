# Node

* Node is a simple core building block of the `flower` language
* Node could be either `IsValue()` or `IsContainer()`
* Value nodes is a nodes which can have `Token` value
* Container nodes could have `Children` nodes inside. Otherwise it always `nil`

## Structure
```go
type Node struct {
	Token    *Token
	Children []*Node
}
```


## API
* Let's assume that we have `node` as `*Node`
```go
// Create a new Node which has `Children` array to append new nodes
node := NewChildrenNode()

// Create token Node (Has `Children` as `nil` value)
node := NewTokenNode(token *Token)

// Check that this node is a value
// If yes then you could get node.Token
node.IsValue() bool

// Check that this node is container (Has nodes inside)
// If yes then you could get node.Children (Array of containing nodes)
node.IsContainer() bool

// Get *Token of the Node (If: IsValue() == true)
if node.IsValue() {
    // Get token here (Check 'Token' for more details)
    node.Token
}


// Get []*Node (If: IsContainer() == true)
if node.IsContainer() {
    // Get children here (Array)
    node.Children
}
```
[More about 'Token'](Token.md)
