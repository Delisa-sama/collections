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

	"github.com/Delisa-sama/collections/containers/forwardlist"
)

func main() {
	list := forwardlist.NewForwardList(1, 2, 3, 4, 5)
	fmt.Println("Size:", list.Size())

	for it := list.Begin(); !it.Equals(list.End()); it.Next() {
		fmt.Println("Value:", it.Value())
	}
}
```
### Двусвязный список

```go
package main

import (
	"fmt"

	"github.com/Delisa-sama/collections/containers/list"
)

func main() {
	l := list.NewList(1, 2, 3, 4, 5)
	fmt.Println("Size:", l.Size())

	for it := l.Begin(); !it.Equals(l.End()); it.Next() {
		fmt.Println("Value:", it.Value())
	}
}
```
### Вектор

```go
package main

import (
	"fmt"

	"github.com/Delisa-sama/collections/containers/vector"
)

func main() {
	vec := vector.NewVector(1, 2, 3, 4, 5)
	fmt.Println("Size:", vec.Size())

	for it := vec.Begin(); !it.Equals(vec.End()); it.Next() {
		fmt.Println("Value:", it.Value())
	}
}
```

### Множество

```go
package main

import (
	"fmt"

	"github.com/Delisa-sama/collections/containers/set"
)

func main() {
	s := set.NewSet(1, 2, 3, 4, 5)
	fmt.Println("Size:", s.Size())

	for it := s.Begin(); !it.Equals(s.End()); it.Next() {
		fmt.Println("Value:", it.Value())
	}
}
```

### Пример использования Equals и EqualsRanges

```go
package main

import (
	"fmt"

	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/containers"
	"github.com/Delisa-sama/collections/containers/forwardlist"
	"github.com/Delisa-sama/collections/containers/vector"
)

func main() {
	// Создаем коллекции для сравнения
	vec1 := vector.NewVector(1, 2, 3, 4, 5)
	vec2 := vector.NewVector(1, 2, 3, 4, 5)
	list1 := forwardlist.NewForwardList(1, 2, 3, 4, 5)

	// Определяем компаратор для сравнения целых чисел
	cmp := comparator.DefaultComparator[int]()

	// Сравниваем векторы
	areEqual := containers.Equals(vec1.Begin(), vec2.Begin(), cmp)
	fmt.Printf("vec1 and vec2 are equal: %v\n", areEqual)

	// Сравниваем вектор и список
	areEqualRanges := containers.EqualsRanges(vec1.Begin(), vec1.End(), list1.Begin(), list1.End(), cmp)
	fmt.Printf("vec1 and list1 are equal: %v\n", areEqualRanges)
}
```
### Пример использования ForEach

```go
package main

import (
	"fmt"

	"github.com/Delisa-sama/collections/containers"
	"github.com/Delisa-sama/collections/containers/vector"
)

func main() {
	// Создаем вектор
	vec := vector.NewVector(1, 2, 3, 4, 5)

	// Применяем функцию ко всем элементам вектора
	containers.ForEach(vec.Begin(), vec.End(), func(val int) {
		fmt.Println(val)
	})
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
	for it := binaryTree.InOrderBegin(); !it.Equals(binaryTree.InOrderEnd()); it.Next() {
		fmt.Println(it.Value())
	}
}
```

## Лицензия

Этот проект лицензируется на условиях лицензии MIT. Подробности смотрите в файле LICENSE.