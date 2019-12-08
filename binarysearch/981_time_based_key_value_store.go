package binarysearch

import (
	"fmt"
	"sort"
)

// Todo
type TimeMap struct {
	store map[string]TimeMapHelper
}

type TimeMapHelper struct {
	store  map[int]string
	tstore []int
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string]TimeMapHelper),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	store, exists := this.store[key]
	if !exists {
		store = TimeMapHelper{
			store:  make(map[int]string),
			tstore: make([]int, 0),
		}
		this.store[key] = store
	}
	_, exists = store.store[timestamp]
	if !exists {
		store.tstore = append(store.tstore, timestamp)
		sort.Ints(store.tstore)
	}
	store.store[timestamp] = value
	this.store[key] = store
}

func (this *TimeMap) Get(key string, timestamp int) string {
	store, exists := this.store[key]
	if !exists {
		fmt.Println("asadsadas")
		return ""
	}

	v, exists := store.store[timestamp]
	if exists {
		return v
	}
	length := len(store.tstore)
	low, high := 0, length-1
	// find low  bound
	for low <= high {
		mid := low + (high-low)/2

		if store.tstore[mid] >= timestamp {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if low > length-1 {
		low = length - 1
	}
	return store.store[store.tstore[low]]
}
