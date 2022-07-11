package nbds

import "errors"

var (
	ErrIndexOutOfBounds = errors.New("index out of bounds")
	ErrNotFound         = errors.New("not found")
	ErrEmpty            = errors.New("no data in structure")
)

// List implements a double linked list using generic type D with should be a type that is comparable
type List[D comparable] struct {
	first  *node[D]
	last   *node[D]
	length int
}

// NewList returns a new list of type D, which is a type that should be comparable.
//
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

// NewListFromSlice returns a new list from slice.
// The newly created list will copy over the elements from the given slice and share its type
// Example:
//
//    // Creates a new list of type string
//    list := NewListFromSlice([]string{"a", "b", "c"})
func NewListFromSlice[D comparable](slice []D) *List[D] {
	newList := &List[D]{}

	for _, item := range slice {
		newList.AddTail(item)
	}

	return newList
}

// ReadAt returns the data stored at the given index, or returns ErrIndexOutOfBounds
// when the index doesn't exist
//
// Performance (worst case): O(n) (linear)
func (l *List[D]) ReadAt(index int) (D, error) {
	cursor := 0
	currentNode := l.first

	for cursor < index && currentNode != nil {
		cursor++
		currentNode = currentNode.right
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
		currentNode = currentNode.right
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

// AddHead adds a new element to head (beginning) of the List
//
// Performance: O(1) (constant)
func (l *List[D]) AddHead(data D) *List[D] {
	newNode := newNode[D]().setData(data).setRight(l.first)

	if l.first != nil {
		l.first.setLeft(newNode)
	}

	if l.last == nil {
		l.last = newNode
	}

	l.length++
	l.first = newNode

	return l
}

// PopHead removes the first element of the list and returns it. If the list is empty ErrEmpty is returned
//
// Performance: O(1) (constant)
func (l *List[D]) PopHead() (D, error) {
	if l.first == nil {
		var noData D
		return noData, ErrEmpty
	}

	data := l.first.data

	l.length--
	l.first = l.first.right
	l.first.setLeft(nil)

	return data, nil
}

// AddTail adds a new element to tail (end) of the List
//
// Performance: O(1) (constant)
func (l *List[D]) AddTail(data D) *List[D] {
	newNode := newNode[D]().setData(data).setLeft(l.last)

	if l.first == nil && l.last == nil {
		l.first = newNode
		l.last = newNode
		l.length++
		return l
	}

	l.length++
	l.last.right = newNode
	l.last = newNode

	return l
}

// PopTail removes the last element of the list and returns it. If the list is empty ErrEmpty is returned
//
// Performance: O(1) (constant)
func (l *List[D]) PopTail() (D, error) {
	if l.last == nil {
		var noData D
		return noData, ErrEmpty
	}

	data := l.last.data

	l.length--
	l.last = l.last.left

	return data, nil
}

// InxertAt inserts a new element to the given index. If the given index is higher
// than the current length of the list ErrIndexOutOfBounds is returned
//
// Performance (worst case): O(n) (linear)
func (l *List[D]) InsertAt(index int, data D) error {
	if index == 0 {
		l.AddHead(data)
		return nil
	}

	cursor := 0
	currentNode := l.first

	for currentNode != nil && cursor < (index-1) {
		currentNode = currentNode.right
		cursor++
	}

	if cursor < (index-1) || currentNode == nil {
		return ErrIndexOutOfBounds
	}

	currentNode.right = newNode[D]().setData(data).setRight(currentNode.right)
	l.length++

	return nil
}

// DeleteAt deletes the element from the list at the given index and returns the data stored for that element.
// If the given index is higher than the current length of the list ErrIndexOutOfBounds is returned
//
// Performance (worst case): O(n) (linear)
func (l *List[D]) DeleteAt(index int) (D, error) {
	if index == 0 {
		return l.PopHead()
	}

	cursor := 0
	currentNode := l.first

	for currentNode != nil && cursor < index {
		currentNode = currentNode.right
		cursor++
	}

	if cursor < index || currentNode == nil {
		var noData D
		return noData, ErrIndexOutOfBounds
	}

	l.length--
	returnData := currentNode.data
	currentNode.unlink()

	return returnData, nil
}

// Map applies mapFunc on every element in the list, and returns the list with the updated elements
func (l *List[D]) Map(mapFunc func(data D) D) *List[D] {
	currentNode := l.first

	for currentNode != nil {
		currentNode.updateData(mapFunc)
		currentNode = currentNode.right
	}

	return l
}

// Each invokes eachFunc on every element in the list
func (l *List[D]) Each(eachFunc func(data D)) {
	currentNode := l.first
	for currentNode != nil {
		eachFunc(currentNode.data)
		currentNode = currentNode.right
	}
}

// Filter returns a list with only the elements for which filterFunc returns true
func (l *List[D]) Filter(filterFunc func(data D) bool) *List[D] {
	currentNode := l.first
	for currentNode != nil {
		RightNode := currentNode.right
		if !filterFunc(currentNode.data) {
			currentNode.unlink()
			l.length--
		}
		currentNode = RightNode
	}

	return l
}

// ToSlice returns a slice with a copy of all elements from the list
func (l *List[D]) ToSlice() []D {
	currentNode := l.first

	slice := []D{}

	for currentNode != nil {
		slice = append(slice, currentNode.data)
		currentNode = currentNode.right
	}

	return slice
}
