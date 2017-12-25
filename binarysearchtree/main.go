package main

import "fmt"

func main() {
	bst := Tree{}
	bst.Insert(5)
	bst.Insert(1)
	bst.Insert(0)
	bst.Insert(3)
	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(8)
	bst.Insert(4)
	n, err := bst.Insert(2)
	bst.Delete(5)
	fmt.Printf("%d %s\n", n.Value, err)

	found, _ := bst.Find(4)
	fmt.Println(found)
	//f, p := bst.Root.Left.findMax(nil)
	//fmt.Printf("%d %d\n", f.Value, p.Value)
}
