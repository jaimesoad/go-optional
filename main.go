package optional

import (
	"log"
)

type Optional[T any] struct {
	data T
	err  error
}

func New[T any](data T) Optional[T] {
	return Optional[T]{data: data}
}

func Get[T any](val T, err error) Optional[T] {
	return Optional[T]{data: val, err: err}
}

func (o *Optional[T]) Default(data T) T {
	if o.err != nil {
		return data
	}
	return o.data
}

func (o Optional[T]) Unwrap() T {
	if o.err != nil {
		log.Fatal(o.err.Error())
	}
	return o.data
}

func (o *Optional[T]) Log() {
	if o.err != nil {
		log.Println(o.err.Error())
	}
}

func (o *Optional[T]) Fatal() {
	if o.err != nil {
		log.Fatal(o.err.Error())
	}
}
