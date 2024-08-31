## [Copiable](copiable/copiable.go)
Определяет интерфейс для объектов, которые могут быть скопированы.

Для корректной работы алгоритмов, все контейнеры и итераторы обязаны реализовывать интерфейс Copiable.

## Контейнеры
Контейнеры — это объекты, которые служат для хранения других объектов.

Библиотека `collections` предоставляет следующие реализации контейнеров:
### Последовательные
- [**Vector**](#vector)
- [**ForwardList**](#forwardlist)
- [**List**](#list)
- [**Deque**](#deque)

### Ассоциативные
- [**Set**](#set)
- [**BST**](#bst)
- [**AVLTree**](#avltree)

## Vector
`Vector` представляет собой динамический массив (вектор), предоставляющий функции для работы с коллекцией элементов.

### Пример использования

```go
package main

import (
	"fmt"

	"github.com/Delisa-sama/collections/sequence/vector"
)

func main() {
	vec := vector.NewVector(1, 2, 3)
	fmt.Println("Размер вектора:", vec.Size())
	fmt.Println("Последний элемент:", vec.Back())
	vec.PushBack(4)
	fmt.Println("Новый размер вектора:", vec.Size())
	vec.PopBack()
	fmt.Println("Размер после удаления:", vec.Size())
}
```

### Конструкторы
#### NewVector
```go
func NewVector[T any](items ...T) *Vector[T]
```
Создает новый Vector и заполняет его переданными элементами.

Time complexity: `O(n)`, где n — количество переданных элементов.

#### NewVectorFromSlice
```go
func NewVectorFromSlice[T any](items []T) *Vector[T]
```
Создает новый Vector на основе переданного слайса, без копирования элементов.

Time complexity: `O(1)`

### Методы
#### Size
```go
func (l *Vector[T]) Size() uint
```
Возвращает количество элементов в векторе.

Time complexity: `O(1)`

#### IsEmpty
```go
func (l *Vector[T]) IsEmpty() bool
```
Проверяет, что вектор пустой.

Time complexity: `O(1)`

#### PushBack
```go
func (l *Vector[T]) PushBack(value T)
```
Добавляет новый элемент в конец вектора.

Time complexity: `O(1)+`

#### Back
```go
func (l *Vector[T]) Back() T
```
Возвращает последний элемент в векторе.

Time complexity: `O(1)`

#### PopBack
```go
func (l *Vector[T]) PopBack()
```
Удаляет последний элемент из вектора.

Time complexity: `O(1)`

#### Copy
```go
func (l *Vector[T]) Copy() copiable.Copiable
```
Возвращает копию вектора.

Time complexity: `O(n)`, где n — количество элементов в векторе.

#### Erase
```go
func (l *Vector[T]) Erase(begin, end interfaces.Iterator)
```
Удаляет элементы в диапазоне [begin, end) из вектора.

Time complexity: `O(1)`.

#### RemoveRange
```go
func (l *Vector[T]) RemoveRange(from, to uint)
```
Удаляет элементы в диапазоне [from, to).

Time complexity: `O(1)`.

#### Remove
```go
func (l *Vector[T]) Remove(index uint)
```
Удаляет элемент по индексу.

Time complexity: `O(1)`.

#### At
```go
func (l *Vector[T]) At(index uint) T
```
Возвращает элемент вектора по переданному индексу.

### Итераторы вектора
#### Begin
```go
func (l *Vector[T]) Begin() interfaces.RandomAccessIterator[T]
```
Возвращает итератор на первый элемент вектора.

#### End
```go
func (l *Vector[T]) End() interfaces.RandomAccessIterator[T]
```
Возвращает итератор на последний элемент вектора.

#### RBegin
```go
func (l *Vector[T]) RBegin() interfaces.BidirectionalIterator[T]
```
Возвращает перевернутый итератор на последний элемент вектора.

#### REnd
```go
func (l *Vector[T]) REnd() interfaces.Iterator
```
Возвращает перевернутый итератор на первый элемент вектора.

## ForwardList
ForwardList представляет собой односвязный список, предоставляющий функции для работы с коллекцией элементов.
Пример использования

```go
package main

import (
	"fmt"

	"github.com/Delisa-sama/collections/sequence/forwardlist"
)

func main() {
	fl := forwardlist.NewForwardList(1, 2, 3)
	fmt.Println("Размер списка:", fl.Size())
	fmt.Println("Первый элемент:", fl.Front())
	fl.PushFront(0)
	fmt.Println("Новый размер списка:", fl.Size())
	fl.PopFront()
	fmt.Println("Размер после удаления:", fl.Size())
}
```

### Конструкторы
#### NewForwardList

```go
func NewForwardList[T any](items ...T) *ForwardList[T]
```

Создает новый ForwardList и заполняет его переданными элементами.

Time complexity: `O(n)`, где n — количество переданных элементов.

### Методы
#### Size
```go
func (l *ForwardList[T]) Size() uint
```
Возвращает количество элементов в списке.

Time complexity: `O(1)`

#### IsEmpty
```go
func (l *ForwardList[T]) IsEmpty() bool
```

Проверяет, что список пустой.

Time complexity: `O(1)`

#### PushBack
```go
func (l *ForwardList[T]) PushBack(value T)
```

Добавляет новый элемент в конец списка.

Time complexity: `O(1)`

#### Back
```go
func (l *ForwardList[T]) Back() T
```

Возвращает последний элемент в списке.

Time complexity: `O(1)`

#### PopBack
```go
func (l *ForwardList[T]) PopBack()
```

Удаляет последний элемент из списка.

Time complexity: `O(1)`

#### PushFront
```go
func (l *ForwardList[T]) PushFront(value T)
```

Добавляет новый элемент в начало списка.

Time complexity: `O(1)`

#### Front
```go
func (l *ForwardList[T]) Front() T
```

Возвращает первый элемент в списке.

Time complexity: `O(1)`

#### PopFront
```go
func (l *ForwardList[T]) PopFront()
```

Удаляет первый элемент из списка.

Time complexity: `O(1)`

#### Copy
```go
func (l *ForwardList[T]) Copy() copiable.Copiable
```
Возвращает копию списка.

Time complexity: `O(n)`, где n — количество элементов в списке.

#### Erase
```go
func (l *ForwardList[T]) Erase(begin, end interfaces.Iterator)
```
Удаляет элементы в диапазоне [begin, end) из списка.

Time complexity: `O(n)` - где n это Distance(begin, end).

### Итераторы односвязного списка
#### Begin
```go
func (l *ForwardList[T]) Begin() interfaces.ForwardIterator[T]
```
Возвращает итератор на первый элемент списка.

#### End
```go
func (l *ForwardList[T]) End() interfaces.Iterator
```
Возвращает итератор на конец списка.

## List
List представляет собой двусвязный список, предоставляющий функции для работы с коллекцией элементов.

### Пример использования
```go
package main

import (
	"fmt"

	"github.com/Delisa-sama/collections/sequence/list"
)

func main() {
	l := list.NewList(1, 2, 3)
	fmt.Println("Размер списка:", l.Size())
	fmt.Println("Первый элемент:", l.Front())
	fmt.Println("Последний элемент:", l.Back())
	l.PushFront(0)
	fmt.Println("Новый размер списка:", l.Size())
	l.PopFront()
	fmt.Println("Размер после удаления:", l.Size())
}
```

### Конструкторы
#### NewList
```go
func NewList[T any](items ...T) *List[T]
```

Создает новый List и заполняет его переданными элементами.

Time complexity: `O(n)`, где n — количество переданных элементов.

### Методы
#### Size
```go
func (l *List[T]) Size() uint
```

Возвращает количество элементов в списке.

Time complexity: `O(1)`

#### IsEmpty
```go
func (l *List[T]) IsEmpty() bool
```

Проверяет, что список пустой.

Time complexity: `O(1)`

#### PushFront
```go
func (l *List[T]) PushFront(value T)
```

Добавляет новый элемент в начало списка.

Time complexity: `O(1)`

#### PushBack
```go
func (l *List[T]) PushBack(value T)
```

Добавляет новый элемент в конец списка.

Time complexity: `O(1)`

#### Front
```go
func (l *List[T]) Front() T
```

Возвращает первый элемент в списке.

Time complexity: `O(1)`

#### Back
```go
func (l *List[T]) Back() T
```

Возвращает последний элемент в списке.

Time complexity: `O(1)`

#### PopFront
```go
func (l *List[T]) PopFront()
```

Удаляет первый элемент из списка.

Time complexity: `O(1)`

#### PopBack
```go
func (l *List[T]) PopBack()
```

Удаляет последний элемент из списка.

Time complexity: `O(1)`

#### Copy
```go
func (l *List[T]) Copy() copiable.Copiable
```
Возвращает копию списка.

Time complexity: `O(n)`, где n — количество элементов в списке.

#### Erase
```go
func (l *List[T]) Erase(begin, end interfaces.Iterator)
```
Удаляет элементы в диапазоне [begin, end) из списка.

Time complexity: `O(n)` - где n это Distance(begin, end).

### Итераторы двусвязного списка
#### Begin
```go
func (l *List[T]) Begin() interfaces.BidirectionalIterator[T]
```
Возвращает итератор на первый элемент списка.

#### End
```go
func (l *List[T]) End() interfaces.Iterator
```
Возвращает итератор на конец списка.

#### RBegin
```go
func (l *List[T]) RBegin() interfaces.BidirectionalIterator[T]
```
Возвращает перевернутый итератор на последний элемент списка.

#### REnd
```go
func (l *List[T]) REnd() interfaces.Iterator
```
Возвращает итератор на конец списка.

## Deque
`Deque` это контейнер, который позволяет быстро вставлять и удалять как в начале, так и в конце.
Кроме того, вставка и удаление на любом конце двухсторонней очереди никогда не делает недействительными
указатели или ссылки на остальные элементы.

### Пример использования

```go
package main

import (
	"fmt"

	"github.com/Delisa-sama/collections/sequence/deque"
)

func main() {
	deq := deque.NewDeque(1, 2, 3)
	fmt.Println("Размер деки:", deq.Size())
	fmt.Println("Последний элемент:", deq.Back())
	deq.PushBack(4)
	fmt.Println("Новый размер деки:", deq.Size())
	deq.PopBack()
	fmt.Println("Размер после удаления:", deq.Size())
}
```

### Конструкторы
#### NewDeque
```go
func NewDeque[T any](items ...T) *Deque[T]
```
Создает новую Deque и заполняет ее переданными элементами.

Time complexity: `O(n)`, где n — количество переданных элементов.

### Методы
#### Size
```go
func (d *Deque[T]) Size() uint
```
Возвращает количество элементов в деке.

Time complexity: `O(1)`

#### IsEmpty
```go
func (d *Deque[T]) IsEmpty() bool
```
Проверяет, что дека пустая.

Time complexity: `O(1)`

#### PushBack
```go
func (d *Deque[T]) PushBack(value T)
```
Добавляет новый элемент в конец деки.

Time complexity: `O(1)+`

#### Back
```go
func (d *Deque[T]) Back() T
```
Возвращает последний элемент в деке.

Time complexity: `O(1)`

#### PopBack
```go
func (d *Deque[T]) PopBack()
```
Удаляет последний элемент из деки.

Time complexity: `O(1)`

#### PushFront
```go
func (d *Deque[T]) PushFront(value T)
```
Добавляет новый элемент в начало деки.

Time complexity: `O(1)+`

#### Front
```go
func (d *Deque[T]) Front() T
```
Возвращает первый элемент в деке.

Time complexity: `O(1)`

#### PopFront
```go
func (d *Deque[T]) PopFront()
```
Удаляет первый элемент из деки.

Time complexity: `O(1)`

#### Copy
```go
func (d *Deque[T]) Copy() copiable.Copiable
```
Возвращает копию деки.

Time complexity: `O(n)`, где n — количество элементов в деке.

#### At
```go
func (d *Deque[T]) At(index uint) T
```
Возвращает элемент деки по переданному индексу.

Time complexity: `O(1)`

#### AtPtr
```go
func (d *Deque[T]) AtPtr(index uint) *T
```
Возвращает указатель на элемент деки по переданному индексу.

Time complexity: `O(1)`

### Итераторы деки
#### Begin
```go
func (d *Deque[T]) Begin() interfaces.RandomAccessIterator[T]
```
Возвращает итератор на первый элемент деки.

#### End
```go
func (d *Deque[T]) End() interfaces.RandomAccessIterator[T]
```
Возвращает итератор на последний элемент деки.

## Set
Set представляет собой коллекцию уникальных элементов, предоставляющую функции для работы с коллекцией.
### Пример использования

```go

package main

import (
	"fmt"

	"github.com/Delisa-sama/collections/associative/set"
)

func main() {
	set := set.NewSet(1, 2, 3)
	fmt.Println("Размер набора:", set.Size())
	set.Insert(4)
	fmt.Println("Новый размер набора:", set.Size())
	set.Remove(2)
	fmt.Println("Размер после удаления:", set.Size())
	fmt.Println("Набор содержит элемент 3:", set.Contains(3))
}
```

### Конструкторы
#### NewSet
```go
func NewSet[T comparable](items ...T) *Set[T]
```

Создает новый Set и заполняет его переданными элементами.

Time complexity: `O(n)`, где n — количество переданных элементов.

### Методы
#### Size
```go
func (s *Set[T]) Size() uint
```

Возвращает количество элементов в наборе.

Time complexity: `O(1)`

#### IsEmpty
```go
func (s *Set[T]) IsEmpty() bool
```

Проверяет, что набор пустой.

Time complexity: `O(1)`

#### Insert
```go
func (s *Set[K]) Set(k K)
```

Вставляет новый элемент в набор.

Time complexity: `O(1)`

#### Contains
```go
func (s *Set[K]) Contains(k K) bool
```

Проверяет, содержится ли элемент в наборе.

Time complexity: `O(1)`

#### Copy
```go
func (s *Set[K]) Copy() copiable.Copiable
```
Возвращает копию множества.

Time complexity: `O(n)`, где n — количество элементов в множестве.

#### Erase
```go
func (s *Set[K]) Erase(begin, end interfaces.Iterator)
```
Удаляет элементы в диапазоне [begin, end) из множества.

Time complexity: `O(n)` - где n это Distance(begin, end).

### Итераторы множества
#### Begin
```go
func (s *Set[K]) Begin() interfaces.BidirectionalIterator[K]
```
Возвращает итератор на первый элемент множества.

#### End
```go
func (s *Set[K]) End() interfaces.Iterator
```
Возвращает итератор на конец множества.

#### RBegin
```go
func (s *Set[K]) RBegin() interfaces.BidirectionalIterator[K]
```
Возвращает перевернутый итератор на последний элемент множества.

#### REnd
```go
func (s *Set[K]) REnd() interfaces.Iterator
```
Возвращает итератор на конец множества.

## BST
BST (Binary Search Tree) представляет собой структуру данных в виде двоичного дерева поиска.
### Пример использования

```go

package main

import (
	"fmt"

	"github.com/Delisa-sama/collections/associative/bst"
)

func main() {
	bst := bst.NewBST(10, 5, 15)
	bst.Insert(7)
	fmt.Println("Размер дерева:", bst.Size())
	fmt.Println("Дерево содержит элемент 7:", bst.Contains(7))
	bst.Remove(5)
	fmt.Println("Размер после удаления:", bst.Size())
}
```

### Конструкторы
#### NewBST
```go
func NewBST[T comparable](items ...T) *BST[T]
```

Создает новое двоичное дерево поиска и заполняет его переданными элементами.

Time complexity: `O(n log n)`, где n — количество переданных элементов.

### Методы
#### Size
```go
func (t *BST[T]) Size() uint
```

Возвращает количество элементов в дереве.

Time complexity: `O(1)`

#### IsEmpty
```go
func (t *BST[T]) IsEmpty() bool
```

Проверяет, что дерево пустое.

Time complexity: `O(1)`

#### Insert
```go
func (t *BST[T]) Insert(value T)
```

Вставляет новый элемент в дерево.

Time complexity: `O(log n)`

#### Remove
```go
func (t *BST[T]) Remove(value T)
```

Удаляет элемент из дерева.

Time complexity: `O(log n)`

#### Contains
```go
func (t *BST[T]) Contains(value T) bool
```

Проверяет, содержится ли элемент в дереве.

Time complexity: `O(log n)`

#### Copy
```go
func (t *BST[T]) Copy() copiable.Copiable
```
Возвращает копию дерева.

Time complexity: `O(n)`, где n — количество элементов в дереве.

### Итераторы BST
#### InOrder
```go
func (t *BST[T]) InOrderBegin() interfaces.ForwardIterator[T]
func (t *BST[T]) InOrderEnd() interfaces.Iterator
```
Итераторы для обхода дерева в in-order порядке.

#### PreOrder
```go
func (t *BST[T]) PreOrderBegin() interfaces.ForwardIterator[T]
func (t *BST[T]) PreOrderEnd() interfaces.Iterator
```
Итераторы для обхода дерева в pre-order порядке.

#### PostOrder
```go
func (t *BST[T]) PostOrderBegin() interfaces.ForwardIterator[T]
func (t *BST[T]) PostOrderEnd() interfaces.Iterator
```
Итераторы для обхода дерева в post-order порядке.

## AVLTree
AVLTree представляет собой самобалансирующееся двоичное дерево поиска.
### Пример использования

```go

package main

import (
	"fmt"

	"github.com/Delisa-sama/collections/associative/avltree"
)

func main() {
	avl := avltree.NewAVLTree(10, 5, 15)
	avl.Insert(7)
	fmt.Println("Размер дерева:", avl.Size())
	fmt.Println("Дерево содержит элемент 7:", avl.Contains(7))
	avl.Remove(5)
	fmt.Println("Размер после удаления:", avl.Size())
}
```

### Конструкторы
#### NewAVLTree
```go
func NewAVLTree[T comparable](items ...T) *AVLTree[T]
```
Создает новое самобалансирующееся двоичное дерево поиска и заполняет его переданными элементами.

Time complexity: `O(n log n)`, где n — количество переданных элементов.

### Методы
#### Size
```go
func (t *AVLTree[T]) Size() uint
```

Возвращает количество элементов в дереве.

Time complexity: `O(1)`

#### IsEmpty
```go
func (t *AVLTree[T]) IsEmpty() bool
```

Проверяет, что дерево пустое.

Time complexity: `O(1)`

#### Insert
```go
func (t *AVLTree[T]) Insert(value T)
```

Вставляет новый элемент в дерево.

Time complexity: `O(log n)`

#### Remove
```go
func (t *AVLTree[T]) Remove(value T)
```

Удаляет элемент из дерева.

Time complexity: `O(log n)`

#### Contains
```go
func (t *AVLTree[T]) Contains(value T) bool
```

Проверяет, содержится ли элемент в дереве.

Time complexity: `O(log n)`

#### Copy
```go
func (tree *AVLTree[K, V]) Copy() copiable.Copiable
```
Возвращает копию дерева.

Time complexity: `O(n)`, где n — количество элементов в дереве.

### Итераторы AVLTree
#### InOrder
```go
func (tree *AVLTree[K, V]) InOrderBegin() interfaces.ValueIterator[pair.Pair[K, V]]
func (tree *AVLTree[K, V]) InOrderEnd() interfaces.Iterator
```
Итераторы для обхода дерева в in-order порядке.

#### PreOrder
```go
func (tree *AVLTree[K, V]) PreOrderBegin() interfaces.ValueIterator[pair.Pair[K, V]]
func (tree *AVLTree[K, V]) PreOrderEnd() interfaces.Iterator
```
Итераторы для обхода дерева в pre-order порядке.

#### PostOrder
```go
func (tree *AVLTree[K, V]) PostOrderBegin() interfaces.ValueIterator[pair.Pair[K, V]]
func (tree *AVLTree[K, V]) PostOrderEnd() interfaces.Iterator
```
Итераторы для обхода дерева в post-order порядке.
