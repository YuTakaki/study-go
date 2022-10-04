package types

type Todo struct {
	Todo       string
	IsComplete bool
}

type SliceFilterType func(index int) bool

type SliceMapType func(x Todo, index int) Todo