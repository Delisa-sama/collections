# Алгоритмы

## Сортировка
- [Sort](#sort)
- [SortC](#sortc)

## Пакетные операции
- [Advance](#advance)
- [AdvanceCopy](#advancecopy)
- [ForEach](#foreach)
- [ForEachPtr](#foreachptr)
- [ForEachIdx](#foreachidx)
- [ForEachIdxPtr](#foreachidxptr)
- [Distance](#distance)
- [Equals](#equals)
- [EqualsC](#equalsc)
- [EqualsRanges](#equalsranges)
- [EqualsRangesC](#equalsrangesc)
- [NextBound](#nextbound)
- [Print](#print)
- [Println](#println)
- [PrintFunc](#printfunc)
- [PrintF](#printf)
- [PrintFFunc](#printffunc)

## Поиск
- [Find](#find)
- [FindC](#findc)
- [FindIf](#findif)
- [FindIfNot](#findifnot)
- [FindFirstOf](#findfirstof)
- [FindFirstOfC](#findfirstofc)
- [FindFirstOfIf](#findfirstofif)
- [AllOf](#allof)
- [AnyOf](#anyof)
- [NoneOf](#noneof)
- [Count](#count)
- [CountC](#countc)
- [CountIf](#countif)

# Бинарный поиск
- [LowerBound](#lowerbound)
- [LowerBoundC](#lowerboundc)
- [UpperBound](#upperbound)
- [UpperBoundC](#upperboundc)

## Свертки
- [Accumulate](#accumulate)
- [FoldLeft](#foldleft)
- [FoldLeftFirst](#foldleftfirst)
- [FoldRight](#foldright)
- [FoldRightLast](#foldrightlast)

# Минимум/Максимум
- [Min](#min)
- [MinC](#minc)
- [Max](#max)
- [MaxC](#maxc)
- [MinMax](#minmax)
- [MinMaxC](#minmaxc)

# Генераторы
- [Generate](#generate)
- [GenerateN](#generaten)
- [Fill](#fill)

# Трансформации
- [TransformUnary](#transformunary)
- [TransformBinary](#transformbinary)
- [Replace](#replace)
- [ReplaceIf](#replaceif)
- [ReplaceCopy](#replacecopy)
- [ReplaceCopyIf](#replacecopyif)
- [Swap](#swap)
- [SwapIter](#swapiter)
- [SwapRanges](#swapranges)
- [Remove](#remove)
- [RemoveC](#removec)
- [RemoveIf](#removeif)
- [Reverse](#reverse)
- [ReverseCopy](#reversecopy)
- [Rotate](#rotate)
- [RotateCopy](#rotatecopy)
- [Unique](#unique)
- [UniqueC](#uniquec)
- [UniqueIf](#uniqueif)
- [Copy](#copy)
- [CopyIf](#copyif)

# Операции с кучей
- [MakeHeap](#makeheap)
- [SortHeap](#sortheap)
- [PopHeap](#popheap)
- [PushHeap](#pushheap)

### [AllOf](algorithms/all_of.go)
```go
func AllOf[T any](begin interfaces.ForwardIterator[T], end interfaces.Iterator, predicate unaryPredicate[T]) bool
```
Проверяет, удовлетворяют ли все элементы в диапазоне [begin, end) предикату.

### [AnyOf](algorithms/any_of.go)
```go
func AnyOf[T any](begin interfaces.ValueIterator[T], end interfaces.Iterator, predicate unaryPredicate[T]) bool
```
Проверяет, удовлетворяет ли хотя бы один элемент в диапазоне [begin, end) предикату.

### [NoneOf](algorithms/none_of.go)
```go
func NoneOf[T any](begin interfaces.ValueIterator[T], end interfaces.Iterator, predicate unaryPredicate[T]) bool
```
Проверяет, удовлетворяет ли ни один элемент в диапазоне [begin, end) предикату.

### [Copy](algorithms/copy.go)
```go
func Copy[T any](
    begin interfaces.ValueIterator[T], end interfaces.Iterator,
    destBegin interfaces.PointerIterator[T],
) interfaces.PointerIterator[T]
```
Копирует элементы из диапазона [begin, end) в диапазон, начинающийся с destBegin.

### [CopyIf](algorithms/copy.go)
```go
func CopyIf[T any](
    begin interfaces.ValueIterator[T], end interfaces.Iterator,
    destBegin interfaces.PointerIterator[T],
    predicate unaryPredicate[T],
) interfaces.PointerIterator[T]
```
Копирует элементы из диапазона [begin, end), которые удовлетворяют предикату, в диапазон, начинающийся с destBegin.

### [CountC](algorithms/count.go)
```go
func CountC[T any](
    begin interfaces.ValueIterator[T],
    end interfaces.Iterator,
    value T,
    cmp comparator.Comparator[T],
) uint 
```
Подсчитывает количество элементов в диапазоне [begin, end), равных заданному значению, используя пользовательский компаратор.

### [Count](algorithms/count.go)
```go
func Count[T comparable](
    begin interfaces.ValueIterator[T],
    end interfaces.Iterator,
    value T,
) uint
```
Подсчитывает количество элементов в диапазоне [begin, end), равных заданному значению, используя оператор сравнения ==.

### [CountIf](algorithms/count.go)
```go
func CountIf[T any](
    begin interfaces.ValueIterator[T],
    end interfaces.Iterator,
    predicate unaryPredicate[T],
) uint
```
Подсчитывает количество элементов в диапазоне [begin, end), удовлетворяющих предикату.

### [EqualsC](algorithms/equals.go)
```go
func EqualsC[T any](a interfaces.ForwardIterator[T], b interfaces.ForwardIterator[T], cmp comparator.Comparator[T]) bool
```
Проверяет, равны ли все элементы двух диапазонов [a, b), используя пользовательский компаратор.

### [Equals](algorithms/equals.go)
```go
func Equals[T comparable](a interfaces.ForwardIterator[T], b interfaces.ForwardIterator[T]) bool
```
Проверяет, равны ли все элементы двух диапазонов [a, b), используя оператор сравнения ==.

### [EqualsRangesC](algorithms/equals.go)
```go
func EqualsRangesC[T any](
	aBegin interfaces.ValueIterator[T], aEnd interfaces.Iterator,
	bBegin interfaces.ValueIterator[T], bEnd interfaces.Iterator,
	cmp comparator.Comparator[T],
) bool
```
Проверяет, равны ли все элементы двух диапазонов [aBegin, aEnd) и [bBegin, bEnd), используя пользовательский компаратор.

### [EqualsRanges](algorithms/equals.go)
```go
func EqualsRanges[T comparable](
	aBegin interfaces.ValueIterator[T], aEnd interfaces.Iterator,
	bBegin interfaces.ValueIterator[T], bEnd interfaces.Iterator,
) bool
```
Проверяет, равны ли все элементы двух диапазонов [aBegin, aEnd) и [bBegin, bEnd), используя оператор сравнения ==.

### [FindC](algorithms/find.go)
```go
func FindC[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	value T,
	cmp comparator.Comparator[T],
) (interfaces.ValueIterator[T], bool)
```
Выполняет поиск элемента в диапазоне [begin, end) с использованием пользовательского компаратора.

Функция возвращает итератор на найденный элемент и булево значение, указывающее на успех поиска.

### [Find](algorithms/find.go)
```go
func Find[T comparable](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	value T,
) (interfaces.ValueIterator[T], bool)
```
Выполняет поиск элемента в диапазоне [begin, end) с использованием оператора сравнения ==.

Функция возвращает итератор на найденный элемент и булево значение, указывающее на успех поиска.

### [FindIf](algorithms/find.go)
```go
func FindIf[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	predicate unaryPredicate[T],
) (interfaces.ValueIterator[T], bool)
```
Выполняет поиск элемента в диапазоне [begin, end), для которого предикат возвращает true.

Функция возвращает итератор на найденный элемент и булево значение, указывающее на успех поиска.

### [FindIfNot](algorithms/find.go)
```go
func FindIfNot[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	predicate unaryPredicate[T],
) (interfaces.ValueIterator[T], bool)
```
Выполняет поиск элемента в диапазоне [begin, end), для которого предикат возвращает false.

Функция возвращает итератор на найденный элемент и булево значение, указывающее на успех поиска.

### [FindFirstOfC](algorithms/find_first_of.go)
```go
func FindFirstOfC[T any](
	begin interfaces.ValueIterator[T], end interfaces.Iterator,
	sBegin interfaces.ValueIterator[T], sEnd interfaces.Iterator,
	cmp comparator.Comparator[T],
) (interfaces.ValueIterator[T], bool)
```
Ищет первый элемент из диапазона [begin, end), который также содержится в диапазоне [sBegin, sEnd), используя пользовательский компаратор.

Функция возвращает итератор на найденный элемент и булево значение, указывающее на успех поиска.

### [FindFirstOf](algorithms/find_first_of.go)
```go
func FindFirstOf[T comparable](
	begin interfaces.ValueIterator[T], end interfaces.Iterator,
	sBegin interfaces.ValueIterator[T], sEnd interfaces.Iterator,
) (interfaces.ValueIterator[T], bool)
```
Ищет первый элемент из диапазона [begin, end), который также содержится в диапазоне [sBegin, sEnd), используя оператор сравнения ==.

Функция возвращает итератор на найденный элемент и булево значение, указывающее на успех поиска.

### [FindFirstOfIf](algorithms/find_first_of.go)
```go
func FindFirstOfIf[T any](
	begin interfaces.ValueIterator[T], end interfaces.Iterator,
	sBegin interfaces.ValueIterator[T], sEnd interfaces.Iterator,
	predicate binaryPredicate[T],
) (interfaces.ValueIterator[T], bool)
```
Ищет первый элемент из диапазона [begin, end), который также содержится в диапазоне [sBegin, sEnd), для которого предикат возвращает true.

Функция возвращает итератор на найденный элемент и булево значение, указывающее на успех поиска.

### [ForEach](algorithms/for_each.go)
```go
func ForEach[T any](begin interfaces.ValueIterator[T], end interfaces.Iterator, f forEachFunc[T])
```
Применяет функцию к каждому элементу в диапазоне [begin, end).

### [ForEachPtr](algorithms/for_each.go)
```go
func ForEachPtr[T any](begin interfaces.PointerIterator[T], end interfaces.Iterator, f forEachPtrFunc[T])
```
Применяет функцию к указателю на каждый элемент в диапазоне [begin, end).

### [ForEachIdx](algorithms/for_each.go)
```go
func ForEachIdx[T any](begin interfaces.ValueIterator[T], end interfaces.Iterator, f forEachIdxFunc[T])
```
Применяет функцию к каждому элементу в диапазоне [begin, end),
передавая в функцию элемент и его индекс.

### [ForEachIdxPtr](algorithms/for_each.go)
```go
func ForEachIdxPtr[T any](begin interfaces.ValueIterator[T], end interfaces.Iterator, f forEachIdxFunc[T])
```
Применяет функцию к указателю каждый элемент в диапазоне [begin, end),
передавая в функцию указатель на элемент и его индекс.

### [Print](algorithms/print.go)
```go
func Print[T any](
    begin interfaces.ValueIterator[T], end interfaces.Iterator,
) (n int, err error)
```
Выводит значения, начиная с итератора begin до итератора end, используя fmt.Print.

### [Println](algorithms/print.go)
```go
func Println[T any](
	begin interfaces.ValueIterator[T], end interfaces.Iterator,
) (n int, err error)
```
Выводит значения, начиная с итератора begin до итератора end, используя fmt.Println.

### [PrintFunc](algorithms/print.go)
```go
func PrintFunc[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	f printFunc,
) (totalN int, err error)
```
Выводит значения, начиная с итератора begin до итератора end, используя заданную функцию вывода.

### [PrintF](algorithms/print.go)
```go
func PrintF[T any](
	begin interfaces.ValueIterator[T], end interfaces.Iterator,
	format string,
) (n int, err error)
```
Выводит значения, начиная с итератора begin до итератора end, используя fmt.Printf и заданный формат.

### [PrintFFunc](algorithms/print.go)
```go
func PrintFFunc[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	format string,
	f printfFunc,
) (totalN int, err error)
```
Выводит значения, начиная с итератора begin до итератора end, используя заданную функцию форматированного вывода.

### [Swap](algorithms/swap.go)
```go
func Swap[T any](a, b *T)
```
Производит обмен значениями двух указателей.

### [SwapIter](algorithms/swap.go)
```go
func SwapIter[T any](a, b interfaces.PointerIterator[T])
```
Производит обмен значениями, на которые указывают два итератора.

### [SwapRanges](algorithms/swap.go)
```go
func SwapRanges[T any](
	aBegin interfaces.PointerIterator[T],
	aEnd interfaces.Iterator,
	bBegin interfaces.PointerIterator[T],
)
```
Производит обмен значениями между двумя диапазонами элементов.

### [Rotate](algorithms/rotate.go)
```go
func Rotate[T any](begin, middle interfaces.ForwardIterator[T], end interfaces.Iterator) interfaces.Iterator
```
Выполняет циклический сдвиг элементов в диапазоне [begin, end) так, чтобы элемент middle стал первым элементом диапазона.
Элементы, находящиеся между begin и middle, будут перемещены в конец диапазона.

### [RotateCopy](algorithms/rotate.go)
```go
func RotateCopy[T any](
    begin, nBegin interfaces.ForwardIterator[T],
    end interfaces.Iterator,
    destBegin interfaces.PointerIterator[T],
) interfaces.PointerIterator[T]
```
Выполняет циклический сдвиг элементов в диапазоне [begin, end) так, чтобы элемент nBegin стал первым элементом диапазона,
и копирует результат в другой диапазон, начинающийся с destBegin.

### [Fill](algorithms/fill.go)
```go
func Fill[T any](begin interfaces.PointerIterator[T], end interfaces.Iterator, value T)
```
Заполняет диапазон [begin, end) значением value.

### [Remove](algorithms/remove.go)
```go
func Remove[T comparable](
    begin interfaces.ForwardIterator[T],
    end interfaces.Iterator,
    value T,
) interfaces.ForwardIterator[T]
```
Удаляет все элементы, равные заданному значению `value`, из диапазона [begin, end).
Элементы, которые не равны `value`, сохраняются в начале диапазона, а оставшиеся элементы не изменяются.

### [RemoveC](algorithms/remove.go)
```go
func RemoveC[T any](
    begin interfaces.ForwardIterator[T],
    end interfaces.Iterator,
    value T,
    cmp comparator.Comparator[T],
) interfaces.ForwardIterator[T]
```
Удаляет все элементы, равные заданному значению `value`, из диапазона [begin, end),
используя пользовательский компаратор для сравнения элементов.
Элементы, которые не равны `value`, сохраняются в начале диапазона, а оставшиеся элементы не изменяются.

### [RemoveIf](algorithms/remove.go)
```go
func RemoveIf[T any](
    begin interfaces.ForwardIterator[T],
    end interfaces.Iterator,
    predicate unaryPredicate[T],
) interfaces.ForwardIterator[T]
```
Удаляет все элементы из диапазона [begin, end), которые удовлетворяют заданному предикату `predicate`.
Элементы, которые не удовлетворяют предикату, сохраняются в начале диапазона, а оставшиеся элементы не изменяются.

### [Unique](algorithms/unique.go)
```go
func Unique[T comparable](
    begin interfaces.ForwardIterator[T],
    end interfaces.Iterator,
) interfaces.ForwardIterator[T]
```
Удаляет дублирующиеся последовательные элементы в диапазоне [begin, end), оставляя только первые
вхождения каждого элемента. Элементы считаются дубликатами, если они равны друг другу (==).

### [UniqueC](algorithms/unique.go)
```go
func UniqueC[T any](
    begin interfaces.ForwardIterator[T],
    end interfaces.Iterator,
    cmp comparator.Comparator[T],
) interfaces.ForwardIterator[T]
```
Удаляет дублирующиеся последовательные элементы в диапазоне [begin, end), оставляя только первые
вхождения каждого элемента. Элементы считаются дубликатами, если они равны согласно пользовательскому компаратору.

### [UniqueIf](algorithms/unique.go)
```go
func UniqueIf[T any](
    begin interfaces.ForwardIterator[T],
    end interfaces.Iterator,
    predicate binaryPredicate[T],
) interfaces.ForwardIterator[T]
```
Удаляет дублирующиеся последовательные элементы в диапазоне [begin, end), оставляя только первые
вхождения каждого элемента. Элементы считаются дубликатами, если они удовлетворяют условию заданному предикатом.

### [Accumulate](algorithms/accumulate.go)
```go
func Accumulate[T interfaces.Numeric](
    begin interfaces.ValueIterator[T],
    end interfaces.Iterator,
    init T,
) T
```
Вычисляет сумму всех элементов в диапазоне [begin, end) начиная с начального значения init.
Функция проходит по всем элементам диапазона, добавляя каждый элемент к переменной init.

### [Distance](algorithms/distance.go)
```go
func Distance[T any](begin, end interfaces.Iterator) uint
```
Вычисляет количество элементов в диапазоне [begin, end).
Если итераторы поддерживают случайный доступ (RandomAccessIterator), вычисление производится за константное время.
В противном случае используется линейный обход диапазона для подсчета количества элементов.

### [Advance](algorithms/advance.go)
```go
func Advance[T any](it interfaces.Iterator, n int)
```
Продвигает итератор it на n шагов вперед или назад.

Функция поддерживает три вида итераторов:
- RandomAccessIterator: итератор с произвольным доступом. Если итератор реализует этот интерфейс,
  он сдвигается на n шагов с помощью метода Shift. Этот метод оптимален по скорости для итераторов
  с произвольным доступом.
- BidirectionalIterator: итератор с двусторонним доступом. Если итератор реализует этот интерфейс,
  и n отрицательное, итератор сдвигается назад на n шагов с помощью метода Prev.
- UnidirectionalIterator: итератор с однонаправленным доступом. Если итератор не поддерживает
  произвольный или двусторонний доступ, он сдвигается на n шагов вперед с помощью метода Next.

### [AdvanceCopy](algorithms/advance.go)
```go
func AdvanceCopy[T any, It interfaces.Iterator](it It, n int) It
```
Продвигает копию итератора it на n шагов вперед или назад и возвращает его.

### [LowerBound](algorithms/bounds.go)
```go
func LowerBound[T cmp.Ordered](
    begin interfaces.ValueIterator[T],
    end interfaces.Iterator,
    value T,
) interfaces.ValueIterator[T]
```
Находит первый элемент, который не меньше чем значение value в отсортированном диапазоне [begin, end).

### [LowerBoundC](algorithms/bounds.go)
```go
func LowerBoundC[T any](
    begin interfaces.ValueIterator[T],
    end interfaces.Iterator,
    value T,
    cmp comparator.Comparator[T],
) interfaces.ValueIterator[T]
```
Находит первый элемент, который не меньше чем значение value в отсортированном диапазоне [begin, end),
используя пользовательский компаратор.

### [UpperBound](algorithms/bounds.go)
```go
func UpperBound[T cmp.Ordered](
    begin interfaces.ValueIterator[T],
    end interfaces.Iterator,
    value T,
) interfaces.ValueIterator[T]
```
Находит первый элемент, который больше чем значение value в отсортированном диапазоне [begin, end).

### [UpperBoundC](algorithms/bounds.go)
```go
func UpperBoundC[T any](
    begin interfaces.ValueIterator[T],
    end interfaces.Iterator,
    value T,
    cmp comparator.Comparator[T],
) interfaces.ValueIterator[T]
```
Находит первый элемент, который больше чем значение value в отсортированном диапазоне [begin, end),
используя пользовательский компаратор.

### [Min](algorithms/minmax.go)
```go
func Min[T cmp.Ordered](
    begin interfaces.ValueIterator[T],
    end interfaces.Iterator,
) interfaces.ValueIterator[T]
```
Находит минимальный элемент в диапазоне [begin, end) используя естественный порядок элементов (cmp.Ordered).

### [MinC](algorithms/minmax.go)
```go
func MinC[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	cmp comparator.Comparator[T],
) interfaces.ValueIterator[T]
```
Находит минимальный элемент в диапазоне [begin, end) используя пользовательский компаратор.

### [Max](algorithms/minmax.go)
```go
func Max[T cmp.Ordered](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
) interfaces.ValueIterator[T]
```
Находит максимальный элемент в диапазоне [begin, end) используя естественный порядок элементов (cmp.Ordered).

### [MaxC](algorithms/minmax.go)
```go
func MaxC[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	cmp comparator.Comparator[T],
) interfaces.ValueIterator[T]
```
Находит максимальный элемент в диапазоне [begin, end) используя пользовательский компаратор.

### [MinMax](algorithms/minmax.go)
```go
func MinMax[T cmp.Ordered](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
) (min, max interfaces.ValueIterator[T])
```
Находит одновременно минимальный и максимальный элементы в диапазоне [begin, end)
используя естественный порядок элементов (cmp.Ordered).

### [MinMaxC](algorithms/minmax.go)
```go
func MinMaxC[T any](
    first interfaces.ValueIterator[T],
    last interfaces.Iterator,
    cmp comparator.Comparator[T],
) (min, max interfaces.ValueIterator[T])
```
Находит одновременно минимальный и максимальный элементы в диапазоне [begin, end)
используя пользовательский компаратор.

### [NextBound](algorithms/next.go)
```go
func NextBound(
    it interfaces.Iterator,
    bound interfaces.Iterator,
)
```
Продвигает итератор it до тех пор, пока он не достигнет итератора bound.

### [FoldLeft](algorithms/fold.go)
```go
func FoldLeft[T any](
    begin interfaces.ValueIterator[T],
    end interfaces.Iterator,
    init T,
    f binaryFoldFunc[T],
) T
```
Выполняет левую свёртку последовательности, начиная с инициализирующего значения init
и применяя функцию f последовательно ко всем элементам от begin до end.

Если итератор begin равен итератору end, функция вернёт значение init.

### [FoldLeftFirst](algorithms/fold.go)
```go
func FoldLeftFirst[T any](
    begin interfaces.ValueIterator[T],
    end interfaces.Iterator,
    f binaryFoldFunc[T],
) *T
```
Выполняет левую свёртку последовательности, используя в качестве начального значения
первый элемент последовательности. Если последовательность пуста, возвращается nil.

### [FoldRight](algorithms/fold.go)
```go
func FoldRight[T any](
    begin interfaces.BidirectionalIterator[T],
    end interfaces.Iterator,
    init T,
    f binaryFoldFunc[T],
) T
```
Выполняет правую свёртку последовательности, начиная с инициализирующего значения init
и применяя функцию f последовательно ко всем элементам от end до begin в обратном порядке.

Если итератор begin равен итератору end, функция вернёт значение init.

### [FoldRightLast](algorithms/fold.go)
```go
func FoldRightLast[T any](
    begin interfaces.BidirectionalIterator[T],
    end interfaces.Iterator,
    f binaryFoldFunc[T],
) *T
```
Выполняет правую свёртку последовательности, используя в качестве начального значения
последний элемент последовательности. Если последовательность пуста, возвращается nil.

### [TransformUnary](algorithms/transform.go)
```go
func TransformUnary[T any, O any](
    begin interfaces.ValueIterator[T],
    end interfaces.Iterator,
    destBegin interfaces.PointerIterator[O],
    f unaryTransformFunc[T, O],
) interfaces.PointerIterator[O]
```
Применяет унарную функцию f к каждому элементу последовательности,
начиная с итератора begin и заканчивая итератором end, и записывает результаты в
последовательность, начинающуюся с destBegin.

### [TransformBinary](algorithms/transform.go)
```go
func TransformBinary[T1 any, T2 any, O any](
    begin1 interfaces.ValueIterator[T1],
    end1 interfaces.Iterator,
    begin2 interfaces.ValueIterator[T2],
    destBegin interfaces.PointerIterator[O],
    f binaryTransformFunc[T1, T2, O],
) interfaces.PointerIterator[O]
```
Применяет бинарную функцию f к парам элементов из двух последовательностей,
начиная с итераторов begin1 и begin2, и записывает результаты в последовательность, начинающуюся с destBegin.

### [Generate](algorithms/generate.go)
```go
func Generate[T any](
    begin interfaces.PointerIterator[T],
    end interfaces.Iterator,
    g generatorFunc[T],
)
```
Заполняет последовательность значений, начиная с итератора begin и заканчивая итератором end,
значениями, генерируемыми функцией g.

### [GenerateN](algorithms/generate.go)
```go
func GenerateN[T any](
    begin interfaces.PointerIterator[T],
    n uint
    g generatorFunc[T],
)
```
Заполняет n элементов последовательности, начиная с итератора begin,
значениями, генерируемыми функцией g.

### [Replace](algorithms/replace.go)
```go
func Replace[T comparable](
    begin interfaces.PointerIterator[T],
    end interfaces.Iterator,
    oldValue T,
    newValue T,
)
```
Заменяет все вхождения значения oldValue на newValue в диапазоне от begin до end.

### [ReplaceIf](algorithms/replace.go)
```go
func ReplaceIf[T any](
    begin interfaces.PointerIterator[T],
    end interfaces.Iterator,
    predicate unaryPredicate[T],
    newValue T,
)
```
Заменяет все значения, которые удовлетворяют предикату, на newValue
в диапазоне от begin до end.

### [ReplaceCopy](algorithms/replace.go)
```go
func ReplaceCopy[T comparable](
    begin interfaces.ValueIterator[T],
    end interfaces.Iterator,
    destBegin interfaces.PointerIterator[T],
    oldValue T,
    newValue T,
)
```
Копирует элементы из диапазона [begin, end) в destBegin,
заменяя все вхождения oldValue на newValue.

### [ReplaceCopyIf](algorithms/replace.go)
```go
func ReplaceCopyIf[T comparable](
    begin interfaces.ValueIterator[T],
    end interfaces.Iterator,
    destBegin interfaces.PointerIterator[T],
    predicate unaryPredicate[T],
    newValue T,
)
```
Копирует элементы из диапазона [begin, end) в destBegin,
заменяя все значения, которые удовлетворяют предикату, на newValue.

### [Reverse](algorithms/reverse.go)
```go
func Reverse[T any](begin, end interfaces.BidirectionalIterator[T])
```
Разворачивает элементы в последовательности, определяемой итераторами
begin и end, используя двунаправленные итераторы.

### [ReverseCopy](algorithms/reverse.go)
```go
func ReverseCopy[T any](
    begin, end interfaces.BidirectionalIterator[T],
    destBegin interfaces.PointerIterator[T],
) interfaces.PointerIterator[T]
```
Копирует элементы из диапазона [begin, end) в destBegin в обратном порядке.

### [MakeHeap](algorithms/heap.go)
```go
func MakeHeap[T any](
    begin, end interfaces.RandomAccessIterator[T],
    cmp comparator.Comparator[T],
)
```
Преобразует диапазон [begin, end) в кучу, используя предоставленный компаратор cmp.

### [SortHeap](algorithms/heap.go)
```go
func SortHeap[T any](
    begin, end interfaces.RandomAccessIterator[T],
    cmp comparator.Comparator[T],
)
```
Выполняет сортировку кучи на месте, упорядочивая элементы в диапазоне [begin, end).

### [PopHeap](algorithms/heap.go)
```go
func PopHeap[T any](
    begin, end interfaces.RandomAccessIterator[T],
    cmp comparator.Comparator[T],
)
```
Удаляет максимальный элемент из кучи и перестраивает её.

### [PushHeap](algorithms/heap.go)
```go
func PushHeap[T any](
    begin, end interfaces.RandomAccessIterator[T],
    cmp comparator.Comparator[T],
)
```
Добавляет элемент в конец диапазона и перестраивает кучу, чтобы сохранить её свойства.

### [Sort](algorithms/sort.go)
```go
func Sort[T cmp.Ordered](
    begin interfaces.RandomAccessIterator[T],
    end interfaces.RandomAccessIterator[T],
)
```
Выполняет сортировку диапазона [begin, end) с использованием алгоритма pdqsort.

### [SortC](algorithms/sort.go)
```go
func SortC[T any](
    begin interfaces.RandomAccessIterator[T],
    end interfaces.RandomAccessIterator[T],
    cmp comparator.Comparator[T],
)
```
Выполняет сортировку диапазона [begin, end) с использованием алгоритма pdqsort и переданного компаратора.