package processor

import (
	"context"
	"fmt"
	"os"
	"sort"
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
			value, alert := m.Check(ctx)

			stat := models.SystemStat{
				Name:  m.Name(),
				Value: value,
			}

			statCh <- stat

			if alert {
				LogAlert(stat)
			}
		}
	}
}

func GetTopProcesses(ctx context.Context) string {
	var output string

	vmStat, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return fmt.Sprintf("[Get Top Processes] Could not retrieve mem info: %v \n", err)
	}

	totalMemory := vmStat.Total

	processes, err := process.ProcessesWithContext(ctx)
	if err != nil {
		return fmt.Sprintf("[Get Top Processes] Could not retrieve process list: %v \n", err)
	}

	var wg sync.WaitGroup
	var cpuList, memList []models.ProcStat
	procChan := make(chan models.ProcStat, len(processes))

	for _, p := range processes {
		wg.Add(1)

		go func(proc *process.Process) {
			defer wg.Done()

			select {
			case <-ctx.Done():
				return
			default:
			}

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

			if cpuPercent <= 1 || ramPercent <= 1 {
				return
			}
			
			// createTime đang có dữ liệu là milliseconds
			createTime, err := proc.CreateTimeWithContext(ctx)
			if err != nil {
				return
			}

			// Time Unix là tổng số giây từ mốc thời gian 1970-01-01 00:00:00 UTC đến hiện tại
			// Vì createTime là milliseconds nên cần phải chia 1000 để convert thành second
			runningTime := time.Since(time.Unix(createTime/1000, 0))

			procStat := models.ProcStat{
				PID:         proc.Pid,
				Name:        name,
				CPU:         cpuPercent,
				Memory:      memInfo.RSS,
				RamPercent:  ramPercent,
				RunningTime: runningTime,
			}

			procChan <- procStat
		}(p)
	}

	go func() {
		wg.Wait()
		close(procChan)
	}()

	for stat := range procChan {
		if stat.CPU > 1 {
			cpuList = append(cpuList, stat)
		}

		if stat.RamPercent > 1 {
			memList = append(memList, stat)
		}
	}

	sort.Slice(cpuList, func(i, j int) bool {
		return cpuList[i].CPU > cpuList[j].CPU
	})

	sort.Slice(memList, func(i, j int) bool {
		return memList[i].RamPercent > memList[j].RamPercent
	})

	output += "== Top 5 CPU cosuming processes == \n"
	for k, c := range cpuList[:5] {
		output += fmt.Sprintf("%d. [%d] %s - CPU: %.2f%% - RAM: %.2f MB (%.2f%%) - Running: %s \n",
			k + 1,
			c.PID,
			c.Name,
			c.CPU,
			float64(c.Memory)/1024.0/1024.0,
			c.RamPercent,
			c.RunningTime,
		)
	}

	output += "== Top 5 RAM cosuming processes == \n"
	for k, m := range memList[:5] {
		output += fmt.Sprintf("%d. [%d] %s - CPU: %.2f%% - RAM: %.2f MB (%.2f%%) - Running: %s \n",
			k+1,
			m.PID,
			m.Name,
			m.CPU,
			float64(m.Memory)/1024.0/1024.0,
			m.RamPercent,
			m.RunningTime,
		)
	}

	ExportToCSV(cpuList, memList)

	return output
}

func ExportToCSV(cpuList, memList []models.ProcStat) {
	f, err := os.OpenFile("process_stats.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("[Export To CSV] Failed to write CSV: ", err)
		return
	}
	defer f.Close()

	if stat, err := f.Stat(); err == nil && stat.Size() == 0 {
		f.WriteString("Timestamp,PID,Name,CPU (%),RAM (MB),RAM (%),Running Time \n")
	}

	timestamp := time.Now().Format(time.RFC3339)
	for _, c := range cpuList[:5] {
		line := fmt.Sprintf("%s,%d,%s,%.2f,%.2f,%.2f,%s\n",
			timestamp,
			c.PID,
			c.Name,
			c.CPU,
			float64(c.Memory)/1024.0/1024.0,
			c.RamPercent,
			c.RunningTime,
		)
		f.WriteString(line)
	}

	for _, c := range memList[:5] {
		line := fmt.Sprintf("%s,%d,%s,%.2f,%.2f,%.2f,%s\n",
			timestamp,
			c.PID,
			c.Name,
			c.CPU,
			float64(c.Memory)/1024.0/1024.0,
			c.RamPercent,
			c.RunningTime,
		)
		f.WriteString(line)
	}
}

func LogAlert(stat models.SystemStat) {
	models.StatMutex.Lock()
	defer models.StatMutex.Unlock()

	f, err := os.OpenFile("alert.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("[Log Alert] Failed to write log: ", err)
		return
	}
	defer f.Close()

	timestamp := time.Now().Format(time.RFC3339)
	logLine := fmt.Sprintf("[%s] ALERT: %s = %s \n", timestamp, stat.Name, stat.Value)
	f.WriteString(logLine)
}