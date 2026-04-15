package main

import (
	"fmt"
	"os"
)

const BlockSize = 4096
const NumDisks = 5

type Disk struct {
	file *os.File
}

func (d *Disk) WriteBlock(blockNum int, data []byte) error {
	offset := int64(blockNum * BlockSize)
	_, err := d.file.WriteAt(data, offset)
	if err != nil {
		return err
	}
	return d.file.Sync()
}
func (d *Disk) ReadBlock(blockNum int) ([]byte, error) {
	buf := make([]byte, BlockSize)
	offset := int64(blockNum * BlockSize)
	_, err := d.file.ReadAt(buf, offset)
	return buf, err
}

func InitDisks() ([]Disk, error) {
	disks := make([]Disk, NumDisks)
	for i := 0; i < NumDisks; i++ {
		f, err := os.OpenFile(
			fmt.Sprintf("disk%d.dat", i),
			os.O_CREATE|os.O_RDWR,
			0666,
		)
		if err != nil {
			return nil, err
		}
		disks[i] = Disk{file: f}
	}
	return disks, nil
}
