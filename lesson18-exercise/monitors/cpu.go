package monitors

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

type CPUMonitor struct {
}

func (m *CPUMonitor) Name() string {
	return "CPU"
}

func (m *CPUMonitor) Check(ctx context.Context) string {
	cpuStat, err := cpu.PercentWithContext(ctx, 1*time.Second, false)
	if err != nil && len(cpuStat) == 0 {
		return fmt.Sprintf("[CPU Monitor] Could not retrieve CPU info: %v \n", err)
	}

	value := fmt.Sprintf("%.2f%%", cpuStat[0])

	return value
}
