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

func (c cacheValueObj) String() string {
	return fmt.Sprintf("(%v, %v,%v)", c.value, c.hitCount, c.lastTime.Format("2006/01/02 15:04:05"))
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
		cacheObj.hitCount += 1
		cacheObj.lastTime = time.Now()
		log.Printf("query %v：%v\n", key, cacheObj)
		return cacheObj.value
	} else {
		return -1
	}
}

func (lfu *LFUCache) put(key, value int) {
	cacheObj := lfu.cache[key]
	fmt.Println(cacheObj)
	if cacheObj == nil { //第一次插入
		fmt.Println(len(lfu.cache))
		if len(lfu.cache) == lfu.cap {
			toDelKey := lfu.getEvicKey()
			log.Printf("被淘汰的key:%d\n", toDelKey)
			delete(lfu.cache, toDelKey)
		}
		//新增
		cacheObj := NewCacheValueObj(value, 1, time.Now())
		lfu.cache[key] = cacheObj
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
	//按访问次数从高到底排序
	decrHitcount := func(c1, c2 *cacheValueObj) bool {
		return c1.hitCount > c2.hitCount
	}
	//按时间访问早晚排序，越小时间越早，保留最近被访问的
	incrLastTime := func(c1, c2 *cacheValueObj) bool {
		return c1.lastTime.Unix() > c2.lastTime.Unix()
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
func (lfu *LFUCache) printCacheDetails() {
	for key, cacheValObj := range lfu.cache {
		log.Printf("current details %v:%v\n", key, cacheValObj)
	}
}

func main() {
	lfuCache := NewLFUCache(3)
	lfuCache.put(1, 1)
	time.Sleep(time.Second)
	lfuCache.put(2, 2)
	time.Sleep(time.Second)
	lfuCache.getValue(1)
	time.Sleep(time.Second)
	lfuCache.getValue(2)
	lfuCache.put(3, 3)
	time.Sleep(time.Second)
	lfuCache.getValue(3)
	lfuCache.printCacheDetails()
	time.Sleep(time.Second)
	lfuCache.put(4, 4)
	time.Sleep(time.Second)
	lfuCache.printCacheDetails()
	lfuCache.put(5, 5)
	lfuCache.printCacheDetails()
}
