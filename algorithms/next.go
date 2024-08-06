package algorithms

import (
	"github.com/Delisa-sama/collections/interfaces"
)

// NextBound продвигает итератор it до тех пор, пока он не достигнет итератора bound.
//
// Аргументы:
// - it: итератор, который нужно продвигать.
// - bound: итератор, до которого нужно продвигать it.
func NextBound(
	it interfaces.Iterator,
	bound interfaces.Iterator,
) {
	for !it.Equals(bound) {
		it.Next()
	}
}
