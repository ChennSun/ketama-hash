package ketama

import (
	"crypto/md5"
	"hash"
)

func init() {

}

// generate hash digest array
func HashFunc(content []byte) []byte {
	var commonHash hash.Hash = md5.New()
	commonHash.Write(content)
	return commonHash.Sum(nil)
}

// 获取key所在的物理节点信息
func NodeLocation(key string) {

}

// 构建虚拟环，将物理节点映射为虚拟节点并落在环上
func ringConstruct() {

}

// 新增物理节点
func AddNode() {

}

// 移除物理节点
func RemoveNode() {

}

// 数据rehash
func dataMigrate() {

}
