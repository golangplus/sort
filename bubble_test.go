package sortp

import (
	"fmt"
	"testing"

	"github.com/golangplus/testing/assert"
)

func ExampleBubbleF() {
	data := []int{5, 3, 1, 8, 0}

	BubbleF(len(data), func(i, j int) bool {
		return data[i] < data[j]
	}, func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})

	fmt.Println(data)
	// OUTPUT:
	// [0 1 3 5 8]
}

func TestBubble(t *testing.T) {
	data := []int{5, 3, 1, 8, 0}

	Bubble(InterfaceStruct {
		LenF: func() int {
			return len(data)
		},
		LessF: func(i, j int) bool {
			return data[i] < data[j]
		},
		SwapF: func(i, j int) {
			data[i], data[j] = data[j], data[i]
		},
	})
	assert.Equal(t, "data", data, []int{0, 1, 3, 5, 8})
}
