package nbds

import "errors"

var (
	ErrIndexOutOfBounds = errors.New("index out of bounds")
	ErrNotFound         = errors.New("not found")
	ErrEmpty            = errors.New("no data in structure")
)

// List implements a linked list using generic type D with should be a type that is comparable
type List[D comparable] struct {
	first  *node[D]
	last   *node[D]
	length int
}

// NewList returns a new list of type D, which is a type that should be comparable.
// Example:
//
//    // Creates a new list of type string
//    list := NewList[string]
//
//    // Creates a new list of type int
//    list := NewList[int]
func NewList[D comparable]() *List[D] {
	return &List[D]{}
}

// ReadAtIndex returns the data stored at the given index, or returns ErrIndexOutOfBounds
// when the index doesn't exist
//
// Performance (worst case): O(n) (linear)
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

// Find returns the index for data of type D if present in List, or returns ErrNotFound
// if the given data could not be found in the list.
//
// Performance (worst case): O(n) (linear)
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

// HasMember returns a bool indicating whether data is in the List
//
// Performance (worst case): O(n) (linear)
func (l *List[D]) HasMember(data D) bool {
	_, err := l.Find(data)
	return err == nil
}

// Prepend adds a new element to head (beginning) of the List
//
// Performance: O(1) (constant)
func (l *List[D]) Prepend(data D) *List[D] {
	newNode := NewNode[D]().setData(data).setNext(l.first)

	if l.last == nil {
		l.last = newNode
	}

	l.length++
	l.first = newNode

	return l
}

// Append adds a new element to tail (end) of the List
//
// Performance: O(1) (constant)
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

// Adds a new element to the given index. If the given index is higher
// than the current length of the list ErrIndexOutOfBounds is returned
//
// Performance (worst case): O(n) (linear)
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
