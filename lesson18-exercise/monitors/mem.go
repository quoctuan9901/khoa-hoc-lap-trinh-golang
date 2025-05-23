package monitors

import (
	"context"
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

type MemMonitor struct {
}

func (m *MemMonitor) Name() string {
	return "Memory"
}

func (m *MemMonitor) Check(ctx context.Context) (string, bool) {
	vmStat, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return fmt.Sprintf("[Memory Monitor] Could not retrieve Memory info: %v \n", err), false
	}

	value := fmt.Sprintf("%.2f%%", vmStat.UsedPercent)

	return value, vmStat.UsedPercent > 60
}
