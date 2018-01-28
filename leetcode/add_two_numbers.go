package main

import "fmt"

type ListNode struct {
	value int
	next  *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	sumNode := &ListNode{value: 0}
	head := sumNode

	for l1 != nil || l2 != nil {
		x := l1.value
		y := l2.value
		sum := x + y + carry
		carry = sum / 10
		fmt.Printf("Sum: %d, Carry: %d\n", sum, carry)
		sumNode.next = &ListNode{value: (sum % 10)}
		sumNode = sumNode.next

		if l1 != nil {
			l1 = l1.next
		}
		if l2 != nil {
			l2 = l2.next
		}
	}

	if carry != 0 {
		sumNode.next = &ListNode{value: carry}
	}

	return head.next
}

func main() {
	l1 := &ListNode{value: 2, next: &ListNode{value: 4, next: &ListNode{value: 3}}}
	l2 := &ListNode{value: 5, next: &ListNode{value: 6, next: &ListNode{value: 4}}}
	l3 := addTwoNumbers(l1, l2)
	for l3 != nil {
		fmt.Print(l3.value)
		l3 = l3.next
	}
	fmt.Println("")
}
