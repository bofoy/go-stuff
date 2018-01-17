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
	err := bst.Insert(2)
	bst.Delete(5)
	fmt.Printf("%s\n", err)
	fmt.Printf("Value is: %d\n", bst.Root.right.right.value)
	found, _ := bst.Find(2)
	fmt.Println(found)
	//f, p := bst.Root.Left.findMax(nil)
	//fmt.Printf("%d %d\n", f.Value, p.Value)
}
