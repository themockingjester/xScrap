package machine

import (
	"github.com/shirou/gopsutil/mem"
)

func GetAvailableRAMMB() uint64 {
	v, _ := mem.VirtualMemory()
	return v.Available / (1024 * 1024) // MB
}
