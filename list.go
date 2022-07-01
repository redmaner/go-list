package list

import "errors"

var (
	ErrIndexOutOfBounds = errors.New("index out of bounds")
)

type List struct {
	first  *node
	last   *node
	length int
}

func New() *List {
	return &List{}
}

func (l *List) ReadAtIndex(index int) (interface{}, bool) {
	cursor := 0
	currentNode := l.first

	for cursor < index && currentNode != nil {
		cursor++
		currentNode = currentNode.next
	}

	if currentNode == nil {
		return "", false
	}

	return currentNode.data, true
}

func (l *List) Find(data interface{}) (index int, found bool) {
	cursor := 0
	currentNode := l.first

	for currentNode != nil {
		if currentNode.data == data {
			return cursor, true
		}
	}

	return 0, false
}

func (l *List) Prepend(data interface{}) {
	newNode := NewNode().setData(data).setNext(l.first)

	if l.last == nil {
		l.last = newNode
	}

	l.length++
	l.first = newNode
}

func (l *List) Append(data interface{}) {
	newNode := NewNode().setData(data)

	if l.first == nil && l.last == nil {
		l.first = newNode
		l.last = newNode
		return
	}

	l.length++
	l.last.next = newNode
	l.last = newNode
}

func (l *List) InsertAtIndex(index int, data interface{}) error {
	if index == 0 {
		l.Prepend(data)
		return nil
	}

	cursor := 0
	currentNode := l.first

	for currentNode != nil && cursor < (index-1) {
		currentNode = currentNode.next
		cursor++
	}

	if cursor < (index-1) || currentNode == nil {
		return ErrIndexOutOfBounds
	}

	currentNode.next = NewNode().setData(data).setNext(currentNode.next)

	return nil
}
