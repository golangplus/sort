// Copyright 2015 The Golang Plus Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package sortp is a plus to standard "sort" package.
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
