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
	return &LRUCache{capacity: cap, data: make([]int, 0)}
}

func (s *LRUCache) AddHead(dataItem int) error {
	tempSlice := []int{dataItem}
	if s.contains(dataItem) {
		if s.data[0] != dataItem {
			s.data = s.moveItemToHead(dataItem)
			fmt.Println(s.data)
		} else {
			fmt.Println(s.data)
		}
		return nil
	} else {
		if len(s.data) < s.capacity {
			s.data = append(tempSlice, s.data...)
			fmt.Println(s.data)
			return nil
		} else if len(s.data) == s.capacity {
			tempSlice := []int{dataItem}
			s.data = append(tempSlice, s.data[:len(s.data)-1]...)
			fmt.Println(s.data)
			return nil
		} else {
			return errors.New("out of index")
		}

	}
}

func (s *LRUCache) moveItemToHead(dataItem int) []int {
	itemIndex := s.indexOf(dataItem)
	rmSlice := s.removeItemByIndex(itemIndex)
	tempSlice := []int{dataItem}
	return append(tempSlice, rmSlice...)
}

func (s *LRUCache) indexOf(dataItem int) int {
	for index, value := range s.data {
		if value == dataItem {
			return index
		}
	}
	return -1
}

//删除函数
func (s *LRUCache) removeItemByIndex(i int) []int {
	return append(s.data[:i], s.data[i+1:]...)
}

func (s *LRUCache) contains(dataItem int) bool {
	if s.data == nil || len(s.data) == 0 {
		return false
	} else {
		for _, value := range s.data {
			if value == dataItem {
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
		fmt.Println(input)
		lruCache.AddHead(input)
	}

}
