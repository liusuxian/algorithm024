package Week_01

import "testing"

// 合并2个有序链表

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归法
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val <= l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}

	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

func TestMergeTwoLists(t *testing.T) {
	nodes1 := [3]*ListNode{
		{Val: 1, Next: nil},
		{Val: 2, Next: nil},
		{Val: 4, Next: nil},
	}
	nodes2 := [3]*ListNode{
		{Val: 1, Next: nil},
		{Val: 3, Next: nil},
		{Val: 4, Next: nil},
	}

	for i := 0; i < 2; i++ {
		nodes1[i].Next = nodes1[i+1]
	}
	for i := 0; i < 2; i++ {
		nodes2[i].Next = nodes2[i+1]
	}

	for i := nodes1[0]; i != nil; i = i.Next {
		t.Log(i.Val)
	}
	for i := nodes2[0]; i != nil; i = i.Next {
		t.Log(i.Val)
	}

	h1 := mergeTwoLists(nil, nil)
	for i := h1; i != nil; i = i.Next {
		t.Log(i.Val)
	}

	h2 := mergeTwoLists(nodes1[0], nil)
	for i := h2; i != nil; i = i.Next {
		t.Log(i.Val)
	}

	h3 := mergeTwoLists(nil, nodes2[0])
	for i := h3; i != nil; i = i.Next {
		t.Log(i.Val)
	}

	h4 := mergeTwoLists(nodes1[0], nodes2[0])
	for i := h4; i != nil; i = i.Next {
		t.Log(i.Val)
	}
}
