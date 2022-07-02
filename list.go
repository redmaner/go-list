package nbds

import "errors"

var (
	ErrIndexOutOfBounds = errors.New("index out of bounds")
	ErrNotFound         = errors.New("not found")
)

type List[D comparable] struct {
	first  *node[D]
	last   *node[D]
	length int
}

func NewList[D comparable]() *List[D] {
	return &List[D]{}
}

func (l *List[D]) ReadAtIndex(index int) (D, error) {
	cursor := 0
	currentNode := l.first

	for cursor < index && currentNode != nil {
		cursor++
		currentNode = currentNode.next
	}

	if currentNode == nil {
		var noData D
		return noData, ErrIndexOutOfBounds
	}

	return currentNode.data, nil
}

func (l *List[D]) Find(data D) (int, error) {
	cursor := 0
	currentNode := l.first

	for currentNode != nil {
		if currentNode.data == data {
			return cursor, nil
		}

		cursor++
		currentNode = currentNode.next
	}

	return 0, ErrNotFound
}

func (l *List[D]) HasMember(data D) bool {
	_, err := l.Find(data)
	return err == nil
}

func (l *List[D]) Prepend(data D) *List[D] {
	newNode := NewNode[D]().setData(data).setNext(l.first)

	if l.last == nil {
		l.last = newNode
	}

	l.length++
	l.first = newNode

	return l
}

func (l *List[D]) Append(data D) *List[D] {
	newNode := NewNode[D]().setData(data)

	if l.first == nil && l.last == nil {
		l.first = newNode
		l.last = newNode
		return l
	}

	l.length++
	l.last.next = newNode
	l.last = newNode

	return l
}

func (l *List[D]) InsertAtIndex(index int, data D) error {
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

	currentNode.next = NewNode[D]().setData(data).setNext(currentNode.next)

	return nil
}
