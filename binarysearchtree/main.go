package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	bst := Tree{}
	bst.Insert(5)
	bst.Insert(1)
	bst.Insert(0)
	bst.Insert(3)
	n := bst.Insert(2)
	bst.Insert(7)
	bst.Insert(8)
	bst.Insert(4)
	fmt.Printf("Inserted: %d\n", n.Value)
	//	bst.Insert(49)
	//	bst.Insert(1)
	//	bst.Insert(199)
	//	bst.Insert(5)
	bst.Find(59)
	f, p := bst.Root.Left.findMax(nil)
	fmt.Printf("%d %d\n", f.Value, p.Value)
}
