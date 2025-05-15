package processor

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
	"quoctuan.com/hoc-golang/models"
)

func RunMonitor(ctx context.Context, wg *sync.WaitGroup, statCh chan<- models.SystemStat, m models.Monitor) {
	defer wg.Done()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			statCh <- models.SystemStat{
				Name:  m.Name(),
				Value: m.Check(ctx),
			}
		}
	}
}

func GetTopProcesses(ctx context.Context) string {
	vmStat, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return fmt.Sprintf("[Get Top Processes] Could not retrieve mem info: %v \n", err)
	} 

	totalMemory := vmStat.Total

	processes, err := process.ProcessesWithContext(ctx)
	if err != nil {
		return fmt.Sprintf("[Get Top Processes] Could not retrieve process list: %v \n", err)
	}

	var mu sync.Mutex

	var wg sync.WaitGroup
	for _, p := range processes {
		wg.Add(1)

		go func(proc *process.Process) {
			defer wg.Done()

			select {
			case <-ctx.Done():
				return
			default:
				name, err := proc.NameWithContext(ctx)
				if err != nil {
					return
				}

				cpuPercent, err := proc.CPUPercentWithContext(ctx)
				if err != nil {
					return
				}

				// RSS: Resident Set Size => Ram thực sự mà tiến trình đang sử dụng
				// VMS: Virtual Memory Size => Tổng bộ nhớ ảo mà hệ điều hành cấp phát cho tiến trình
				memInfo, err := proc.MemoryInfoWithContext(ctx)
				if err != nil {
					return
				}

				ramPercent := (float64(memInfo.RSS) / float64(totalMemory)) * 100

				// createTime đang có dữ liệu là milliseconds
				createTime, err := proc.CreateTimeWithContext(ctx)
				if err != nil {
					return
				}

				// Time Unix là tổng số giây từ mốc thời gian 1970-01-01 00:00:00 UTC đến hiện tại
				// Vì createTime là milliseconds nên cần phải chia 1000 để convert thành second
				runningTime := time.Since(time.Unix(createTime/1000, 0))

				if cpuPercent > 5 || ramPercent > 5 {
					mu.Lock()
					procStat := models.ProcStat{
						PID: proc.Pid,
						Name: name,
						CPU: cpuPercent,
						Memory: memInfo.RSS,
						RamPercent: ramPercent,
						RunningTime: runningTime,
					}

					fmt.Printf("%+v \n", procStat)
					mu.Unlock()
				}
			}
		}(p)
	}

	wg.Wait()


	return ""
}
