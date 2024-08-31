# Адаптеры
Адаптеры контейнеров — это интерфейсы, созданные путем ограничения функциональности уже существующего контейнера и предоставления другого набора функций.
Когда вы объявляете адаптеры контейнера, у вас есть возможность указать, какой последовательный контейнер будет базовым.

## Адаптеры
- [**Stack**](#stack)
- [**Queue**](#queue)
- [**PriorityQueue**](#priorityqueue)

## Stack
Стек — это адаптер, обеспечивающий доступ «последним пришел — первым вышел» (LIFO) к элементам базового контейнера.
### Пример использования

```go
package main

import (
	"fmt"

	"github.com/Delisa-sama/collections/adapters/stack"
	"github.com/Delisa-sama/collections/sequence/vector"
)

func main() {
	stack := stack.NewStack(vector.NewVector[int])
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println("Размер стека:", stack.Size())
	fmt.Println("Верхний элемент:", stack.Top())
	stack.Pop()
	fmt.Println("Новый верхний элемент:", stack.Top())
	fmt.Println("Размер после удаления:", stack.Size())
}
```

### Требования к интерфейсу контейнера
- `Size() uint` - возвращает количество элементов в контейнере.
- `IsEmpty() bool` - проверяет, что контейнер пустой.
- `Back() T` - возвращает последний элемент в контейнере.
- `PushBack(T)` - добавляет элемент в конец контейнера.
- `PopBack()` - удаляет последний элемент из контейнера.

### Конструкторы
#### NewStack
```go
func NewStack[T any, C container[T]](cc containerConstructor[T, C], items ...T) *Stack[T, C]
```
Создает новый Stack и заполняет его переданными элементами.

Time complexity: `O(n)`, где n — количество переданных элементов.

### Методы
#### Push
```go
func (s *Stack[T, C]) Push(v T)
```
Добавляет новый элемент на вершину стека.

Time complexity: Time complexity метода `PushBack` базового контейнера.

#### Pop
```go
func (s *Stack[T, C]) Pop()
```
Удаляет верхний элемент из стека.

Time complexity: Time complexity метода `PopBack` базового контейнера.

#### Top
```go
func (s *Stack[T, C]) Top() T
```
Возвращает верхний элемент стека, не удаляя его.

Time complexity: Time complexity метода `Back` базового контейнера.

#### Size
```go
func (s *Stack[T, C]) Size() uint
```
Возвращает количество элементов в стеке.

Time complexity: Time complexity метода `Size` базового контейнера.

#### IsEmpty
```go
func (s *Stack[T, C]) IsEmpty() bool
```
Проверяет, что стек пустой.

Time complexity: Time complexity метода `IsEmpty` базового контейнера.

#### Copy
```go
func (s *Stack[T, C]) Copy() *Stack[T, C]
```
Возвращает копию стека, копирую контейнер.

Time complexity: O(n), где n — количество элементов в контейнере.

## Queue
Очередь — это адаптер, обеспечивающий доступ «первым пришел — первым вышел» (FIFO) к элементам базового контейнера.
### Пример использования

```go
package main

import (
	"fmt"

	"github.com/Delisa-sama/collections/adapters/queue"
	"github.com/Delisa-sama/collections/sequence/deque"
)

func main() {
	q := queue.NewQueue(deque.NewDeque[int])
	q.PushBack(1)
	q.PushBack(2)
	q.PushBack(3)
	fmt.Println("Размер очереди:", q.Size())
	fmt.Println("Верхний элемент:", q.Front())
	q.PopFront()
	fmt.Println("Новый верхний элемент:", q.Front())
	fmt.Println("Размер после удаления:", q.Size())
}
```

### Требования к интерфейсу контейнера
- `Size() uint` - возвращает количество элементов в контейнере.
- `IsEmpty() bool` - проверяет, что контейнер пустой.
- `Back() T` - возвращает последний элемент в контейнере.
- `Front() T` - возвращает первый элемент в контейнере.
- `PushBack(T)` - добавляет элемент в конец контейнера.
- `PopFront()` - удаляет первый элемент из контейнера.

### Конструкторы
#### NewQueue
```go
func NewQueue[T any, C container[T]](cc containerConstructor[T, C], items ...T) *Queue[T, C]
```
Создает новую Queue и заполняет ее переданными элементами.

Time complexity: `O(n)`, где n — количество переданных элементов.

### Методы
#### Front
```go
func (q *Queue[T, C]) Front() T
```
Возвращает первый элемент в очереди (без его удаления).

Time complexity: Time complexity метода `Front` базового контейнера.

#### Back
```go
func (q *Queue[T, C]) Back() T
```
Возвращает последний элемент в очереди (без его удаления).

Time complexity: Time complexity метода `Back` базового контейнера.

#### PushBack
```go
func (q *Queue[T, C]) PushBack(value T)
```
Добавляет элемент в конец очереди.

Time complexity: Time complexity метода `PushBack` базового контейнера.

#### PopFront
```go
func (q *Queue[T, C]) PopFront()
```
Добавляет элемент в конец очереди.

Time complexity: Time complexity метода `PopFront` базового контейнера.

#### Size
```go
func (q *Queue[T, C]) Size() uint
```
Возвращает количество элементов в очереди.

Time complexity: Time complexity метода `Size` базового контейнера.

#### IsEmpty
```go
func (q *Queue[T, C]) IsEmpty() bool
```
Проверяет, что очередь пуста.

Time complexity: Time complexity метода `IsEmpty` базового контейнера.

#### Copy
```go
func (q *Queue[T, C]) Copy() *Stack[T, C]
```
Возвращает копию очереди, копирую контейнер.

Time complexity: O(n), где n — количество элементов в контейнере.

## PriorityQueue
Очередь с приоритетами — это адаптер, который обеспечивает постоянный поиск самого большого (по умолчанию).
### Пример использования

```go
package main

import (
	"fmt"

	"github.com/Delisa-sama/collections/adapters/priorityqueue"
	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/sequence/deque"
)

func main() {
	q := priorityqueue.NewPriorityQueue(deque.NewDeque[int], comparator.DefaultLess[int]())
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println("Размер очереди:", q.Size())
	fmt.Println("Верхний элемент:", q.Top())
	q.Pop()
	fmt.Println("Новый верхний элемент:", q.Top())
	fmt.Println("Размер после удаления:", q.Size())
}
```

### Требования к интерфейсу контейнера
- `Size() uint` - возвращает количество элементов в контейнере.
- `IsEmpty() bool` - проверяет, что контейнер пустой.
- `Front() T` - возвращает первый элемент в контейнере.
- `PushBack(T)` - добавляет элемент в конец контейнера.
- `PopFront()` - удаляет первый элемент из контейнера.
- `Begin()` - возвращает RandomAccessIterator, указывающий на первый элемент контейнера.
- `End()` - возвращает RandomAccessIterator, указывающий на последний элемент контейнера.

### Конструкторы
#### NewPriorityQueue
```go
func NewPriorityQueue[T any, C container[T]](
    cc containerConstructor[T, C],
    comp comparator.Less[T],
    items ...T,
) *PriorityQueue[T, C]
```
Создает новую PriorityQueue и заполняет ее переданными элементами.

Time complexity: `O(n + n log n)`, где n — количество переданных элементов 
(с учетом худшей сложности алгоритма сортировки).

### Методы
#### Push
```go
func (q *PriorityQueue[T, C]) Push(value T)
```
Добавляет новый элемент в приоритетную очередь. После добавления, элементы переупорядочиваются
согласно компаратору, чтобы поддерживать правильный порядок приоритетов.

Time complexity: Time complexity метода `Push` базового контейнера + O(n log n) сложность сортировки в худшем случае.

#### Top
```go
func (q *PriorityQueue[T, C]) Top() T
```
Возвращает элемент с наивысшим приоритетом (первый элемент в очереди) без его удаления.

Time complexity: Time complexity метода `Front` базового контейнера.

#### Pop
```go
func (q *PriorityQueue[T, C]) Pop()
```
Удаляет элемент с наивысшим приоритетом (первый элемент) из очереди.

Time complexity: Time complexity метода `PopFront` базового контейнера.

#### Size
```go
func (q *PriorityQueue[T, C]) Size() uint
```
Возвращает количество элементов в очереди.

Time complexity: Time complexity метода `Size` базового контейнера.

#### IsEmpty
```go
func (q *PriorityQueue[T, C]) IsEmpty() bool
```
Проверяет, что очередь пуста.

Time complexity: Time complexity метода `IsEmpty` базового контейнера.

#### Copy
```go
func (q *PriorityQueue[T, C]) Copy() *Stack[T, C]
```
Возвращает копию очереди, копирую контейнер.

Time complexity: O(n), где n — количество элементов в контейнере.
