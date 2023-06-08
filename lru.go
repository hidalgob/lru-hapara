package main

import "fmt"

type LRUCache struct {
	capacity int
	cache    map[int]int
	keys     []int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]int),
		keys:     make([]int, 0, capacity),
	}
}

func (this *LRUCache) Put(key int, value int) {
	this.cache[key] = value
}

func main() {
	c := Constructor(2)
	fmt.Println(c)
	c.Put(1, 1)
	fmt.Println(c)
	fmt.Println("It works")
}
