// Middle of the Linked List
// https://leetcode.com/problems/middle-of-the-linked-list
package main

import "fmt"

// ListNode3 Definition for singly-linked list.
type ListNode3 struct {
	Val  int
	Next *ListNode3
}

func middleNode(head *ListNode3) *ListNode3 {
	var listSize int

	if head == nil {
		return head
	} else {
		listSize = 1
	}

	copiedHead := &ListNode3{
		Val:  head.Val,
		Next: head.Next,
	}

	for head.Next != nil {
		listSize += 1
		head = head.Next
	}

	remainder := listSize / 2

	i := 0
	for i < remainder {
		copiedHead = copiedHead.Next
		i += 1
	}
	return copiedHead
}

func main() {
	tc := []ListNode3{
		{
			Val: 1, Next: &ListNode3{
				Val: 2, Next: &ListNode3{
					Val: 3, Next: &ListNode3{
						Val: 4, Next: &ListNode3{
							Val: 5, Next: nil,
						},
					},
				},
			},
		},
		{
			Val: 1, Next: &ListNode3{
				Val: 2, Next: &ListNode3{
					Val: 3, Next: &ListNode3{
						Val: 4, Next: &ListNode3{
							Val: 5, Next: &ListNode3{
								Val: 6, Next: nil,
							},
						},
					},
				},
			},
		},
	}

	rs := [][]int{
		{3, 4, 5},
		{4, 5, 6},
	}

	for i, v := range tc {
		var result = middleNode(&v)
		j := 0
		for result.Next != nil {
			if result.Val != rs[i][j] {
				fmt.Println(result.Val, rs[i][j])
				panic("틀려")
			}
			result = result.Next
			j += 1
		}

		if result.Val != rs[i][j] {
			fmt.Println(result.Val, rs[i][j])
			panic("결국 틀려")
		}
		fmt.Println("성공")
	}
}
