package benchmark

import "runtime"

// ConsumedMem returns consumed memory
func ConsumedMem() uint64 {
	runtime.GC()
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	return s.Sys
}
