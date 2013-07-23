package main

import (
	"log"
	"runtime"
	"syscall"
	"time"
)

var (
	prevTotalAlloc uint64
	prevNumGC      uint32
	uCpu0          int64
	sCpu0          int64
	resT0          time.Time = time.Now()
)

type MemStats struct {
	AllocKB       int
	AllocKBps int
	GCPerMin        int
	Gortns    int
}

type CPUStats struct {
	Userms    int64
	Userp1000 int
	Sysms     int64
	Sysp1000  int
}

type RsrcStats struct {
	CPU CPUStats
	Mem MemStats
}

func logResourceUsage() {
	t1 := time.Now()
	diff := t1.Sub(resT0).Seconds()

	rs := RsrcStats{}

	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	curTotalAlloc := ms.TotalAlloc
	curNumGC := ms.NumGC

	rs.Mem.AllocKB = int(ms.Alloc / 1024)
	rs.Mem.AllocKBps = int(float64(curTotalAlloc-prevTotalAlloc) / 1024 / diff)
	rs.Mem.GCPerMin = int(float64(curNumGC-prevNumGC) * 60 / diff)
	rs.Mem.Gortns = runtime.NumGoroutine()

	prevTotalAlloc = curTotalAlloc
	prevNumGC = curNumGC

	var ru syscall.Rusage
	syscall.Getrusage(syscall.RUSAGE_SELF, &ru)

	uCpu1 := ru.Utime.Sec*1000 + int64(ru.Utime.Usec)/1000
	sCpu1 := ru.Stime.Sec*1000 + int64(ru.Stime.Usec)/1000
	rs.CPU.Userms = uCpu1 - uCpu0
	rs.CPU.Sysms = sCpu1 - sCpu0
	rs.CPU.Userp1000 = int(float64(rs.CPU.Userms) / diff)
	rs.CPU.Sysp1000 = int(float64(rs.CPU.Sysms) / diff)
	uCpu0 = uCpu1
	sCpu0 = sCpu1

	log.Printf("%#v\n", rs)

	resT0 = t1
}
