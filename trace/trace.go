package trace

import (
	"fmt"
	"log/slog"
	"runtime"
)

func Trace() {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("for some reason runtime.Caller failed")
	}

	slog.Info(fmt.Sprintf("%s:%v", file, line))
}
