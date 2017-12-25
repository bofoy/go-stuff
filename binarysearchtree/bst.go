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
		return t.Root.insert(value)
	}
	return nil
}

func (t *Tree) Delete(value int) error {
	if t == nil {
		return errors.New("Tree is empty, cannot delete node")
	}
	return t.Root.delete(value, t.Root)
}

func (t *Tree) Find(value int) (bool, error) {
	if t == nil {
		return false, errors.New("Tree is empty, cannot find node")
	} else if t.Root.find(value) == true {
		return true, nil
	} else {
		return false, nil
	}
}

func (n *node) insert(value int) error {
	if value < n.value {
		if n.left == nil { // if node is nil insert new node as left child
			n.left = &node{value: value}
		}
		// recursively traverse left node and return node inserted
		n.left.insert(value)
	} else if value > n.value {
		if n.right == nil { // if node is nil insert new node as right child
			n.right = &node{value: value}
		}
		// recursively traverse right node and return node inserted
		n.right.insert(value)
	}
	return fmt.Errorf("Value %d already exists", value)
}

// finds max node of a subtree and its parent node
func (n *node) findMax(parent *node) (node *node, p *node) {
	if n.right == nil {
		return n, parent
	}
	return n.right.findMax(n)
}

// replaces parent's child pointer from n to replacement
func (n *node) replacenode(parent *node, replacement *node) {
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
			n.replacenode(parent, nil) // if no children, then simply set to nil
		}
		// if one child, point parent to either the left or right child
		if n.left == nil {
			n.replacenode(parent, n.right)
			return nil
		}
		if n.right == nil {
			n.replacenode(parent, n.left)
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
		if n.left == nil {
			return false
		}
		// recursively traverse left node and return the node found
		return n.left.find(value)
	}
	if n.right == nil {
		return false
	}
	// recursively traverse right node and return the node found
	return n.right.find(value)
}
