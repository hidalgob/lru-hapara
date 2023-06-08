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

func (this *LRUCache) Get(key int) int {
	if val, ok := this.cache[key]; ok {
		return val
	}
	return -1
}

func (this *LRUCache) Delete(key int) {
	delete(this.cache, key)
	fmt.Println("deleted")
}

func main() {
	c := Constructor(2)
	c.Put(1, 1)
	fmt.Println(c)
	fmt.Println(c.Get(1))
	c.Delete(1)
	fmt.Println("Exiting")
}
