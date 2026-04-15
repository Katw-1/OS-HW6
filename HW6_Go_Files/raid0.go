package main

type RAID0 struct {
	disks []Disk
}

func (r *RAID0) Write(blockNum int, data []byte) error {
	d := blockNum % NumDisks
	b := blockNum / NumDisks
	return r.disks[d].WriteBlock(b, data)
}
func (r *RAID0) Read(blockNum int) ([]byte, error) {
	d := blockNum % NumDisks
	b := blockNum / NumDisks
	return r.disks[d].ReadBlock(b)
}
