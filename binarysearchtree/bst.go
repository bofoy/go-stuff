package main

import (
	"errors"
	"fmt"
)

type Tree struct {
	Root *Node
}

// lowercase things are not exported
type Node struct {
	Left  *Node
	Right *Node
	Value int
}

// Go doesn't have classes but you can define methods on types with "receivers"
// (t *Tree) is a "receiver". Similar to "this" in Java
// Use pointer receiver when method should modify the Value the receiver points
// to. Also used to avoid copying Value on each method call since everything in
// Go is pass by Value
func (t *Tree) Insert(value int) (insert *Node, err error) {
	if t == nil {
		t = &Tree{&Node{Value: value}}
	} else if t.Root == nil {
		t.Root = &Node{Value: value}
	} else {
		return t.Root.insert(value)
	}
	return t.Root, nil
}

func (t *Tree) Delete(value int) error {
	if t == nil {
		return errors.New("Tree is empty, cannot delete Node")
	}
	return t.Root.delete(value, t.Root)
}

func (t *Tree) Find(value int) (bool, error) {
	if t == nil {
		return false, errors.New("Tree is empty, cannot find Node")
	} else if t.Root.find(value) == true {
		return true, nil
	} else {
		return false, nil
	}
}

func (n *Node) insert(value int) (insert *Node, err error) {
	if value < n.Value {
		if n.Left == nil { // if Node is nil insert new Node as Left child
			n.Left = &Node{Value: value}
		}
		// recursively traverse Left Node and return Node inserted
		return n.Left.insert(value)
	} else if value > n.Value {
		if n.Right == nil { // if Node is nil insert new Node as Right child
			n.Right = &Node{Value: value}
		}
		// recursively traverse Right Node and return Node inserted
		return n.Right.insert(value)
	}
	return n, fmt.Errorf("Value %d already exists", value)
}

// finds max Node of a subtree and its parent Node
func (n *Node) findMax(parent *Node) (Node *Node, p *Node) {
	if n.Right == nil {
		return n, parent
	}
	return n.Right.findMax(n)
}

// replaces parent's child pointer from n to replacement
func (n *Node) replaceNode(parent *Node, replacement *Node) {
	if n == parent.Left {
		parent.Left = replacement
	} else {
		parent.Right = replacement
	}
}

func (n *Node) delete(value int, parent *Node) error {
	if n == nil {
		return errors.New("Value to be deleted does not exist in tree")
	} else if value < n.Value {
		n.Left.delete(value, n)
	} else if value > n.Value {
		n.Right.delete(value, n)
	} else { // found Node to be deleted
		if n.Left == nil && n.Right == nil { // leaf Node
			n.replaceNode(parent, nil) // if no children, then simply set to nil
		}
		// if one child, point parent to either the left or right child
		if n.Left == nil {
			n.replaceNode(parent, n.Right)
			return nil
		}
		if n.Right == nil {
			n.replaceNode(parent, n.Left)
			return nil
		}

		// if Node has 2 children, recursively find the max of the
		// left subtree and its parent Node
		replacement, replParent := n.Left.findMax(n)
		n.Value = replacement.Value

		return replacement.delete(replacement.Value, replParent)
	}
	return nil
}

func (n *Node) find(value int) bool {
	if value == n.Value {
		return true
	} else if value < n.Value {
		if n.Left == nil {
			return false
		}
		// recursively traverse Left Node and return the Node found
		return n.Left.find(value)
	}
	if n.Right == nil {
		return false
	}
	// recursively traverse Right Node and return the Node found
	return n.Right.find(value)
}
