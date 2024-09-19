package dsa

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (l *LinkedList) InsertAtHead(data int) {
	temp1 := Node{Data: data, Next: nil}

	if l.Head == nil {
		l.Head = &temp1
	} else {
		temp2 := l.Head
		l.Head = &temp1
		temp1.Next = temp2
	}
	l.Length += 1
}

func (l *LinkedList) InsertAtTail(data int) {
	temp1 := Node{Data: data, Next: nil}

	if l.Head == nil {
		l.Head = &temp1
	} else {
		temp2 := l.Head
		for temp2.Next != nil {
			temp2 = temp2.Next
		}
		temp2.Next = &temp1
	}
	l.Length += 1
}

func (l *LinkedList) Insert(n, data int) {
	if n == 0 {
		l.InsertAtHead(data)
	} else if n == l.Length-1 {
		l.InsertAtTail(data)
	} else {
		temp1 := Node{Data: data, Next: nil}
		temp2 := l.Head

		for i := 0; i < n-1; i++ {
			temp2 = temp2.Next
		}

		temp1.Next = temp2.Next
		temp2.Next = &temp1
		l.Length += 1
	}
}

func (l *LinkedList) DeleteAtHead() {
	temp := l.Head
	l.Head = temp.Next

	l.Length -= 1
}

func (l *LinkedList) DeleteAtTail() {
	temp1 := l.Head
	var temp2 *Node
	for temp1.Next != nil {
		temp2 = temp1
		temp1 = temp1.Next
	}
	temp2.Next = nil

	l.Length -= 1
}

func (l *LinkedList) Delete(n int) {
	if n == 0 {
		l.DeleteAtHead()
	} else if n == l.Length-1 {
		l.DeleteAtTail()
	} else {
		temp1 := l.Head
		for i := 0; i < n-1; i++ {
			temp1 = temp1.Next
		}
		temp2 := temp1.Next
		temp1.Next = temp2.Next
		l.Length -= 1
	}
}

func (l *LinkedList) Reverse() {
	var curr, prev, next *Node
	curr = l.Head
	prev = nil

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head = prev
}
