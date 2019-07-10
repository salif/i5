package types

type Node struct {
	Kind  string
	Value string
	Body  []Node
	Dlm   string
}
