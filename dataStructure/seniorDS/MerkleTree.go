package main

import (
	"crypto/sha256"
	"fmt"
	"math"
)

type merkleNode struct {
	left  *merkleNode
	right *merkleNode
	data  []byte
}

type merkleTree struct {
	headNode *merkleNode
}

func NewMerkleNode(leftNode, rightNode *merkleNode, data []byte) *merkleNode {
	var node merkleNode
	if leftNode == nil && rightNode == nil {
		hashValue := sha256.Sum256(data)
		node.data = hashValue[:]
	} else {
		preHash := append(leftNode.data, rightNode.data...)
		hashValue := sha256.Sum256(preHash)
		node.data = hashValue[:]
	}
	node.left = leftNode
	node.right = rightNode
	return &node
}

//根据输入的[]byte数组，构造Merkel树
func NewMerkelTree(data [][]byte) *merkleTree {
	var nodes []merkleNode

	//当用叶子节点去生成一颗默克尔树时，必须保证叶子节点的数量为偶数，如果不是需要复制一份最后的交易哈西值到最后拼凑成偶数个交易哈希。
	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}

	for _, dataitem := range data {
		node := NewMerkleNode(nil, nil, dataitem)
		nodes = append(nodes, *node)
	}
	fmt.Println(int(math.Log2(float64(len(nodes)))) + 1)
	for i := 0; i < int(math.Log2(float64(len(nodes))))+1; i++ {
		var newNodes []merkleNode

		for j := 0; j < len(nodes); j += 2 {
			fmt.Printf("%d,%d\n", i, j)
			//防止数组越界
			if j+1 < len(nodes) {
				node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
				newNodes = append(newNodes, *node)
			}
		}
		nodes = newNodes
		fmt.Printf("node.length:%d\n", len(nodes))
	}

	mTree := merkleTree{&nodes[0]}

	return &mTree
}

func main() {
	node1 := NewMerkleNode(nil, nil, []byte("test1"))
	node2 := NewMerkleNode(nil, nil, []byte("test2"))
	node3 := NewMerkleNode(nil, nil, []byte("test3"))
	secondNode1 := NewMerkleNode(node1, node2, []byte(""))
	secconNode2 := NewMerkleNode(node3, node3, []byte(""))
	topNode := NewMerkleNode(secondNode1, secconNode2, []byte(""))
	fmt.Println(topNode.data)
	fmt.Println(len(topNode.data))
	fmt.Printf("%b\n", topNode.data)
	fmt.Println("--------------------------")
	fmt.Println(secondNode1.data)
	testArray := [][]byte{[]byte("test1"), []byte("test2"), []byte("test3"), []byte("test4"), []byte("test5")}
	fmt.Println(testArray)
	merkleTree1 := NewMerkelTree(testArray)
	fmt.Println(merkleTree1.headNode.data)
}
