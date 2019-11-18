package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

/**
第一笔hash：
16f0eb42cb4d9c2374b2cb1de4008162c06fdd8f1c18357f0c849eb423672f5f
大小端转换为：
5f2f6723b49e840c7f35181c8fdd6fc0628100e41dcbb274239c4dcb42ebf016

第二笔hash：
cce2f95fc282b3f2bc956f61d6924f73d658a1fdbc71027dd40b06c15822e061
大小端转换为：
61e02258c1060bd47d0271bcfda158d6734f92d6616f95bcf2b382c25ff9e2cc

将两个拼接在一起：
5f2f6723b49e840c7f35181c8fdd6fc0628100e41dcbb274239c4dcb42ebf01661e02258c1060bd47d0271bcfda158d6734f92d6616f95bcf2b382c25ff9e2cc

将上面拼接的字符串进行两次hash如下：

第一次hash结果：
9b2ec096d49fee8b310752082d63d8fe198386ae2172d90533d9186bb28df63d

将上面计算出的hash值再次进行hash：
525894ddd0891b36c5ff8658e2a978d615b35ce6dedb5cb83f2420dbcd40a0c7

大小端转换即为结果：
c7a040cddb20243fb85cdbdee65cb315d678a9e25886ffc5361b89d0dd945852
*/

func ReverseBytes2(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func main() {

	//字符串hash转换为字节
	hash1, _ := hex.DecodeString("16f0eb42cb4d9c2374b2cb1de4008162c06fdd8f1c18357f0c849eb423672f5f")

	hash2, _ := hex.DecodeString("cce2f95fc282b3f2bc956f61d6924f73d658a1fdbc71027dd40b06c15822e061")

	//大小端的转换
	ReverseBytes2(hash1)

	ReverseBytes2(hash2)

	//拼接在一起
	rawdata := append(hash1, hash2...)
	//double hash256
	firsthash := sha256.Sum256(rawdata)
	secondhash := sha256.Sum256(firsthash[:])
	merkroot := secondhash[:]

	//反转，与浏览器当中的数据对比
	ReverseBytes2(merkroot)

	fmt.Printf("%x", merkroot)

}
