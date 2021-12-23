package vector

type Vector[T any] []T

type Fn[T any] func(item T)

func (v *Vector[T]) Push(m T) {
	*v = append(*v, m)
}

func New[T any]() *Vector[T] {
	var vec Vector[T]
	return &vec
}

func From[T any](items ...T) *Vector[T] {
	var vec Vector[T]
	for _, i := range items {
		vec.Push(i)
	}
	return &vec
}

func (v *Vector[T]) Each(fn Fn[T]) {
	for _, i := range *v {
		fn(i)
	}
}
