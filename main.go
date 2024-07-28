package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Simple in-memory cache
type Cache struct {
	store map[string]string
	ttl   time.Duration
}

func NewCache(ttl time.Duration) *Cache {
	return &Cache{
		store: make(map[string]string),
		ttl:   ttl,
	}
}

func (c *Cache) Set(key, value string) {
	c.store[key] = value
	time.AfterFunc(c.ttl, func() {
		delete(c.store, key)
	})
}

func (c *Cache) Get(key string) (string, bool) {
	value, exists := c.store[key]
	return value, exists
}

func main() {
	cache := NewCache(5 * time.Minute)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Simple REPL with caching")
	fmt.Println("------------------------")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		if value, exists := cache.Get(input); exists {
			fmt.Println("From cache:", value)
		} else {
			cache.Set(input, input)
			fmt.Println("Stored:", input)
		}
	}
}
