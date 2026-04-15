package main

type RAID5 struct {
	disks []Disk
}

func (r *RAID5) Write(blockNum int, data []byte) error {
	parityDisk := blockNum % NumDisks
	dataDisk := (blockNum + 1) % NumDisks
	block := blockNum / NumDisks
	oldData, _ := r.disks[dataDisk].ReadBlock(block)
	oldParity, _ := r.disks[parityDisk].ReadBlock(block)
	newParity := xor(xor(oldParity, oldData), data)
	if err := r.disks[dataDisk].WriteBlock(block, data); err != nil {
		return err
	}
	return r.disks[parityDisk].WriteBlock(block, newParity)
}

func (r *RAID5) Read(blockNum int) ([]byte, error) {
	dataDisk := (blockNum + 1) % NumDisks
	block := blockNum / NumDisks
	return r.disks[dataDisk].ReadBlock(block)
}
