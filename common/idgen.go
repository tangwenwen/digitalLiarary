package common

import (
	"sync"
	"time"
)

const (
	Book int64 = iota
	User
)

const (
	nodeBits  uint8 = 6
	stepBits  uint8 = 6
	nodeMax   int64 = -1 ^ (-1 << nodeBits)
	stepMax   int64 = -1 ^ (-1 << stepBits)
	timeShift uint8 = nodeBits + stepBits
	nodeShift uint8 = stepBits
)

type Node struct {
	mu        sync.Mutex
	timestamp int64
	workId    int64
	number    int64
}

var Epoch int64 = 1564588800000

func NewNode(id int64) *Node {
	return &Node{
		timestamp: 0,
		workId:    id,
		number:    0,
	}
}
func GetId(id int64) int64 {
	node := NewNode(id)
	return node.Generate()
}
func (n *Node) Generate() int64 {
	n.mu.Lock()
	defer n.mu.Unlock()

	now := time.Now().UnixNano() / 1e6
	if n.timestamp == now {
		n.number++
		if n.number > stepMax {
			for now <= n.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}

	} else {
		n.number = 0
		n.timestamp = now
	}
	result := int64((now-Epoch)<<timeShift | (n.workId << nodeShift) | (n.number))
	return result
}
