package main

import (
	"log"
)

func main() {
	lruCache := NewLRUCache(3)
	lruCache.Put("b", 2)
	lruCache.Put("a", 1)
	lruCache.Put("c", 3)

	value, _ := lruCache.Get("a")
	if value != 1 {
		log.Fatalf("Require: %d, Got: %d", 1, value)
	}

	lruCache.Put("d", 4)
	_, found := lruCache.Get("b")
	if found {
		log.Fatal("Require: false, Got: true")
	}

	lruCache.Put("a", 5)
	value, _ = lruCache.Get("a")
	if value != 5 {
		log.Fatalf("Require: %d, Got: %d", 5, value)
	}

	log.Println("Passed")
}
