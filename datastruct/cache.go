package main

import (
	"errors"
	"fmt"
	"sync"
)

0 -> 1 -> 2 -> 3 -> 4
2
// В написании LRU кэша нам поможет структура двусвязного списка.
// Теперь в хэштаблице лежат ноды (узлы) списка, и при переполнении мы откидываем последний элемент,
// а при вставке же или чтении мы двигаем узел с нашим ключом в начало списка.
// Можно было использовать "container/list" из базовой библиотеки go или написать список руками, как сделал я.
type Cache struct {
	storage map[string]*listNode
	mu      sync.RWMutex
	cap     int
	head    *listNode
	tail    *listNode
}

type listNode struct {
	key   string
	value any
	prev  *listNode
	next  *listNode
}

func New(cap int) (*Cache, error) {
	if cap < 1 {
		return nil, errors.New("cap must be >= 1")
	}

	return &Cache{
		storage: make(map[string]*listNode, cap),
		cap:     cap,
	}, nil
}

func (c *Cache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// если ключ уже есть в хранилище, просто обновим значение
	if node, ok := c.storage[key]; ok {
		node.value = value
		// новая запись становится в начало списка
		c.moveToHead(node)

		return
	}

	node := &listNode{key: key, value: value}

	// вытесняет самый холодный кэш при превышении cap
	if len(c.storage) >= c.cap {
		c.removeTail()
	}

	// таковой записи еще не было, добавим в хранилище и список
	c.addToHead(node)
	c.storage[key] = node
}

func (c *Cache) Get(key string) (value any, success bool) {
	// ALERT! помимо доступа в хэш таблице, нужно так же поставить ноду в начало списка: moveToHead
	// поэтому мы не в праве использовать RLock() - это вызовет гонку между двумя конкурирующими
	// вызовами Get
	c.mu.Lock()
	defer c.mu.Unlock()

	if node, ok := c.storage[key]; ok {
		c.moveToHead(node)
		return node.value, true
	}

	return nil, false
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	node, ok := c.storage[key]
	// если записи нет, то и дело с концом
	if !ok {
		return
	}

	// удалим нашу ноду из списка:
	// текущее состояние: node.prev.next -> node <- node.next.prev
	// желаемое:          node.prev.next <-> node.next.prev
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	// чистим из хранилища
	delete(c.storage, key)
}

func (c *Cache) moveToHead(node *listNode) {
	// если нода уже во главе, ничего делать не надо
	if node == c.head || node == nil {
		return
	}

	// текущее состояние: node.prev.next -> node <- node.next.prev
	// желаемое:          node.prev.next <-> node.next.prev
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	// предпоследняя нода становится последней
	if node == c.tail {
		c.tail = node.prev
	}

	// текущая нода становится в начало
	node.prev = nil
	node.next = c.head
	c.head.prev = node
	c.head = node
}

func (c *Cache) addToHead(node *listNode) {
	if node == nil {
		return
	}

	// поставим нашу ноду во главу
	node.prev = nil
	node.next = c.head
	if c.head != nil {
		c.head.prev = node
	}
	c.head = node

	// хвоста нет? значит список пуст, а наша нода стала и головой и хвостом
	if c.tail == nil {
		c.tail = node
	}
}

func (c *Cache) removeTail() {
	if c.tail == nil {
		return
	}

	// отрежем хвост, теперь предпоследняя нода стала хвостовой
	if c.tail.prev != nil {
		c.tail.prev.next = nil
		c.tail = c.tail.prev
	} else {
		// перед хвостом никого нет, значит наша нода и голова и хвост одновременно
		c.head = nil
	}

	delete(c.storage, c.tail.key)
}

func main() {
	mainRace()
}

func mainRace() {
	cache, err := New(3)
	if err != nil {
		panic(err)
	}
	wg := sync.WaitGroup{}

	cache.Set("name", "Alex")
	cache.Set("hobby", "BJJ")
	cache.Set("hobby3", "BJJ")
	cache.Set("hobby2", "BJJ")
	cache.Set("hobby4", "BJJ")
	cache.Set("name", "Alex")

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(cache.Get("name"))
		fmt.Println(cache.Get("hobby"))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cache.Delete("hobby")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(cache.Get("name"))
		fmt.Println(cache.Get("hobby"))
	}()

	wg.Wait()

	fmt.Println(cache)
}
