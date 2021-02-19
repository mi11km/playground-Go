package chapter2

type SinglyLinkedList struct {
	head   *node
	length int
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{nil, 0}
}

func (l *SinglyLinkedList) Insert(val int) {
	newNode := &node{value: val}
	n := l.head
	if n == nil {
		l.head = newNode
	} else {
		for n.next != nil {
			n = n.next
		}
		n.next = newNode
	}
	l.length++
}

func (l *SinglyLinkedList) Delete(index int) {
	if l.length <= index || index < 0 {
		return
	}
	if index == 0 {
		l.head = l.head.next
	} else {
		pre := l.head
		for i := 0; i < index-1; i++ {
			pre = pre.next
		}
		if pre.next.next == nil {
			pre.next = nil
		} else {
			pre.next = pre.next.next
		}
	}
	l.length--
}

func (l *SinglyLinkedList) getNode(index int) *node {
	if l.length <= index {
		return nil
	}
	n := l.head
	for i := 0; i < index; i++ {
		n = n.next
	}
	return n
}

func (l *SinglyLinkedList) Get(index int) int {
	if n := l.getNode(index); n != nil {
		return n.value
	}
	return -1
}

func (l *SinglyLinkedList) Len() int {
	return l.length
}

/*--------------*/

// Mostly used for testing and debug
func (l *SinglyLinkedList) Slice() []int {
	slice := make([]int, l.length)
	n := l.head
	for i := 0; i < l.length; i++ {
		slice[i] = n.value
		n = n.next
	}
	return slice
}
