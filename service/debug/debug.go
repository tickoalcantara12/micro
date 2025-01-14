package debug

import (
	"github.com/tickoalcantara12/micro/v3/service/debug/log"
	memLog "github.com/tickoalcantara12/micro/v3/service/debug/log/memory"
	"github.com/tickoalcantara12/micro/v3/service/debug/profile"
	"github.com/tickoalcantara12/micro/v3/service/debug/stats"
	memStats "github.com/tickoalcantara12/micro/v3/service/debug/stats/memory"
	"github.com/tickoalcantara12/micro/v3/service/debug/trace"
	memTrace "github.com/tickoalcantara12/micro/v3/service/debug/trace/memory"
)

var (
	DefaultLog      log.Log         = memLog.NewLog()
	DefaultTracer   trace.Tracer    = memTrace.NewTracer()
	DefaultStats    stats.Stats     = memStats.NewStats()
	DefaultProfiler profile.Profile = nil
)
