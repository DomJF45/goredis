package main

import (
	"fmt"
)

type Store struct {
	data  map[int]string
	cache Cacher
}

func NewStore(c Cacher) *Store {
	data := map[int]string{
		1: "Elon musk owns twitter",
		2: "Foo is not bar and bar is not baz",
		3: "Git gud",
	}
	return &Store{
		data:  data,
		cache: c,
	}
}

func (s *Store) Get(key int) (string, error) {
	val, ok := s.cache.Get(key)
	if ok {
		// bust the cache
		if err := s.cache.Remove(key); err != nil {
			fmt.Println(err)
		}
		fmt.Println("Returning the value from cache")
		return val, nil
	}
	val, ok = s.data[key]
	if !ok {
		return "", fmt.Errorf("key not found")
	}

	if err := s.cache.Set(key, val); err != nil {
		return "", err
	}

	fmt.Println("returning key from internal storage")

	return val, nil
}
