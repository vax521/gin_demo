package main

import (
	"errors"
	"fmt"
)

type LRUCache struct {
	capacity int
	data     []int
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{capacity: cap, data: make([]int, cap)}
}

func (s *LRUCache) AddHead(dataItem int) error {
	tempSlice := make([]int, 1)
	tempSlice = append(tempSlice, dataItem)
	if s.contains(dataItem) {
		if s.data[0] != dataItem {
			s.moveItemToHead(dataItem)
		}
		return nil
	} else {
		if len(s.data) < s.capacity {
			s.data = append(tempSlice, s.data...)
			return nil
		} else {
			return errors.New("out of index")
		}

	}
}

func (s *LRUCache) moveItemToHead(dataItem int) []int {
	itemIndex := s.indexOf(dataItem)
	sliceTemp := s.removeItemByIndex(itemIndex)
	tempSlice := make([]int, 1)
	tempSlice = append(tempSlice, dataItem)
	return append(tempSlice, sliceTemp...)
}

func (s *LRUCache) indexOf(dataItem int) int {
	for index, value := range s.data {
		if value == dataItem {
			return index
		} else {
			return -1
		}
	}
	return -1
}

//删除函数
func (s *LRUCache) removeItemByIndex(i int) []int {
	return append(s.data[:len(s.data)-1], s.data[len(s.data):]...)
}

func (s *LRUCache) contains(dataItem int) bool {
	if s.data == nil || len(s.data) == 0 {
		return false
	} else {
		for item := range s.data {
			if item == dataItem {
				return true
			}
		}
		return false
	}
}

func main() {
	lruCache := NewLRUCache(4)
	inputItem := []int{1, 2, 3, 4, 5, 5, 5, 4, 3, 2, 1, 1, 1, 3, 1}

	for _, input := range inputItem {
		lruCache.AddHead(input)
		fmt.Println(input)
		fmt.Println(lruCache.data)
	}

}
