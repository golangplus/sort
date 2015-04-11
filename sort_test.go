package sortp

import (
	"fmt"
	"sort"
)

// Make sure InterfaceStruct implements sort.Interface
var _ sort.Interface = InterfaceStruct{}

func ExampleSortF() {
	data := []int{5, 3, 1, 8, 0}

	SortF(len(data), func(i, j int) bool {
		return data[i] < data[j]
	}, func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})
	
	fmt.Println(data)
	// OUTPUT:
	// [0 1 3 5 8]
}

func ExampleInterfaceStruct() {
	data := []int{5, 3, 1, 8, 0}

	sort.Sort(InterfaceStruct{
		LenF: func() int {
			return len(data)
		}, LessF: func(i, j int) bool {
			return data[i] < data[j]
		}, SwapF: func(i, j int) {
			data[i], data[j] = data[j], data[i]
		},
	})
	
	fmt.Println(data)
	// OUTPUT:
	// [0 1 3 5 8]
}
