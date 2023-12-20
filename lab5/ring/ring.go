package ring

import "fmt"

type Ring[T comparable] struct {
	slice []T
	currentPos int
}

func NewRing[T comparable]() *Ring[T] {
	return &Ring[T]{
		slice: make([]T, 0),
	}
}

func (r *Ring[T]) Put(v T) {
	if r.currentPos == len(r.slice) - 1 {
		r.slice = append(r.slice, v)
		return
	}
	if r.currentPos == 0 {
		r.slice = append([]T{v}, r.slice...)
		return
	}
	r.slice = append(r.slice[:r.currentPos], append([]T{v}, r.slice[r.currentPos:]...)...)
}

func (r *Ring[T]) Remove() {
	if r.currentPos == len(r.slice) - 1 {
		r.slice = r.slice[:r.currentPos]
		return
	}
	if r.currentPos == 0 {
		r.slice = r.slice[1:]
		return
	}
	r.slice = append(r.slice[:r.currentPos], r.slice[r.currentPos+1:]...)
}

func (r *Ring[T]) Move(steps int) {
	current := r.currentPos + steps
	if current > 0 {
		r.currentPos = current % len(r.slice)
	} else {
		r.currentPos = (current % len(r.slice) + len(r.slice)) % len(r.slice)
	}
}

func (r *Ring[T]) Peek() T {
	return r.slice[r.currentPos]
}

func (r *Ring[T]) Size() int { return len(r.slice) }

func (r *Ring[T]) String() string {
	green := "\033[32m"
	reset := "\033[0m"
	res := ""

	for idx, el := range r.slice {
		if r.currentPos == idx {
			res += fmt.Sprintf("%s%c%s ", green, el, reset)
		} else {
			res += fmt.Sprintf("%c ", el)
		}
	}
	return res
}