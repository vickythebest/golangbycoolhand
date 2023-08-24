package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// fmt.Println("Remove duplicate")
	list := &ListNode{}
	list.Val = 1

	list1 := &ListNode{}
	list1.Val = 1
	list.Next = list1

	list2 := &ListNode{}
	list2.Val = 1
	list1.Next = list2

	// list3 := &ListNode{}
	// list3.Val = 3
	// list2.Next = list3

	// list4 := &ListNode{}
	// list4.Val = 3
	// list3.Next = list4

	result := deleteDuplicates(list)

	for result != nil {
		fmt.Println("result : ", result.Val)
		result = result.Next
	}

}

func deleteDuplicates(head *ListNode) *ListNode {

	currentNode := head

	for currentNode != nil {

		if currentNode.Next != nil && currentNode.Val == currentNode.Next.Val {
			temp := currentNode.Next
			currentNode.Next = temp.Next
		} else {
			currentNode = currentNode.Next
		}

	}

	return head
}
func deleteDuplicates1(head *ListNode) *ListNode {

	currentNode := head

	if currentNode != nil && currentNode.Next != nil {

		left := currentNode.Val
		right := currentNode.Next.Val

		for currentNode.Next != nil {

			if left == right {
				currentNode.Next = currentNode.Next.Next
				if currentNode.Next != nil {
					right = currentNode.Next.Val
				}

			} else {
				currentNode = currentNode.Next
				left = currentNode.Val
			}

		}
		return head
	} else {
		return head
	}
}
