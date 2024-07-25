# Containers Library
## Описание

Библиотека containers предоставляет набор структур данных и итераторов для работы с коллекциями в языке Go. Библиотека включает в себя односвязные списки, двусвязные списки и векторы с соответствующими итераторами.
### Установка

Для установки библиотеки используйте команду:

```bash
go get github.com/Delisa-sama/collections
```
## Использование
### Односвязный список

```go

package main

import (
	"fmt"

	"github.com/Delisa-sama/collections/sequence/forwardlist"
)

func main() {
	list := forwardlist.NewForwardList(1, 2, 3, 4, 5)
	fmt.Println("Size:", l.Size())

	for it := l.Begin(); !l.IsEmpty(); it.Next() {
		fmt.Println("Value:", it.Value())

		// Проверяем что итератор не дошел до конца
		if it.Equals(l.End()) {
			break
		}
	}
}
```
### Двусвязный список

```go

package main

import (
	"fmt"

	"github.com/Delisa-sama/collections/sequence/list"
)

func main() {
	l := list.NewList(1, 2, 3, 4, 5)
	fmt.Println("Size:", l.Size())

	for it := l.Begin(); !l.IsEmpty(); it.Next() {
		fmt.Println("Value:", it.Value())

		// Проверяем что итератор не дошел до конца
		if it.Equals(l.End()) {
			break
		}
	}
}
```
### Вектор

```go

package main

import (
	"fmt"

	"github.com/Delisa-sama/collections/sequence/vector"
)

func main() {
	vec := vector.NewVector(1, 2, 3, 4, 5)
	fmt.Println("Size:", vec.Size())

	for it := vec.Begin(); !vec.IsEmpty(); it.Next() {
		fmt.Println("Value:", it.Value())

		// Проверяем что итератор не дошел до конца
		if it.Equals(vec.End()) {
			break
		}
	}
}
```

### Множество

```go

package main

import (
	"fmt"

	"github.com/Delisa-sama/collections/associative/set"
)

func main() {
	s := set.NewSet(1, 2, 3, 4, 5)
	fmt.Println("Size:", s.Size())

	for it := s.Begin(); !s.IsEmpty(); it.Next() {
		fmt.Println("Value:", it.Value())

		// Проверяем что итератор не дошел до конца
		if it.Equals(s.End()) {
			break
		}
	}
}
```

### Пример использования EqualsByIterators
```go
package main

import (
	"fmt"

	"github.com/Delisa-sama/collections"
	"github.com/Delisa-sama/collections/sequence/list"
	"github.com/Delisa-sama/collections/sequence/vector"
)

func main() {
	l := list.NewList(1, 2, 3, 4, 5)
	vec1 := vector.NewVector(1, 2, 3, 4, 5)
	vec2 := vector.NewVector(1, 2, 3, 5)

	// Определяем компаратор для сравнения целых чисел
	cmp := containers.DefaultComparator[int]()

	// Сравниваем вектор и лист
	areEqual := containers.EqualsByIterators(vec1.Begin(), l.Begin(), cmp)
	fmt.Printf("vec1 and l are equal: %v\n", areEqual)

	// Сравниваем первый и второй векторы
	areEqual = containers.EqualsByIterators(vec1.Begin(), vec2.Begin(), cmp)
	fmt.Printf("vec1 and vec2 are equal: %v\n", areEqual)

	// Перебор и печать элементов второго вектора
	fmt.Println("Iterating over vec2:")
	for it := vec2.Begin(); !vec2.IsEmpty(); it.Next() {
		fmt.Printf("%d ", it.Value())

		// Проверяем что итератор не дошел до конца
		if it.Equals(vec2.End()) {
			break
		}
	}

	emptyVector := vector.NewVector[int]()
	// Проверка на IsEmpty позволяет корректно итерироваться по пустому вектору
	for it := emptyVector.Begin(); !emptyVector.IsEmpty(); it.Next() {
		fmt.Printf("%d ", it.Value())
		if it.Equals(emptyVector.End()) {
			break
		}
	}
}
```

### Бинарное поисковое дерево (BST)

```go
package main

import (
	"fmt"

	"github.com/Delisa-sama/collections/bst"
	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/interfaces"
)

func main() {
	// Создаем бинарное поисковое дерево с компаратором для целых чисел
	binaryTree := bst.NewBST(comparator.DefaultComparator[int](), 4, 3, 5, 1, 0, 2, 6, 8, 7)

	// Печатаем элементы дерева в порядке in-order обхода
	for ; it.HasNext(); it.Next() {
		fmt.Println(it.Value())
	}
}
```

## Лицензия

Этот проект лицензируется на условиях лицензии MIT. Подробности смотрите в файле LICENSE.