package main

import "fmt"

func main() {
	disks, err := InitDisks()
	if err != nil {
		fmt.Println("Error initializing disks", err)
		return
	}
	Benchmark(&RAID0{disks}, "RAID 0")
	Benchmark(&RAID1{disks}, "RAID 1")
	Benchmark(&RAID4{disks}, "RAID 4")
	Benchmark(&RAID5{disks}, "RAID 5")
}
