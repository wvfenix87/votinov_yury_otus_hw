package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	// элемент с информацией о первом и последнем элементе.
	mainElem *ListItem
	// mainElem.Prev - первый элемент листа.
	// mainElem.Next - последний элемент листа.
	// длинна списка.
	len int
}

// длинна списка.
func (l *list) Len() int {
	return l.len
}

// возвращение первого элемента.
func (l *list) Front() *ListItem {
	return l.mainElem.Prev
}

// возвращение последнего элемента.
func (l *list) Back() *ListItem {
	return l.mainElem.Next
}

// добавить элемент в начало листа.
func (l *list) PushFront(v interface{}) *ListItem {
	listItem := ListItem{
		Value: v,
		Prev:  nil,
		Next:  l.Front(),
	}
	if l.len == 0 {
		l.mainElem.Next = &listItem
	} else {
		l.mainElem.Prev.Prev = &listItem
	}
	l.mainElem.Prev = &listItem
	l.len++
	return &listItem
}

// добавить элемент в конец списка.
func (l *list) PushBack(v interface{}) *ListItem {
	listItem := ListItem{
		Value: v,
		Prev:  l.Back(),
		Next:  nil,
	}
	if l.len == 0 {
		l.mainElem.Prev = &listItem
	} else {
		l.mainElem.Next.Next = &listItem
	}
	l.mainElem.Next = &listItem
	l.len++
	return &listItem
}

// удаление элемента из листа.
func (l *list) Remove(i *ListItem) {
	prevElem := i.Prev
	nextElem := i.Next
	// проверка является ли элемент единственным.
	isLast := false
	if l.len == 1 && prevElem == nil && nextElem == nil && l.Front() == i && l.Back() == i {
		isLast = true
	}
	// проверка удален ли элемент из массива.
	isDeleted := false
	if prevElem == nil && nextElem == nil && l.Front() != i && l.Back() != i {
		isDeleted = true
	}
	if isDeleted {
		// если элемента нет в массиве ничего не делаем.
		return
	}
	if isLast {
		// если элемент последний просто обнуляем рутовой элемент.
		l.mainElem.Next = nil
		l.mainElem.Prev = nil
		l.len = 0
		i.Prev = nil
		i.Next = nil
		return
	}
	// если элемент последний то переносим рут на предидущий.
	if l.Back() == i {
		l.mainElem.Next = prevElem
		prevElem.Next = nil
	} else {
		// если нет то сшиваем лист.
		nextElem.Prev = prevElem
	}
	// если элемент первый то переносим рут на следующий.
	if l.Front() == i {
		l.mainElem.Prev = nextElem
		nextElem.Prev = nil
	} else {
		// если нет то продолжаем сшивать.
		prevElem.Next = nextElem
	}
	l.len--
	i.Prev = nil
	i.Next = nil
}

// переместить элемент в начало листа.
func (l *list) MoveToFront(i *ListItem) {
	// если в листе 1 элемент ничего не делаем.
	if l.len == 1 {
		return
	}
	// если уже в топе то ничего не делаем.
	if l.Front() == i {
		return
	}
	prevElem := i.Prev
	nextElem := i.Next
	// если элемент последний то переносим рут на предидущий.
	if l.Back() == i {
		l.mainElem.Next = prevElem
	} else {
		// если нет сшиваем лист.
		nextElem.Prev = prevElem
	}
	prevElem.Next = nextElem
	// вставка во фронт.
	i.Next = l.Front()
	i.Prev = nil
	l.mainElem.Prev.Prev = i
	l.mainElem.Prev = i
}

func NewList() List {
	// создание мейн (начального элемента).
	mainItem := ListItem{
		Prev:  nil,
		Next:  nil,
		Value: nil,
	}
	newList := list{
		mainElem: &mainItem,
		len:      0,
	}
	return &newList
}
