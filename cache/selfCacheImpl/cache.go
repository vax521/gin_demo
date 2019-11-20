package main

var minShards = 1024

type cache struct {
	shards []*cacheShard
	hash   fnv64a
}

func newCache() *cache {
	cache := &cache{
		hash:   newDefaultHasher(),
		shards: make([]*cacheShard, minShards),
	}
	for i := 0; i < minShards; i++ {
		cache.shards[i] = initNewShard()
	}

	return cache
}

/**
选取特定 shard 的函数，该函数的实现是首先通过前面的哈希函数将 key 转换成一个 hash 值，
然后用 hash 值与 shard 的数量计算出 shard 数组的下标。
值得一提的是，这里不是用取余运算得到结果，而是通过按位与计算的。
hashedkey&mask
0111
AND 1101 (mask) = 0101
*/
func (c *cache) getShard(hashedKey uint64) (shard *cacheShard) {
	return c.shards[hashedKey&uint64(minShards-1)]
}

func (c *cache) set(key string, value []byte) {
	hashedKey := c.hash.Sum64(key)
	shard := c.getShard(hashedKey)
	shard.set(hashedKey, value)
}

func (c *cache) get(key string) ([]byte, error) {
	hashedKey := c.hash.Sum64(key)
	shard := c.getShard(hashedKey)
	return shard.get(key, hashedKey)
}
