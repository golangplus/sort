package sortp

import(
	"sort"
	"testing"
)

// Make sure InterfaceStruct implements sort.Interface
var _ sort.Interface = InterfaceStruct{}

func TestSortF(t *testing.T) {
	data := []int{5, 3, 1, 8, 0}
	
	SortF(len(data), func(i, j int) bool {
		return data[i] < data[j]
	}, func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})
	
	if !sort.IntsAreSorted(data) {
		t.Errorf("Data is not sorted by SortF: %v", data)
	}
}