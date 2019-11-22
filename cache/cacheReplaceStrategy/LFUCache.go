package main

import "time"

type cacheValueObj struct {
	value    int
	hitCount int
	lastTime time.Time
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

func (lfu *LFUCache) values() []*cacheValueObj {
	var temList []*cacheValueObj
	for _, value := range lfu.cache {
		temList = append(temList, value)
	}
	return temList
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
