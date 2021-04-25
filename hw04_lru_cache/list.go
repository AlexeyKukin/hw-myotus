package hw04lrucache

// Интерфейс List описывает методы, которые должны быть у структуры.
type List interface {
	Len() int
	Front() *ListItem                  // первый элемент списка			++
	Back() *ListItem                   // последний элемент списка		++
	PushFront(v interface{}) *ListItem // добавить значение в начало	++
	PushBack(v interface{}) *ListItem  // добавить значение в конец		++
	Remove(i *ListItem)                // удалить элемент				++
	MoveToFront(i *ListItem)           // переместить элемент в начало	++
}

// Структура "элемент списка". Из этих элементов будет состоять наш лист.
type ListItem struct {
	Value interface{} // значение в элементе пустой интерфейс, может быть что угодно.
	Next  *ListItem   // ссылка на такой же элемент (следующий)
	Prev  *ListItem   // ссылка на такой же элемент (предыдущий)
}

// Структура лист.
type list struct {
	elemCount int
	firstNode *ListItem
	lastnode  *ListItem
}

// Ко-во элементов списка.
func (l list) Len() int {
	return l.elemCount
}

// Функция возвращающая интерфейс, объявляет новый список при помощи встроенной ф-ии new типа list.
func NewList() List {
	return new(list)
}

// добавить значение в НАЧАЛО списка.
func (l *list) PushFront(v interface{}) *ListItem {
	newNode := new(ListItem)
	if l.elemCount == 0 {
		newNode.Prev, newNode.Next, newNode.Value = nil, nil, v
		l.lastnode, l.firstNode = newNode, newNode
		l.elemCount++
	} else {
		// Формируем новый элемент ноды.
		newNode.Value = v
		newNode.Next = l.firstNode
		newNode.Prev = nil
		// Правим ссылку на следующий элемент ноды у предыдущей ноды.
		l.firstNode.Prev = newNode
		// Меняем указатель в листе на последнюю ноду.
		l.firstNode = newNode
		l.elemCount++
	}
	return newNode
}

// добавить значение в КОНЕЦ списка.
func (l *list) PushBack(v interface{}) *ListItem {
	newNode := new(ListItem)
	if l.elemCount == 0 {
		newNode.Prev, newNode.Next, newNode.Value = nil, nil, v
		l.lastnode, l.firstNode = newNode, newNode
		l.elemCount++
	} else {
		// Формируем новый элемент ноды.
		newNode.Value = v
		newNode.Prev = l.lastnode
		newNode.Next = nil
		// Правим ссылку на следующий элемент ноды у предыдущей ноды.
		l.lastnode.Next = newNode
		// Меняем указатель в листе на последнюю ноду.
		l.lastnode = newNode
		l.elemCount++
	}
	return newNode
}

func (l *list) Front() *ListItem {
	return l.firstNode
}

func (l *list) Back() *ListItem {
	return l.lastnode
}

func (l *list) Remove(i *ListItem) {
	switch {
	case l.elemCount == 1:
		l.firstNode, l.lastnode = nil, nil
	case i.Prev == nil:
		l.firstNode = i.Next
		l.firstNode.Prev = nil
	case i.Next == nil:
		l.lastnode = i.Prev
		l.lastnode.Next = nil
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}

	l.elemCount--
}

func (l *list) MoveToFront(i *ListItem) {
	// Для начала списка. Тут нам ничего делать не нужно, просто выходим.
	if i == l.firstNode {
		return
	}
	// Для конца списка.
	if i == l.lastnode {
		// Сводим соседей. В данном случае у нас сосед со стороны Next nil
		i.Prev.Next = nil
		l.lastnode = i.Prev
		// Меняем ссылки у самого элемента
		i.Prev = nil
		i.Next = l.firstNode

		// Меняем ссылку на пред. значение у первого элемента и устанавливаем значение начала списка.
		l.firstNode.Prev = i
		l.firstNode = i
	} else {
		// Для других элементов списка.
		// Сводим наших соседей друг с другом.
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
		// Меняем ссылки у самого элемента
		i.Prev = nil
		i.Next = l.firstNode
		// Меняем ссылку на пред. значение у первого элемента и устанавливаем значение начала списка.
		l.firstNode.Prev = i
		l.firstNode = i
	}
}
