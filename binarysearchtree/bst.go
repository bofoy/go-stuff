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
	left  *node
	right *node
	value int
}

// Go doesn't have classes but you can define methods on types with "receivers"
// (t *Tree) is a "receiver". Similar to "this" in Java
// Use pointer receiver when method should modify the value the receiver points
// to. Also used to avoid copying value on each method call since everything in
// Go is pass by value
func (t *Tree) Insert(value int) error {
	if t == nil {
		t = &Tree{&node{value: value}}
	} else if t.Root == nil {
		t.Root = &node{value: value}
	} else {
		_, err := t.Root.insert(value)
		return err
	}
	return nil
}

func (t *Tree) Delete(value int) error {
	if t == nil || t.Root == nil {
		return errors.New("Tree is empty, cannot delete from empty tree")
	}
	return t.Root.delete(value, t.Root)
}

func (t *Tree) Find(value int) (bool, error) {
	if t == nil || t.Root == nil {
		return false, errors.New("Tree is empty, cannot search empty tree")
	}
	return t.Root.find(value), nil
}

func (n *node) insert(value int) (*node, error) {
	if n == nil {
		return &node{value: value}, nil
	} else if value < n.value { // if node is nil insert new node as left child
		node, _ := n.left.insert(value)
		n.left = node
	} else if value > n.value { // if node is nil insert new node as right child
		node, _ := n.right.insert(value)
		n.right = node
	}
	return n, fmt.Errorf("Value %d already exists", value)
}

// finds max node of a subtree and its parent node
func (n *node) findMax(parent *node) (node *node, p *node) {
	if n.right == nil {
		return n, parent
	}
	return n.right.findMax(n)
}

// replaces parent's child pointer from n to replacement
func (n *node) replaceNode(parent *node, replacement *node) {
	if n == parent.left {
		parent.left = replacement
	} else {
		parent.right = replacement
	}
}

func (n *node) delete(value int, parent *node) error {
	if n == nil {
		return errors.New("Value to be deleted does not exist in tree")
	} else if value < n.value {
		n.left.delete(value, n)
	} else if value > n.value {
		n.right.delete(value, n)
	} else { // found node to be deleted
		if n.left == nil && n.right == nil { // leaf node
			n.replaceNode(parent, nil) // if no children, then simply set to nil
		}
		// if one child, point parent to either the left or right child
		if n.left == nil {
			n.replaceNode(parent, n.right)
			return nil
		}
		if n.right == nil {
			n.replaceNode(parent, n.left)
			return nil
		}

		// if node has 2 children, recursively find the max of the
		// left subtree and its parent node
		replacement, replParent := n.left.findMax(n)
		n.value = replacement.value

		return replacement.delete(replacement.value, replParent)
	}
	return nil
}

func (n *node) find(value int) bool {
	if value == n.value {
		return true
	} else if value < n.value {
		if n.left != nil {
			return n.left.find(value)
		}
	} else if value > n.value {
		if n.right != nil {
			return n.right.find(value)
		}
	}
	return false
}
