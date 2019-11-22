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
	cache map[string]cacheValueObj
}

func (lfu *LFUCache) values() []cacheValueObj {
	var temList []cacheValueObj
	for _, value := range lfu.cache {
		temList = append(temList, value)
	}
	return temList
}
