package ketama

import (
	"crypto/md5"
	"fmt"
	"hash"
	"sort"
)

// server info
type Node struct {
	Ip     string
	Port   string
	Weight int32
}

// virtual ring
var ringMap map[uint32]Node
var sortHashSlice []uint32

func init() {
	ringMap = make(map[uint32]Node)
}

// generate hash digest array
func hashFunc(content []byte) []byte {
	var commonHash hash.Hash = md5.New()
	commonHash.Write(content)
	return commonHash.Sum(nil)
}

// get location node of thw key
func NodeLocation(key string) Node {
	digest := hashFunc([]byte(key))
	keyHash := uint32(digest[0]) |
		uint32(digest[1])<<8 |
		uint32(digest[2])<<16 |
		uint32(digest[3])<<24
	// dichotomy
	highest := len(sortHashSlice)
	lowest := 0
	var matchHash uint32
	for highest > lowest {
		middle := (highest + lowest) / 2
		if keyHash > sortHashSlice[middle-1] && keyHash <= sortHashSlice[middle] {
			matchHash = sortHashSlice[middle]
			break
		}
		if keyHash > sortHashSlice[middle] {
			lowest = middle
		} else {
			highest = middle
		}
	}
	if matchHash == 0 {
		matchHash = sortHashSlice[0]
	}
	return ringMap[matchHash]
}

// build ring, put vitrual node onto ring
func RingBuild(nodeList []Node) {
	var allWeight int32
	var serverNum uint8
	for _, node := range nodeList {
		allWeight += node.Weight
		serverNum++
	}
	for _, node := range nodeList {
		addr := fmt.Sprintf("%s:%s", node.Ip, node.Port)
		// Allocate the number of virtual nodes by weight percentage
		ks := int(float32(node.Weight) / float32(allWeight) * float32(serverNum) * float32(40))

		for i := 0; i < ks; i++ {
			virtualAddr := fmt.Sprintf("%s-%d", addr, i)
			digest := hashFunc([]byte(virtualAddr))
			for i := 0; i < 4; i++ {
				nodeHash := uint32(digest[0+4*i]) |
					uint32(digest[1+4*i])<<8 |
					uint32(digest[2+4*i])<<16 |
					uint32(digest[3+4*i])<<24
				// put it onto ringMap
				ringMap[nodeHash] = node
				sortHashSlice = append(sortHashSlice, nodeHash)
			}
		}

	}
	// sort hash slice
	sort.Slice(sortHashSlice, func(i, j int) bool {
		return sortHashSlice[i] < sortHashSlice[j]
	})
}

// add node, migrate data
func dataMigrate() {

}
