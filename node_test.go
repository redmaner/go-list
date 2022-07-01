package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListReadAndInsertion(t *testing.T) {

	// Basic creation and insertion at tail
	list := New()
	list.InsertAtTail("1") // index 0
	list.InsertAtTail("2") // index 1
	list.InsertAtTail("4") // index 2

	var value string
	var ok bool

	// Basic reading at indexes
	if value, ok = list.ReadAtIndex(0); !ok {
		t.Fatal("Data not found")
	}
	assert.Equal(t, value, "1")

	if value, ok = list.ReadAtIndex(1); !ok {
		t.Fatal("Data not found")
	}
	assert.Equal(t, value, "2")

	if value, ok = list.ReadAtIndex(2); !ok {
		t.Fatal("Data not found")
	}
	assert.Equal(t, value, "4")

	// Test insertion in the middle of a list
	if err := list.InsertAtIndex(2, "3"); err != nil {
		t.Fatal(err)
	}

	if value, ok = list.ReadAtIndex(2); !ok {
		t.Fatal("Data not found")
	}
	assert.Equal(t, value, "3")

	if value, ok = list.ReadAtIndex(3); !ok {
		t.Fatal("Data not found")
	}
	assert.Equal(t, value, "4")

	// Test insertion at index out of bounds
	if err := list.InsertAtIndex(10, "nope"); err != ErrIndexOutOfBounds {
		t.Fatal("Insertion at index out of bounds should not work")
	}
}
