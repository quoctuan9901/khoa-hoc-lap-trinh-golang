package models

import (
	"context"
	"sync"
	"time"
)

type Monitor interface {
	Name() string
	Check(ctx context.Context) string
}

type SystemStat struct {
	Name  string
	Value string
}

type ProcStat struct {
	PID         int32
	Name        string
	CPU         float64
	Memory      uint64
	RamPercent  float64
	RunningTime time.Duration
}

var (
	StatMutex sync.Mutex

	Stats = map[string]SystemStat{}
)
