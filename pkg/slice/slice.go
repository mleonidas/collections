package slice

import (
	"constraints"
	"sort"
)

// Slice higher level type for slices/arrays, if this was going to be used
// in a production context would probably change this to be a struct and put the
// generic array as a field in the struct so we can lock / unlock reads and
// clones
type Slice[T any] []T

type Fn[T any] func(item T)

// Push the quickest way to implement this
func (v *Slice[T]) Push(m T) {
	*v = append(*v, m)
}

// Index returns the item at a given index
func Index[T comparable](v []T, el T) int {
	// loop through the slice ... self explanatory
	for i, e := range v {
		if e == el {
			return i
		}
	}
	return -1
}

// New this might not be needed at all lol
func New[T any]() *Slice[T] {
	var vec Slice[T]
	return &vec
}

// Sort returns a copy of the slice that is sorted, unlike golangs sort.X
// it is not mutating the original object
func Sort[T constraints.Ordered](v []T) []T {

	if v == nil {
		return nil
	}

	vCopy := make([]T, len(v))

	copy(vCopy, v)

	sort.Slice(vCopy, func(i int, j int) bool {
		return vCopy[i] <= vCopy[j]
	})

	return vCopy
}

// From a nod to Kotlins arrayOf taking arguments as a variadic
func From[T any](items ...T) *Slice[T] {
	var vec Slice[T]
	for _, i := range items {
		vec.Push(i)
	}
	return &vec
}

// Each mostly a nod to ruby, loop over each item in the slice and call
// the function passed as an argument with the item as a paramater
func (v *Slice[T]) Each(fn Fn[T]) {
	for _, i := range *v {
		fn(i)
	}
}

// Uniq given a list of comparable items iterate and remove dups
func Uniq[T comparable](v []T) []T {
	keys := make(map[T]bool)
	var list Slice[T]

	for _, entry := range v {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list.Push(entry)
		}
	}

	return list
}
