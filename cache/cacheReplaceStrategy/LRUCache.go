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

/**
  核心思想：最近被访问的将来被访问的几率也更大
步骤：
     1. 新数据插入到链表头部；
    2. 每当缓存命中（即缓存数据被访问,需要遍历链表），则将数据移到链表头部；
    3. 当链表满的时候，将链表尾部的数据丢弃。
*/
func (s *LRUCache) AddHead(dataItem int) error {
	tempSlice := []int{dataItem}
	if s.contains(dataItem) {
		if s.data[0] != dataItem {
			s.data = s.moveItemToHead(dataItem)
		}
		return nil
	} else {
		if len(s.data) < s.capacity {
			s.data = append(tempSlice, s.data...)
			return nil
		} else if len(s.data) == s.capacity {
			tempSlice := []int{dataItem}
			s.data = append(tempSlice, s.data[:len(s.data)-1]...)
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
		fmt.Println(lruCache.data)
	}

}
