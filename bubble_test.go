package sortp

import (
	"fmt"
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
