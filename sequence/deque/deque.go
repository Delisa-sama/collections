package deque

import (
	"github.com/Delisa-sama/collections/copiable"
	"github.com/Delisa-sama/collections/interfaces"
)

const chunkBufferSize = 8 // Размер буфера одного чанка.

// chunk представляет собой чанк данных, хранящую элементы с возможностью
// добавления и удаления элементов спереди и сзади.
type chunk[T any] struct {
	buf   [chunkBufferSize]T // Буфер фиксированного размера для хранения элементов
	begin uint               // Индекс начала заполненной области буфера
	end   uint               // Индекс конца заполненной области буфера
}

// newChunk создает новый chunk с одним начальным значением.
func newChunk[T any](value T) *chunk[T] {
	return &chunk[T]{
		buf:   [chunkBufferSize]T{value},
		begin: 0,
		end:   1,
	}
}

// pushBack добавляет элемент в конец chunk.
// Возвращает false, если буфер полностью заполнен.
func (c *chunk[T]) pushBack(value T) bool {
	if c.end == chunkBufferSize {
		return false
	}
	c.end++
	c.buf[c.end] = value
	return true
}

// popBack удаляет элемент с конца chunk.
// Возвращает false, если буфер пуст.
func (c *chunk[T]) popBack() bool {
	if c.end == 0 {
		return false
	}
	c.end--
	return true
}

// pushFront добавляет элемент в начало chunk.
// Если в начале нет места, сдвигает элементы вправо для освобождения пространства.
// Возвращает false, если буфер полностью заполнен и сдвиг невозможен.
func (c *chunk[T]) pushFront(value T) bool {
	if c.begin == 0 {
		if c.end == chunkBufferSize {
			return false
		}
		c.shiftRight(1)
	}
	c.begin--
	c.buf[c.begin] = value
	return true
}

// popFront удаляет элемент с начала chunk.
// Возвращает false, если в буфере нет элементов.
func (c *chunk[T]) popFront() bool {
	if c.begin == c.end {
		return false
	}
	c.begin++
	return true
}

// shiftRight сдвигает все элементы буфера вправо на заданное количество позиций.
// Используется для освобождения места в начале буфера.
func (c *chunk[T]) shiftRight(offset uint) {
	for i := uint(0); i < offset; i++ {
		length := len(c.buf)
		// Сдвигаем элементы вправо, начиная с конца буфера
		for j := length - 1; j > 0; j-- {
			c.buf[j] = c.buf[j-1]
		}
	}
	c.begin += offset
	c.end += offset
}

// Deque представляет собой двустороннюю очередь, которая состоит из нескольких чанков.
// Она поддерживает добавление и удаление элементов как с начала, так и с конца деки за O(1).
type Deque[T any] struct {
	chunks []*chunk[T] // Список чанков, содержащих элементы
	size   uint        // Общее количество элементов в деке
}

// NewDeque создает новую пустую двустороннюю очередь или очередь с начальными элементами.
// Элементы распределяются по чанкам.
func NewDeque[T any](items ...T) *Deque[T] {
	d := &Deque[T]{
		chunks: make([]*chunk[T], 0, len(items)/chunkBufferSize),
		size:   0,
	}

	chunkNum := 0
	for i := range items {
		if i%chunkBufferSize == 0 {
			chunkNum++
			d.chunks = append(d.chunks, newChunk(items[i]))
			d.size++
			continue
		}
		d.chunks[chunkNum-1].buf[i%chunkBufferSize] = items[i]
		d.chunks[chunkNum-1].end++
		d.size++
	}

	return d
}

// PushBack добавляет элемент в конец деки.
func (d *Deque[T]) PushBack(value T) {
	if len(d.chunks) == 0 {
		d.chunks = append(d.chunks, newChunk(value))
		d.size++
		return
	}

	if !d.chunks[len(d.chunks)-1].pushBack(value) {
		d.chunks = append(d.chunks, newChunk(value))
	}
	d.size++
}

// Back возвращает последний элемент деки без его удаления.
func (d *Deque[T]) Back() T {
	lastChunk := d.chunks[len(d.chunks)-1]
	return lastChunk.buf[lastChunk.end]
}

// PopBack удаляет последний элемент из деки.
func (d *Deque[T]) PopBack() {
	if len(d.chunks) == 0 {
		return
	}

	lastChunk := d.chunks[len(d.chunks)-1]
	if !lastChunk.popBack() {
		d.chunks = d.chunks[:len(d.chunks)-1]
		if d.size == 0 {
			return
		}
		d.chunks[len(d.chunks)-1].popBack()
	}
	d.size--
}

// PushFront добавляет элемент в начало деки.
func (d *Deque[T]) PushFront(value T) {
	if len(d.chunks) == 0 {
		d.chunks = append(d.chunks, newChunk(value))
		d.size++
		return
	}
	if !d.chunks[0].pushFront(value) {
		d.chunks = append([]*chunk[T]{newChunk(value)}, d.chunks...)
	}
	d.size++
}

// Front возвращает первый элемент деки без его удаления.
func (d *Deque[T]) Front() T {
	return d.chunks[0].buf[d.chunks[0].begin]
}

// PopFront удаляет первый элемент из деки.
func (d *Deque[T]) PopFront() {
	if len(d.chunks) == 0 {
		return
	}
	if !d.chunks[0].popFront() {
		d.chunks = d.chunks[1:]
		if d.size == 0 {
			return
		}
		d.chunks[0].popFront()
	}
	d.size--
}

// At возвращает элемент на указанной позиции в деке.
func (d *Deque[T]) At(index uint) T {
	return *d.AtPtr(index)
}

// AtPtr возвращает указатель на элемент на указанной позиции в деке.
func (d *Deque[T]) AtPtr(index uint) *T {
	firstChunkSize := d.chunks[0].end - d.chunks[0].begin
	if index < firstChunkSize {
		return &d.chunks[0].buf[d.chunks[0].begin+index]
	}
	index -= firstChunkSize

	c := d.chunks[index/chunkBufferSize+1]
	return &c.buf[index%chunkBufferSize]
}

// Size возвращает количество элементов в деке.
func (d *Deque[T]) Size() uint {
	return d.size
}

// IsEmpty проверяет, является ли очередь пустой.
func (d *Deque[T]) IsEmpty() bool {
	return d.size == 0
}

// Copy создает и возвращает копию текущей деки.
func (d *Deque[T]) Copy() copiable.Copiable {
	cpy := NewDeque[T]()
	if d.size == 0 {
		return cpy
	}
	cpy.chunks = make([]*chunk[T], 0, len(d.chunks))
	for _, c := range d.chunks {
		for i := range c.buf {
			if i == 0 {
				cpy.chunks = append(cpy.chunks, newChunk(c.buf[i]))
				continue
			}
			cpy.chunks[len(cpy.chunks)-1].buf[i] = c.buf[i]
		}
	}
	return cpy
}

// Begin возвращает итератор на начало деки.
func (d *Deque[T]) Begin() interfaces.RandomAccessIterator[T] {
	return newIterator(d, 0)
}

// End возвращает итератор на конец деки.
func (d *Deque[T]) End() interfaces.RandomAccessIterator[T] {
	return newIterator(d, d.size)
}
