package nbds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListReadAndInsertion(t *testing.T) {
	// Basic creation and insertion at tail
	list := NewList[string]()
	list.AddTail("1") // index 0
	list.AddTail("2") // index 1
	list.AddTail("4") // index 2

	var value interface{}
	var err error

	assert := assert.New(t)

	// Basic reading at indexes
	if value, err = list.ReadAt(0); err != nil {
		t.Fatal("Data not found")
	}
	assert.Equal(value, "1", "Expected 1")

	if value, err = list.ReadAt(1); err != nil {
		t.Fatal("Data not found")
	}
	assert.Equal(value, "2", "Expected 2")

	if value, err = list.ReadAt(2); err != nil {
		t.Fatal("Data not found")
	}
	assert.Equal(value, "4", "Expected 4")

	// Test insertion in the middle of a list
	if err := list.InsertAt(2, "3"); err != nil {
		t.Fatal(err)
	}

	if value, err = list.ReadAt(2); err != nil {
		t.Fatal("Data not found")
	}
	assert.Equal(value, "3", "Expected 3")

	if value, err = list.ReadAt(3); err != nil {
		t.Fatal("Data not found")
	}
	assert.Equal(value, "4", "Expected 4")

	// Test insertion at index out of bounds
	if err := list.InsertAt(10, "nope"); err != ErrIndexOutOfBounds {
		t.Fatal("Insertion at index out of bounds should not work")
	}
}

func TestFindAndHasMember(t *testing.T) {
	assert := assert.New(t)

	// Basic creation and insertion at tail
	list := NewList[string]().AddTail("1").AddTail("2").AddTail("3").AddTail("4").AddTail("5")

	// Test has member
	assert.Equal(list.HasMember("5"), true, "5 should be present in the list")
	assert.Equal(list.HasMember("10"), false, "10 should not be present in the list")

	// Test find
	index, err := list.Find("5")
	if err == ErrNotFound {
		t.Fatal("5 should be present in the list")
	}

	assert.Equal(index, 4, "Index of 5 should be 4")

	// Test read at index
	val, err := list.ReadAt(4)
	if err == ErrIndexOutOfBounds {
		t.Fatal("Index 4 should exist")
	}
	assert.Equal(val, "5", "Index 4 should have the value 5")
}

func TestDeletion(t *testing.T) {
	assert := assert.New(t)

	// Basic creation from slice
	list := NewListFromSlice([]int{0, 1, 2, 3, 4, 5})

	// Pop the head
	item, err := list.PopHead()
	if err == ErrEmpty {
		t.Fatal("List should not be empty")
	}

	assert.Equal(item, 0, "Head should be 0")

	// Delete index 3
	item, err = list.DeleteAt(2)
	if err == ErrEmpty {
		t.Fatal("List should not be empty")
	}

	assert.Equal(item, 3, "Index 2 should be 3")

	// Read new index 3, which should now be 4
	item, err = list.ReadAt(2)
	if err == ErrIndexOutOfBounds {
		t.Fatal("List should contain given index of 2")
	}

	assert.Equal(item, 4, "Index 2 should be 4")

	// Delete tail
	item, err = list.PopTail()
	if err == ErrEmpty {
		t.Fatal("List should not be empty")
	}

	assert.Equal(item, 5, "Tail should be 5")
}

func TestMapAndEach(t *testing.T) {
	assert := assert.New(t)

	list := NewListFromSlice([]int{0, 1, 2, 3, 4, 5})

	cursor := 0
	list.Map(func(num int) int {
		return num * 3
	}).Each(func(num int) {
		assert.Equal(num, cursor, "Number should match cursor")
		cursor += 3
	})
}

func TestFilter(t *testing.T) {
	list := NewListFromSlice([]int{0, 1, 2, 3, 4, 5})
	list.Filter(func(num int) bool {
		return num%2 == 0
	})

	assert.Equal(t, list.ToSlice(), []int{0, 2, 4}, "Slice should only contain even numbers")
}
