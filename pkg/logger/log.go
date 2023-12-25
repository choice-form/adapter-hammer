package logger

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Level = zapcore.Level

const (
	InfoLevel   Level = zap.InfoLevel   // 0, default level
	WarnLevel   Level = zap.WarnLevel   // 1
	ErrorLevel  Level = zap.ErrorLevel  // 2
	DPanicLevel Level = zap.DPanicLevel // 3, used in development log
	// PanicLevel logs a message, then panics
	PanicLevel Level = zap.PanicLevel // 4
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel Level = zap.FatalLevel // 5
	DebugLevel Level = zap.DebugLevel // -1
)

var TxTraceIDKey = "X-Txcn-Trace-Id"

func contextValue(ctx context.Context, key string) string {
	value := ctx.Value(TxTraceIDKey)
	if v, ok := value.(string); ok {
		return v
	}
	if v, ok := value.(*string); ok {
		return *v
	}
	return ""
}

type Field = zap.Field

func (l *Logger) Debug(msg string, fields ...Field) {
	l.l.Debug(msg, fields...)
}

func (l *Logger) DebugWithTxTraceID(ctx context.Context, msg string, fields ...Field) {
	traceID := contextValue(ctx, TxTraceIDKey)
	fields = append(fields, String(TxTraceIDKey, traceID))
	l.l.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.l.Info(msg, fields...)
}

func (l *Logger) InfoWithTxTraceID(ctx context.Context, msg string, fields ...Field) {
	traceID := contextValue(ctx, TxTraceIDKey)
	fields = append(fields, String(TxTraceIDKey, traceID))
	l.l.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.l.Warn(msg, fields...)
}

func (l *Logger) WarnWithTxTraceID(ctx context.Context, msg string, fields ...Field) {
	traceID := contextValue(ctx, TxTraceIDKey)
	fields = append(fields, String(TxTraceIDKey, traceID))
	l.l.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.l.Error(msg, fields...)
}

func (l *Logger) ErrorWithTxTraceID(ctx context.Context, msg string, fields ...Field) {
	traceID := contextValue(ctx, TxTraceIDKey)
	fields = append(fields, String(TxTraceIDKey, traceID))
	l.l.Error(msg, fields...)
}

func (l *Logger) DPanic(msg string, fields ...Field) {
	l.l.DPanic(msg, fields...)
}

func (l *Logger) DPanicWithTxTraceID(ctx context.Context, msg string, fields ...Field) {
	traceID := contextValue(ctx, TxTraceIDKey)
	fields = append(fields, String(TxTraceIDKey, traceID))
	l.l.DPanic(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...Field) {
	l.l.Panic(msg, fields...)
}

func (l *Logger) PanicWithTxTraceID(ctx context.Context, msg string, fields ...Field) {
	traceID := contextValue(ctx, TxTraceIDKey)
	fields = append(fields, String(TxTraceIDKey, traceID))
	l.l.Panic(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...Field) {
	l.l.Fatal(msg, fields...)
}

func (l *Logger) FatalWithTxTraceID(ctx context.Context, msg string, fields ...Field) {
	traceID := contextValue(ctx, TxTraceIDKey)
	fields = append(fields, String(TxTraceIDKey, traceID))
	l.l.Fatal(msg, fields...)
}

// function variables for all field types
// in github.com/uber-go/zap/field.go

var (
	Skip        = zap.Skip
	Binary      = zap.Binary
	Bool        = zap.Bool
	Boolp       = zap.Boolp
	ByteString  = zap.ByteString
	Complex128  = zap.Complex128
	Complex128p = zap.Complex128p
	Complex64   = zap.Complex64
	Complex64p  = zap.Complex64p
	Float64     = zap.Float64
	Float64p    = zap.Float64p
	Float32     = zap.Float32
	Float32p    = zap.Float32p
	Int         = zap.Int
	Intp        = zap.Intp
	Int64       = zap.Int64
	Int64p      = zap.Int64p
	Int32       = zap.Int32
	Int32p      = zap.Int32p
	Int16       = zap.Int16
	Int16p      = zap.Int16p
	Int8        = zap.Int8
	Int8p       = zap.Int8p
	String      = zap.String
	Stringp     = zap.Stringp
	Uint        = zap.Uint
	Uintp       = zap.Uintp
	Uint64      = zap.Uint64
	Uint64p     = zap.Uint64p
	Uint32      = zap.Uint32
	Uint32p     = zap.Uint32p
	Uint16      = zap.Uint16
	Uint16p     = zap.Uint16p
	Uint8       = zap.Uint8
	Uint8p      = zap.Uint8p
	Uintptr     = zap.Uintptr
	Uintptrp    = zap.Uintptrp
	Reflect     = zap.Reflect
	Namespace   = zap.Namespace
	Stringer    = zap.Stringer
	Time        = zap.Time
	Timep       = zap.Timep
	Stack       = zap.Stack
	StackSkip   = zap.StackSkip
	Duration    = zap.Duration
	Durationp   = zap.Durationp
	Any         = zap.Any

	Info                = std.Info
	InfoWithTxTraceID   = std.InfoWithTxTraceID
	Warn                = std.Warn
	WarnWithTxTraceID   = std.WarnWithTxTraceID
	Error               = std.Error
	ErrorWithTxTraceID  = std.ErrorWithTxTraceID
	DPanic              = std.DPanic
	DPanicWithTxTraceID = std.DPanicWithTxTraceID
	Panic               = std.Panic
	PanicWithTxTraceID  = std.PanicWithTxTraceID
	Fatal               = std.Fatal
	FatalWithTxTraceID  = std.FatalWithTxTraceID
	Debug               = std.Debug
	DebugWithTxTraceID  = std.DebugWithTxTraceID
)

// not safe for concurrent use
func ResetDefault(l *Logger) {
	std = l
	Info = std.Info
	InfoWithTxTraceID = std.InfoWithTxTraceID
	Warn = std.Warn
	WarnWithTxTraceID = std.WarnWithTxTraceID
	Error = std.Error
	ErrorWithTxTraceID = std.ErrorWithTxTraceID
	DPanic = std.DPanic
	DPanicWithTxTraceID = std.DPanicWithTxTraceID
	Panic = std.Panic
	PanicWithTxTraceID = std.PanicWithTxTraceID
	Fatal = std.Fatal
	FatalWithTxTraceID = std.FatalWithTxTraceID
	Debug = std.Debug
	DebugWithTxTraceID = std.DebugWithTxTraceID
}

type Logger struct {
	l     *zap.Logger // zap ensure that zap.Logger is safe for concurrent use
	level Level
}

var std = New(os.Stderr, InfoLevel, WithCaller(true))

func Default() *Logger {
	return std
}

type Option = zap.Option

var (
	WithCaller    = zap.WithCaller
	AddStacktrace = zap.AddStacktrace
)

type RotateOptions struct {
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
}

type LevelEnablerFunc func(lvl Level) bool

type TeeOption struct {
	Filename string
	Ropt     RotateOptions
	Lef      LevelEnablerFunc
}

func NewTeeWithRotate(tops []TeeOption, opts ...Option) *Logger {
	var cores []zapcore.Core
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02T15:04:05.000Z0700"))
	}

	for _, top := range tops {
		top := top

		lv := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return top.Lef(Level(lvl))
		})

		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   top.Filename,
			MaxSize:    top.Ropt.MaxSize,
			MaxBackups: top.Ropt.MaxBackups,
			MaxAge:     top.Ropt.MaxAge,
			Compress:   top.Ropt.Compress,
		})

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(cfg.EncoderConfig),
			zapcore.AddSync(w),
			lv,
		)
		cores = append(cores, core)
	}

	cores = append(cores, zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	))

	logger := &Logger{
		l: zap.New(zapcore.NewTee(cores...), opts...),
	}
	return logger
}

// New create a new logger (not support log rotating).
func New(writer io.Writer, level Level, opts ...Option) *Logger {
	if writer == nil {
		panic("the writer is nil")
	}
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02T15:04:05.000Z0700"))
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(writer),
		zapcore.Level(level),
	)
	logger := &Logger{
		l:     zap.New(core, opts...),
		level: level,
	}
	return logger
}

func (l *Logger) Sync() error {
	return l.l.Sync()
}

func Sync() error {
	if std != nil {
		return std.Sync()
	}
	return nil
}

func NewWriteLogsFile(logFolder string) {
	err := os.MkdirAll(logFolder, os.ModePerm)
	if err != nil {
		panic(err)
	}
	var tops = []TeeOption{
		{
			// Filename: "logs/access.log",
			Filename: filepath.Join(logFolder, "access.log"),
			Ropt: RotateOptions{
				MaxSize:    8,
				MaxAge:     1,
				MaxBackups: 10,
				Compress:   true,
			},
			Lef: func(lvl Level) bool {
				return lvl <= InfoLevel
			},
		},
		{
			// Filename: "logs/error.log",
			Filename: filepath.Join(logFolder, "error.log"),
			Ropt: RotateOptions{
				MaxSize:    8,
				MaxAge:     1,
				MaxBackups: 10,
				Compress:   true,
			},
			Lef: func(lvl Level) bool {
				return lvl > InfoLevel
			},
		},
	}

	logger := NewTeeWithRotate(tops)
	ResetDefault(logger)
}

const loggerCtxKey = "logger_ctx_key"

// 给 ctx 注入一个 logger, logger 中包含Field(内含日志打印的 k-v对)
func NewCtx(ctx context.Context, fields ...zapcore.Field) context.Context {
	return context.WithValue(ctx, loggerCtxKey, WithCtx(ctx).With(fields...))
}

// 尝试从 context 中获取带有 traceId Field的 logger
func WithCtx(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return std.l
	}
	ctx_logger, ok := ctx.Value(loggerCtxKey).(*zap.Logger)
	if ok {
		return ctx_logger
	}
	return std.l
}
