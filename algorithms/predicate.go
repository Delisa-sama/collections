package algorithms

// unaryPredicate представляет унарный предикат для элемента типа T.
type unaryPredicate[T any] func(T) bool

// binaryPredicate представляет бинарный предикат для элементов типа T.
type binaryPredicate[T any] func(T, T) bool
