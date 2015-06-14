// Copyright 2015 The Golang Plus Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package sortp is a plus to the standard "sort" package.
*/
package sortp

import (
	"sort"
)

// InterfaceStruct is a struct implementing sort.Interface given closures
type InterfaceStruct struct {
	// Len is the number of elements in the collection.
	LenF func() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	LessF func(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	SwapF func(i, j int)
}

// sort.Interface.Len
func (is InterfaceStruct) Len() int {
	return is.LenF()
}

// sort.Interface.Less
func (is InterfaceStruct) Less(i, j int) bool {
	return is.LessF(i, j)
}

// sort.Interface.Swap
func (is InterfaceStruct) Swap(i, j int) {
	is.SwapF(i, j)
}

// SortF calls sort.Sort by closures. Since Interface.Len always returns a constant,
// it is an int parameter rather than a closure here.
func SortF(Len int, Less func(i, j int) bool, Swap func(i, j int)) {
	sort.Sort(InterfaceStruct{
		LenF: func() int {
			return Len
		},
		LessF: Less,
		SwapF: Swap,
	})
}

// StableF calls sort.Stable by closures. Since Interface.Len always returns a constant,
// it is an int parameter rather than a closure here.
func StableF(Len int, Less func(i, j int) bool, Swap func(i, j int)) {
	sort.Stable(InterfaceStruct{
		LenF: func() int {
			return Len
		},
		LessF: Less,
		SwapF: Swap,
	})
}

// IsSortedF is similar to sort.IsSorted but with closure as arguments.
func IsSortedF(Len int, Less func(i, j int) bool) bool {
	for i := 1; i < Len; i++ {
		if Less(i, i-1) {
			return false
		}
	}
	return true
}

// ReverseLess returns a func which can be used to sort data in a reverse order.
func ReverseLess(Less func(i, j int) bool) func(i, j int) bool {
	return func(i, j int) bool {
		return Less(j, i)
	}
}

// Merge merges two sorted list.
//
// @param LeftLen is the length of the left list.
// @param RightLen is the length of the right list.
// @param Less compares the l-th element of left list and the r-th element of the right list.
// @param AppendLeft appends the l-th element of the left list to the result list.
// @param AppendRight appends the r-th element of the right list to the result list.
func Merge(LeftLen, RightLen int, Less func(l, r int) bool, AppendLeft func(l int), AppendRight func(r int)) {
	l, r := 0, 0
	for l < LeftLen && r < RightLen {
		if Less(l, r) {
			AppendLeft(l)
			l++
		} else {
			AppendRight(r)
			r++
		}
	}
	for ; l < LeftLen; l++ {
		AppendLeft(l)
	}
	for ; r < RightLen; r++ {
		AppendRight(r)
	}
}
