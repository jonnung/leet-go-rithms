// Reverse Linked List
// https://leetcode.com/problems/reverse-linked-list/
package main

import "fmt"

// ListNode2 Definition for singly-linked list.
type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func reverseList(head *ListNode2) *ListNode2 {
	if head == nil {
		return nil
	}

	prev := &ListNode2{
		Val:  head.Val,
		Next: nil,
	}

	if head.Next == nil {
		return prev
	}

	current := head.Next

	var next *ListNode2
	var tempNode *ListNode2

	for current.Next != nil {
		next = current.Next
		tempNode = &ListNode2{
			Val:  current.Val,
			Next: prev,
		}

		prev = tempNode
		current = next
	}

	tempNode = &ListNode2{
		Val:  current.Val,
		Next: prev,
	}

	return tempNode
}

func main() {
	tc := []ListNode2{
		{
			Val:  1,
			Next: nil,
		},
	}

	for _, v := range tc {
		result := reverseList(&v)
		if result == nil {
			fmt.Println("nil")
		} else {
			fmt.Println(result.Val)
			fmt.Println(result.Next)
		}
	}
}
