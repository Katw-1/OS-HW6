package main

import (
	"crypto/rand"
	"fmt"
	"time"
)

func Benchmark(r RAID, name string) {
	totalBlocks := (100 * 1024 * 1024) / BlockSize //100 MB
	data := make([]byte, BlockSize)
	rand.Read(data)
	start := time.Now()
	for i := 0; i < totalBlocks; i++ {
		r.Write(i, data)
	}
	writeTime := time.Since(start)
	start = time.Now()
	for i := 0; i < totalBlocks; i++ {
		r.Read(i)
	}
	readTime := time.Since(start)

	fmt.Printf("=== %s ===\n", name)
	fmt.Printf("Write: %v total, %v per block\n",
		writeTime, writeTime/time.Duration(totalBlocks))
	fmt.Printf("Read: %v total, %v per block\n\n", readTime, readTime/time.Duration(totalBlocks))
}
