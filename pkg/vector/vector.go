package vector

// Vector higher level type for slices/arrays, if this was going to be used
// in a production context would probably change this to be a struct and put the
// generic array as a field in the struct so we can lock / unlock reads and
// clones
type Vector[T any] []T

type Fn[T any] func(item T)

// Push, the quickest way to implement this
func (v *Vector[T]) Push(m T) {
	*v = append(*v, m)
}

// New this might not be needed at all lol
func New[T any]() *Vector[T] {
	var vec Vector[T]
	return &vec
}

// From a nod to Kotlins arrayOf taking arguments as a variadic
func From[T any](items ...T) *Vector[T] {
	var vec Vector[T]
	for _, i := range items {
		vec.Push(i)
	}
	return &vec
}

// Each mostly a nod to ruby, loop over each item in the vector and call
// the function passed as an argument with the item as a paramater
func (v *Vector[T]) Each(fn Fn[T]) {
	for _, i := range *v {
		fn(i)
	}
}

// Uniq given a list of comparable items iterate and remove dups
func Uniq[T comparable](v []T) []T {
	keys := make(map[T]bool)
	var list Vector[T]

	for _, entry := range v {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list.Push(entry)
		}
	}

	return list
}
