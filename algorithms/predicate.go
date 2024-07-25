package algorithms

type unaryPredicate[T any] func(T) bool

type binaryPredicate[T any] func(T, T) bool
