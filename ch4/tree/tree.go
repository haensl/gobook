// Package tree contains binary tree functionality
package tree

// Tree represents a binary tree.
type Tree struct {
	Value       int
	Left, Right *Tree
}

// Sort sorts values using a binary tree.
func Sort(values []int) {
	var root *Tree
	for _, v := range values {
		root = Add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *Tree) []int {
	if t != nil {
		values = appendValues(values, t.Left)
		values = append(values, t.Value)
		values = appendValues(values, t.Right)
	}
	return values
}

// Add a value to Tree t
func Add(t *Tree, value int) *Tree {
	if t == nil {
		t = new(Tree)
		t.Value = value
		return t
	}
	if value < t.Value {
		t.Left = Add(t.Left, value)
	} else {
		t.Right = Add(t.Right, value)
	}
	return t
}
