package monitors

import (
	"context"
	"fmt"

	"github.com/shirou/gopsutil/net"
)

type NetMonitor struct {
}

func (m *NetMonitor) Name() string {
	return "Network"
}

func (m *NetMonitor) Check(ctx context.Context) string {
	netStat, err := net.IOCountersWithContext(ctx, false)
	if err != nil && len(netStat) == 0 {
		return fmt.Sprintf("[Network Monitor] Could not retrieve Network info: %v \n", err)
	}

	value := fmt.Sprintf("Send: %d KB, Recv: %d KB", netStat[0].BytesSent/1024, netStat[0].BytesRecv/1024)

	return value
}
