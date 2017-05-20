
package main

import "fmt"


type Node interface {
	SetValue(v int)
	GetValue() int
}

//type SLLNOde
type SLLNode struct {
	next *SLLNode
	value int
}

type SingeLinkedList struct{
	head *SLLNode
	tail *SLLNode
}

func (sNode *SLLNode) SetValue(v int) {
	sNode.value = v
}

func (sNode *SLLNode) GetValue() int {
	return sNode.value
}

func NewSLLNode() *SLLNode {
	return new(SLLNode)
}

func newSingleLinkedList() *SingeLinkedList{
	return new(SingeLinkedList)
}

func (list *SingeLinkedList) Add(v int) {
	newNode := &SLLNode{value: v}
	if list.head == nil {
		list.head = newNode
	} else if list.tail == list.head {
		list.head.next = newNode
	} else if list.tail != nil {
		list.tail.next = newNode
	}
	list.tail = newNode
}
	
func (list  *SingeLinkedList) String() string {
	s := ""
	for n := list.head; n != nil; n = n.next {
		s += fmt.Sprintf(" {%d} ", n.GetValue())
	}
	return s
}

func main() {
	list := newSingleLinkedList()
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	fmt.Println("Hello, playground", list)
}
