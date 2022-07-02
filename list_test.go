package nbds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListReadAndInsertion(t *testing.T) {

	// Basic creation and insertion at tail
	list := NewList[string]()
	list.Append("1") // index 0
	list.Append("2") // index 1
	list.Append("4") // index 2

	var value interface{}
	var err error

	assert := assert.New(t)

	// Basic reading at indexes
	if value, err = list.ReadAtIndex(0); err != nil {
		t.Fatal("Data not found")
	}
	assert.Equal(value, "1", "Expected 1")

	if value, err = list.ReadAtIndex(1); err != nil {
		t.Fatal("Data not found")
	}
	assert.Equal(value, "2", "Expected 2")

	if value, err = list.ReadAtIndex(2); err != nil {
		t.Fatal("Data not found")
	}
	assert.Equal(value, "4", "Expected 4")

	// Test insertion in the middle of a list
	if err := list.InsertAtIndex(2, "3"); err != nil {
		t.Fatal(err)
	}

	if value, err = list.ReadAtIndex(2); err != nil {
		t.Fatal("Data not found")
	}
	assert.Equal(value, "3", "Expected 3")

	if value, err = list.ReadAtIndex(3); err != nil {
		t.Fatal("Data not found")
	}
	assert.Equal(value, "4", "Expected 4")

	// Test insertion at index out of bounds
	if err := list.InsertAtIndex(10, "nope"); err != ErrIndexOutOfBounds {
		t.Fatal("Insertion at index out of bounds should not work")
	}
}

func TestFindAndHasMember(t *testing.T) {
	assert := assert.New(t)

	// Basic creation and insertion at tail
	list := NewList[string]().Append("1").Append("2").Append("3").Append("4").Append("5")

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
	val, err := list.ReadAtIndex(4)
	if err == ErrIndexOutOfBounds {
		t.Fatal("Index 4 should exist")
	}
	assert.Equal(val, "5", "Index 4 should have the value 5")
}
