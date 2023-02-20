package main

import (
	"hash/fnv"
)

func Hash(s []byte) uint32 {
	h := fnv.New32a()
	_, _ = h.Write(s[:])
	return h.Sum32()
}
