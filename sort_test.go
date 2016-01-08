// Copyright 2015 The Golang Plus Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sortp

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/golangplus/testing/assert"
)

// Make sure InterfaceStruct implements sort.Interface
var _ sort.Interface = InterfaceStruct{}

func TestSortF(t *testing.T) {
	data := []int{5, 3, 1, 8, 0}

	less := func(i, j int) bool {
		return data[i] < data[j]
	}
	swap := func(i, j int) {
		data[i], data[j] = data[j], data[i]
	}
	SortF(len(data), less, swap)

	assert.True(t, "IsSortedF", IsSortedF(len(data), less))
	assert.Equal(t, "data", data, []int{0, 1, 3, 5, 8})
}

func TestIndexSort(t *testing.T) {
	data := []int{5, 3, 1, 8, 0}
	indexes := IndexSort(sort.IntSlice(data))
	assert.Equal(t, "indexes", indexes, []int{4, 2, 1, 0, 3})
}

func TestIndexSortF(t *testing.T) {
	data := []int{5, 3, 1, 8, 0}
	indexes := IndexSortF(len(data), func(i, j int) bool {
		return data[i] < data[j]
	})
	assert.Equal(t, "indexes", indexes, []int{4, 2, 1, 0, 3})
}

func TestStableF(t *testing.T) {
	data := []int{5, 1, 1, 8, 1}
	sub := []int{1, 2, 3, 4, 5}
	less := func(i, j int) bool {
		return data[i] < data[j]
	}
	swap := func(i, j int) {
		data[i], data[j] = data[j], data[i]
		sub[i], sub[j] = sub[j], sub[i]
	}
	StableF(len(data), less, swap)
	assert.True(t, "IsSortedF", IsSortedF(len(data), less))
	assert.Equal(t, "data", data, []int{1, 1, 1, 5, 8})
	assert.Equal(t, "sub", sub, []int{2, 3, 5, 1, 4})
}

func TestIndexStable(t *testing.T) {
	data := []int{5, 1, 1, 8, 1}
	indexes := IndexStable(sort.IntSlice(data))
	assert.Equal(t, "indexes", indexes, []int{1, 2, 4, 0, 3})
}

func TestIndexStableF(t *testing.T) {
	data := []int{5, 1, 1, 8, 1}
	indexes := IndexStableF(len(data), func(i, j int) bool {
		return data[i] < data[j]
	})
	assert.Equal(t, "indexes", indexes, []int{1, 2, 4, 0, 3})
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

func TestIsSortedF(t *testing.T) {
	list := []int{1, 2, 3}
	assert.True(t, "IsSortedF", IsSortedF(len(list), func(i, j int) bool {
		return list[i] < list[j]
	}))
	list = []int{2, 1, 3}
	assert.False(t, "IsSortedF", IsSortedF(len(list), func(i, j int) bool {
		return list[i] < list[j]
	}))
}

func TestReverseLess(t *testing.T) {
	arr := []int{1, 1, 2}
	less := func(i, j int) bool {
		return arr[i] < arr[j]
	}
	assert.False(t, "less", less(0, 0))
	assert.False(t, "less", less(0, 1))
	assert.True(t, "less", less(0, 2))
	assert.False(t, "less", less(1, 0))
	assert.False(t, "less", less(1, 1))
	assert.True(t, "less", less(1, 2))
	assert.False(t, "less", less(2, 0))
	assert.False(t, "less", less(2, 1))
	assert.False(t, "less", less(2, 2))

	less = ReverseLess(less)
	assert.False(t, "less", less(0, 0))
	assert.False(t, "less", less(0, 1))
	assert.False(t, "less", less(0, 2))
	assert.False(t, "less", less(1, 0))
	assert.False(t, "less", less(1, 1))
	assert.False(t, "less", less(1, 2))
	assert.True(t, "less", less(2, 0))
	assert.True(t, "less", less(2, 1))
	assert.False(t, "less", less(2, 2))
}

func TestMerge(t *testing.T) {
	left, right := make([]int, 100), make([]int, 200)
	for i := range left {
		left[i] = rand.Int()
	}
	for i := range right {
		right[i] = rand.Int()
	}

	// Expected results can be obtained by sorting.
	expMerged := append(append([]int{}, left...), right...)
	sort.Ints(expMerged)

	// Sort left and right
	sort.Ints(left)
	sort.Ints(right)

	// Do merge
	merged := make([]int, 0, len(left)+len(right))
	Merge(len(left), len(right), func(l, r int) bool {
		return left[l] < right[r]
	}, func(l int) {
		merged = append(merged, left[l])
	}, func(r int) {
		merged = append(merged, right[r])
	})

	assert.StringEqual(t, "merged", merged, expMerged)
}

func ExampleMerge() {
	left := []int{1, 3, 5, 7}
	right := []int{4, 6, 8, 10, 11}

	var merged []int
	Merge(len(left), len(right), func(l, r int) bool {
		return left[l] < right[r]
	}, func(l int) {
		merged = append(merged, left[l])
	}, func(r int) {
		merged = append(merged, right[r])
	})
	fmt.Println(merged)

	left, right = right, left
	merged = nil
	Merge(len(left), len(right), func(l, r int) bool {
		return left[l] < right[r]
	}, func(l int) {
		merged = append(merged, left[l])
	}, func(r int) {
		merged = append(merged, right[r])
	})
	fmt.Println(merged)

	// Output:
	// [1 3 4 5 6 7 8 10 11]
	// [1 3 4 5 6 7 8 10 11]
}

func ExampleDiffSortedList() {
	from := []string{"a", "b", "d", "f"}
	to := []string{"b", "c", "d", "g", "h"}

	var extra, missing []string
	DiffSortedList(len(from), len(to), func(f, t int) int {
		if from[f] < to[t] {
			return -1
		}
		if from[f] > to[t] {
			return 1
		}
		return 0
	}, func(f int) {
		extra = append(extra, from[f])
	}, func(t int) {
		missing = append(missing, to[t])
	})
	fmt.Println("extra:", extra)
	fmt.Println("missing:", missing)

	from, to = to, from
	extra, missing = nil, nil
	DiffSortedList(len(from), len(to), func(f, t int) int {
		if from[f] < to[t] {
			return -1
		}
		if from[f] > to[t] {
			return 1
		}
		return 0
	}, func(f int) {
		extra = append(extra, from[f])
	}, func(t int) {
		missing = append(missing, to[t])
	})
	fmt.Println("extra:", extra)
	fmt.Println("missing:", missing)

	to = from
	extra, missing = nil, nil
	DiffSortedList(len(from), len(to), func(f, t int) int {
		if from[f] < to[t] {
			return -1
		}
		if from[f] > to[t] {
			return 1
		}
		return 0
	}, func(f int) {
		extra = append(extra, from[f])
	}, func(t int) {
		missing = append(missing, to[t])
	})
	fmt.Println("extra:", extra)
	fmt.Println("missing:", missing)

	// Output:
	// extra: [a f]
	// missing: [c g h]
	// extra: [c g h]
	// missing: [a f]
	// extra: []
	// missing: []
}
