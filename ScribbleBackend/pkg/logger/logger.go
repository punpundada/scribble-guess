package logger

import (
	"context"
	"io"
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

func Init() {
	if os.Getenv("ENV") != "production" {
		return
	}
	mw := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "logs/server.log",
		MaxSize:    10, // MB
		MaxBackups: 5,
		MaxAge:     28, // days
		Compress:   true,
	})

	log.SetOutput(mw)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

type ctxKey string

const requestIDKey ctxKey = "request_id"

type Logger struct {
	reqID string
}

func New(reqID string) *Logger {
	return &Logger{reqID: reqID}
}

func FromCtx(ctx context.Context) *Logger {
	if v := ctx.Value(requestIDKey); v != nil {
		if id, ok := v.(string); ok {
			return New(id)
		}
	}
	return New("NO_REQ_ID")
}

func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

func (l *Logger) prefix() string {
	return "[REQ_ID=" + l.reqID + "] "
}

func (l *Logger) Info(v ...any) {
	log.Println(append([]any{l.prefix()}, v...)...)
}

func (l *Logger) Error(v ...any) {
	log.Println(append([]any{l.prefix(), "ERROR:"}, v...)...)
}

func (l *Logger) Printf(format string, v ...any) {
	log.Printf(l.prefix()+format, v...)
}
