package main

import (
	"errors"
	"fmt"
)

type Tree struct {
	Root *node
}

// lowercase things are not exported
type node struct {
	Left  *node
	Right *node
	Value int
}

// Go doesn't have classes but you can define methods on types with "receivers"
// (t *Tree) is a "receiver". Similar to "this" in Java
// Use pointer receiver when method should modify the Value the receiver points
// to. Also used to avoid copying Value on each method call since everything in
// Go is pass by Value
func (t *Tree) Insert(Value int) *node {
	if t == nil {
		t = &Tree{&node{Value: Value}}
	} else if t.Root == nil {
		t.Root = &node{Value: Value}
	} else {
		return t.Root.insert(Value)
	}
	return t.Root
}

func (t *Tree) Delete(Value int) (*node, error) {
	if t == nil {
		return t.Root, errors.New("Tree is empty, cannot delete node")
	}
	return t.Root.delete(Value)
}

func (t *Tree) Find(Value int) (*node, error) {
	if t == nil {
		return t.Root, errors.New("Tree is empty, cannot find node")
	}
	return t.Root.find(Value)
}

func (n *node) insert(Value int) *node {
	if Value < n.Value {
		if n.Left == nil { // if node is nil insert new node as Left child
			n.Left = &node{Value: Value}
		}
		// recursively traverse Left node and return node inserted
		return n.Left.insert(Value)
	} else if Value > n.Value {
		if n.Right == nil { // if node is nil insert new node as Right child
			n.Right = &node{Value: Value}
		}
		// recursively traverse Right node and return node inserted
		return n.Right.insert(Value)
	}
	return n
}

// finds max value of a subtree
func (n *node) findMax(parent *node) (node *node, p *node) {
	if n.Right == nil {
		return n, parent
	}
	return n.Right.findMax(n)
}

func (n *node) delete(Value int) (node *node, err error) {
	if Value == n.Value {
		n = nil
	} else if Value < n.Value {
		if n.Left == nil {
			return nil, errors.New("Node not found, cannot be deleted")
		}
		n.Left.delete(Value)
	} else if Value > n.Value {
		if n.Right == nil {
			return nil, errors.New("Node not found, cannot be deleted")
		}
		n.Right.delete(Value)
	}
	return n, nil
}

func (n *node) find(Value int) (*node, error) {
	if Value == n.Value {
		fmt.Printf("%d exists\n", Value)
	} else if Value < n.Value {
		if n.Left == nil {
			return nil, errors.New("Node not found")
		}
		// recursively traverse Left node and return the node found
		return n.Left.find(Value)
	} else if Value > n.Value {
		if n.Right == nil {
			return nil, errors.New("Node not found")
		}
		// recursively traverse Right node and return the node found
		return n.Right.find(Value)
	}
	return n, nil
}
