package parser

type Node struct {
	Token    *Token
	Children []*Node
}

func NewTokenNode(token *Token) *Node {
	return &Node{
		Token:    token,
		Children: nil,
	}
}

func (n *Node) IsValue() bool {
	return n.Token != nil
}

func (n *Node) IsContainer() bool {
	return n.Children != nil
}

func (n *Node) addsubnode(node *Node) {
	n.Children = append(n.Children, node)
}

func (n *Node) removeLast() {
	n.Children = n.Children[0 : len(n.Children)-1]
}

func NewChildrenNode() *Node {
	return &Node{
		Token:    nil,
		Children: []*Node{},
	}
}
