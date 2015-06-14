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

	// Output:
	// [1 3 4 5 6 7 8 10 11]
}
