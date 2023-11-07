package logger

import (
	"context"
	"testing"
)

func TestNewLogger(t *testing.T) {
	t.Run("default logger info", func(t *testing.T) {
		Info("test", String("key", "value"))
		Sync()
	})
}

func TestNewWriteLogsFile(t *testing.T) {
	t.Run("logger to file", func(t *testing.T) {
		logsFile := "../../logs"
		NewWriteLogsFile(logsFile)

		Info("log at file", Any("logger", std))
		DPanic("dpanic at file")
		// Panic("panic at file ")
		Error("error at file ", String("error", "test error"))
		Sync()
	})
}

func TestLoggerToFileSlice(t *testing.T) {
	t.Run("TestLoggerToFileSlice", func(t *testing.T) {
		logsFile := "../../logs"
		NewWriteLogsFile(logsFile)

		for i := 0; i < 100000; i++ {
			Info("log at file", Int("i", i))
			Error("log at file", Int("i", i))
		}

		Sync()
	})
}

func TestWithCtx(t *testing.T) {
	t.Run("test withCtx", func(t *testing.T) {

		Error("debug", String("a", "b"))

		ctx := context.Background()
		ctx1 := NewCtx(ctx, String("trance_id", "abcdefgccc"))
		WithCtx(ctx1).Info("with", String("ttt", "bbb"))

		Sync()
	})
}
