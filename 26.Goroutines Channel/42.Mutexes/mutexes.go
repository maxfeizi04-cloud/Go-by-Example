package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	cc := Container{
		counters: map[string]int{"a": 1, "b": 2, "c": 3},
	}

	var wg sync.WaitGroup
	doIncrement := func(name string, n int) {
		for range n {
			cc.inc(name)
		}
	}
	wg.Go(func() {
		doIncrement("a", 1000)
	})

	wg.Go(func() {
		doIncrement("a", 1000)
	})

	wg.Go(func() {
		doIncrement("b", 1000)
	})
	wg.Wait()
	fmt.Println(cc.counters)

}
