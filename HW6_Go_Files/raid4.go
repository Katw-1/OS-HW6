package main

type RAID4 struct {
	disks []Disk
}

func xor(a, b []byte) []byte {
	out := make([]byte, len(a))
	for i := range a {
		out[i] = a[i] ^ b[i]
	}
	return out
}

func (r *RAID4) Write(blockNum int, data []byte) error {
	dataDisk := blockNum % (NumDisks - 1)
	block := blockNum / (NumDisks - 1)
	parityDisk := NumDisks - 1
	oldData, _ := r.disks[dataDisk].ReadBlock(block)
	oldParity, _ := r.disks[parityDisk].ReadBlock(block)
	newParity := xor(xor(oldParity, oldData), data)
	if err := r.disks[dataDisk].WriteBlock(block, data); err != nil {
		return err
	}
	return r.disks[parityDisk].WriteBlock(block, newParity)
}

func (r *RAID4) Read(blockNum int) ([]byte, error) {
	dataDisk := blockNum % (NumDisks - 1)
	block := blockNum / (NumDisks - 1)
	return r.disks[dataDisk].ReadBlock(block)
}
