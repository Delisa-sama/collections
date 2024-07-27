# Containers Library
## Описание

Библиотека containers предоставляет набор [контейнеров](#контейнеры), итераторов и алгоритмов для работы с коллекциями в языке Go.

## Установка
Для установки библиотеки используйте команду:

```bash
go get github.com/Delisa-sama/collections
```

## Примеры
Примеры можно посмотреть в файле [examples_test.go](examples_test.go).

## Контейнеры
Контейнеры — это объекты, которые служат для хранения других объектов.

Библиотека `collections` предоставляет следующие реализации контейнеров:
### Последовательные
- [**Vector**](#vector)
- [**ForwardList**](#forwardlist)
- [**List**](#list)

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
    "github.com/Delisa-sama/collections"
)

func main() {
    vec := collections.NewVector(1, 2, 3)
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

#### At
```go
func (l *Vector[T]) At(index uint) interfaces.RandomAccessIterator[T]
```
Возвращает итератор на элемент вектора по переданному индексу.

Time complexity: `O(1)`

#### Begin
```go
func (l *Vector[T]) Begin() interfaces.RandomAccessIterator[T]
```
Возвращает итератор на первый элемент вектора.

Time complexity: `O(1)`

#### End
```go
func (l *Vector[T]) End() interfaces.Iterator
```
Возвращает итератор на последний элемент вектора.

Time complexity: `O(1)`

#### RBegin
```go
func (l *Vector[T]) RBegin() interfaces.BidirectionalIterator[T]
```
Возвращает перевернутый итератор на последний элемент вектора.

Time complexity: `O(1)`

#### REnd
```go
func (l *Vector[T]) REnd() interfaces.Iterator
```
Возвращает перевернутый итератор на первый элемент вектора.

Time complexity: `O(1)`

## ForwardList
ForwardList представляет собой односвязный список, предоставляющий функции для работы с коллекцией элементов.
Пример использования

```go
package main

import (
	"fmt"

	"github.com/Delisa-sama/collections"
)

func main() {
	fl := collections.NewForwardList(1, 2, 3)
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

## List
List представляет собой двусвязный список, предоставляющий функции для работы с коллекцией элементов.

### Пример использования
```go
package main

import (
	"fmt"

	"github.com/Delisa-sama/collections"
)

func main() {
	list := collections.NewList(1, 2, 3)
	fmt.Println("Размер списка:", list.Size())
	fmt.Println("Первый элемент:", list.Front())
	fmt.Println("Последний элемент:", list.Back())
	list.PushFront(0)
	fmt.Println("Новый размер списка:", list.Size())
	list.PopFront()
	fmt.Println("Размер после удаления:", list.Size())
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

## Set
Set представляет собой коллекцию уникальных элементов, предоставляющую функции для работы с коллекцией.
### Пример использования

```go

package main

import (
	"fmt"

	"github.com/Delisa-sama/collections"
)

func main() {
	set := collections.NewSet(1, 2, 3)
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
func (s *Set[T]) Insert(value T)
```

Вставляет новый элемент в набор.

Time complexity: `O(1)`

#### Remove
```go
func (s *Set[T]) Remove(value T)
```

Удаляет элемент из набора.

Time complexity: `O(1)`

#### Contains
```go
func (s *Set[T]) Contains(value T) bool
```

Проверяет, содержится ли элемент в наборе.

Time complexity: `O(1)`

## BST
BST (Binary Search Tree) представляет собой структуру данных в виде двоичного дерева поиска.
### Пример использования

```go

package main

import (
	"fmt"

	"github.com/Delisa-sama/collections"
)

func main() {
	bst := collections.NewBST(10, 5, 15)
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

## AVLTree
AVLTree представляет собой самобалансирующееся двоичное дерево поиска.
### Пример использования

```go

package main

import (
	"fmt"

	"github.com/Delisa-sama/collections"
)

func main() {
	avl := collections.NewAVLTree(10, 5, 15)
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

## Лицензия

Этот проект лицензируется на условиях лицензии MIT. Подробности смотрите в файле LICENSE.