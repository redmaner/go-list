package list

type node struct {
	data string
	next *node
}

type List struct {
	first  *node
	last   *node
	length int
}

func New() *List {
	return &List{}
}

func (l *List) ReadAtIndex(index int) (string, bool) {
	cursor := 0
	currentNode := l.first

	for cursor < index {
		if currentNode == nil {
			return "", false
		}

		cursor++
		currentNode = currentNode.next
	}

	if currentNode == nil {
		return "", false
	}

	return currentNode.data, true
}

func (l *List) Find(data string) (index int, found bool) {
	cursor := 0
	currentNode := l.first

	for currentNode != nil {
		if currentNode.data == data {
			return cursor, true
		}
	}

	return 0, false
}

func (l *List) InsertAtHead(data string) {
	l.length++

	newNode := &node{
		data: data,
	}

	if l.first == nil && l.last == nil {
		l.first = newNode
		l.last = newNode
		return
	}

	newNode.next = l.first
	l.first = newNode
	return
}

func (l *List) InsertAtTail(data string) {
	l.length++

	newNode := &node{
		data: data,
	}

	if l.first == nil && l.last == nil {
		l.first = newNode
		l.last = newNode
		return
	}

	l.last.next = newNode
	l.last = newNode
	return
}
