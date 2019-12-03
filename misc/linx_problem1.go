// Copyright (c) 2016 LINE Corporation. All rights reserved.
// LINE Corporation PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.

package main

import (
	"fmt"
)

type LRU struct {
	hashSet     map[int]int
	accessCount map[int]int // TODO: using priority queue
}

func (c *LRU) Add(k, v int) {
	c.hashSet[k] = v

	if _, ok := c.accessCount[k]; !ok {
		c.accessCount[k] = 1
	} else {
		c.accessCount[k]++
	}
}

func (c *LRU) Remove(k int) {
	if _, ok := c.hashSet[k]; !ok {
		fmt.Println(-1)
		return
	}

	delete(c.accessCount, k)
	delete(c.hashSet, k)
}

func (c *LRU) Get(k int) int {
	if v, ok := c.hashSet[k]; !ok {
		fmt.Println(-1)
		return -1
	} else {
		c.accessCount[k]++
		fmt.Println(v)
		return v
	}
}

func (c *LRU) Evict() {
	if len(c.accessCount) == 0 {
		return
	}

	var min int
	var keyToRemove int
	isSet := false

	for k, v := range c.accessCount {
		if !isSet || min > v {
			fmt.Printf("set key to remove: %d, min: %d, v: %d\n", k, min, v)
			min, keyToRemove, isSet = v, k, true
		}
	}

	c.Remove(keyToRemove)
}

// solveProblem read input and online solve problem
func solveProblem(lru *LRU) {
	var k, v int
	command := ""

	finalOutput := -1
	for {
		fmt.Scan(&command)
		switch command {
		case "exit":
			if finalOutput != -1 {
				fmt.Print(finalOutput)
			}
			return
		case "add":
			fmt.Scan(&k, &v)
			lru.Add(k, v)
		case "evict":
			lru.Evict()
		case "get":
			fmt.Scan(&k)

			if output := lru.Get(k); output != -1 {
				finalOutput = output
			}
		case "remove":
			fmt.Scan(&k)
			lru.Remove(k)
		}
	}
}

func main() {
	solveProblem(&LRU{hashSet: make(map[int]int), accessCount: make(map[int]int)})
}

