package nbds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListReadAndInsertion(t *testing.T) {

	// Basic creation and insertion at tail
	list := New[string]()
	list.Append("1") // index 0
	list.Append("2") // index 1
	list.Append("4") // index 2

	var value interface{}
	var ok bool

	assert := assert.New(t)

	// Basic reading at indexes
	if value, ok = list.ReadAtIndex(0); !ok {
		t.Fatal("Data not found")
	}
	assert.Equal(value, "1", "Expected 1")

	if value, ok = list.ReadAtIndex(1); !ok {
		t.Fatal("Data not found")
	}
	assert.Equal(value, "2", "Expected 2")

	if value, ok = list.ReadAtIndex(2); !ok {
		t.Fatal("Data not found")
	}
	assert.Equal(value, "4", "Expected 4")

	// Test insertion in the middle of a list
	if err := list.InsertAtIndex(2, "3"); err != nil {
		t.Fatal(err)
	}

	if value, ok = list.ReadAtIndex(2); !ok {
		t.Fatal("Data not found")
	}
	assert.Equal(value, "3", "Expected 3")

	if value, ok = list.ReadAtIndex(3); !ok {
		t.Fatal("Data not found")
	}
	assert.Equal(value, "4", "Expected 4")

	// Test insertion at index out of bounds
	if err := list.InsertAtIndex(10, "nope"); err != ErrIndexOutOfBounds {
		t.Fatal("Insertion at index out of bounds should not work")
	}
}
