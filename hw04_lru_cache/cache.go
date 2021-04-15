package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool // Добавить значение в кэш по ключу.
	Get(key Key) (interface{}, bool)     // Получить значение из кэша по ключу.
	Clear()                              // Очистить кэш.
}

type lruCache struct {
	Cache // Remove me after realization.

	capacity int
	queue    List
	items    map[Key]*ListItem
	mutex    *sync.Mutex
}

type cacheItem struct {
	key   string
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		mutex:    &sync.Mutex{},
	}
}

/* Очередь (queue) будет реализована при помощи двусвязного списка.
Значением будет cacheItem key типа string и значение типа интерфейс].
map (items) будет использоваться для поиска элемента по ключу [ключ типа Key] значение - ссылка на элемент списка.
Емкость кеша (capacity).
*/

// Функция Set добавляет элемент, возвращаемое значение указывает был ли элемент в списке ранее (true) или нет.
func (l *lruCache) Set(key Key, value interface{}) bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	ci := cacheItem{
		string(key),
		value,
	}
	// Проверяем, есть ли ключ в кэше. Если есть - обновляем значение и перемещаем элемент в начало списка.
	_, ok := l.items[key]
	if ok {
		l.queue.MoveToFront(l.items[key])
		l.queue.Front().Value = ci
	} else {
		// А если ключ новый, то проверяем есть ли у нас место. Если места нет освобождаем последний элемент.
		if l.queue.Len() == l.capacity {
			delete(l.items, Key(l.queue.Back().Value.(cacheItem).key))
			l.queue.Remove(l.queue.Back())
		}
		l.items[key] = l.queue.PushFront(ci)
	}
	return ok
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	el, ok := l.items[key]
	if ok {
		l.queue.MoveToFront(l.items[key])
		return el.Value.(cacheItem).value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.items = make(map[Key]*ListItem, l.capacity)
	l.queue = NewList()
}
