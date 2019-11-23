package main

import (
	"fmt"
	"log"
	"time"
)

type cacheValueObj struct {
	value    int
	hitCount int
	lastTime time.Time
}

func (c *cacheValueObj) Stirng() string {
	return fmt.Sprintf("value:%d,hitcount:%d,lastTime:%v", c.value, c.hitCount, c.lastTime.Format("15:04:05"))
}
func NewCacheValueObj(val int, hitCo int, lastTime time.Time) *cacheValueObj {
	return &cacheValueObj{value: val, hitCount: hitCo, lastTime: lastTime}
}

/*func(c * cacheValueObj) updateModiTime(modiTime time.Time) {
    c.lastTime = modiTime
}
*/

type LFUCache struct {
	cap   int
	cache map[int]*cacheValueObj
}

func NewLFUCache(cap int) *LFUCache {
	return &LFUCache{cap: cap, cache: make(map[int]*cacheValueObj)}
}

func (lfu *LFUCache) values() []*cacheValueObj {
	var temList []*cacheValueObj
	for _, value := range lfu.cache {
		temList = append(temList, value)
	}
	return temList
}

func (lfu *LFUCache) getValue(key int) int {
	cacheObj := lfu.cache[key]
	if cacheObj != nil {
		return cacheObj.value
	} else {
		return -1
	}
}
func (lfu *LFUCache) remove(key int) {
	cacheObj := lfu.cache[key]
	if cacheObj != nil {
		delete(lfu.cache, key)
		log.Printf("remove key:%d\n", key)
	} else {

	}
}

func (lfu *LFUCache) put(key, value int) {
	cacheObj := lfu.cache[key]
	if cacheObj == nil { //第一次插入
		if len(lfu.cache) == lfu.cap {
			lfu.remove(key)
		}
		//新增
		cacheObj := NewCacheValueObj(value, 1, time.Now())
		lfu.cache[key] = cacheObj
		fmt.Println(cacheObj)
		log.Printf("add %v：%v\n", key, cacheObj)
	} else {
		cacheObj.value = value
		cacheObj.hitCount += 1
		cacheObj.lastTime = time.Now()
		lfu.cache[key] = cacheObj
		log.Printf("update %v:%v\n", key, cacheObj)
	}

}

func (lfu *LFUCache) getEvicKey() int {
	objList := lfu.values()
	decrHitcount := func(c1, c2 *cacheValueObj) bool {
		return c1.hitCount > c2.hitCount
	}
	incrLastTime := func(c1, c2 *cacheValueObj) bool {
		return c1.lastTime.Unix() < c2.lastTime.Unix()
	}
	OrderedBy(decrHitcount, incrLastTime).Sort(objList)
	removedObj := objList[len(objList)-1]
	for key, objCache := range lfu.cache {
		if removedObj.value == objCache.value {
			return key
		}
	}
	return -1
}

func main() {
	lfuCache := NewLFUCache(3)
	lfuCache.put(1, 1)
	lfuCache.put(2, 2)
	lfuCache.getValue(1)
	lfuCache.getValue(2)
}
