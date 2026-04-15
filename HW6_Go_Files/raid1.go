package main

type RAID1 struct {
	disks []Disk
}

func (r *RAID1) Write(blockNum int, data []byte) error {
	for _, d := range r.disks {
		if err := d.WriteBlock(blockNum, data); err != nil {
			return err
		}
	}
	return nil
}

func (r *RAID1) Read(blockNum int) ([]byte, error) {
	return r.disks[0].ReadBlock(blockNum)
}
