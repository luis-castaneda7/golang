package main

type ListNode struct {
	Val  int
	Next *ListNode
}

var collec = make(map[*ListNode]int)

func hasCycle(head *ListNode) bool {
	if head.Next == nil {
		return false
	}
	if _, ok := collec[head]; ok {
		return true
	}

	collec[head] = 0
	return hasCycle(head.Next)

}

func main() {

}
