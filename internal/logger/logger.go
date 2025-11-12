package logger

import (
	"fmt"
	"log"
	"strings"
)

type Logger struct {
	env string
}

func New(env string) *Logger { return &Logger{env: env} }

func (l *Logger) formatKV(kv ...any) string {
	if len(kv) == 0 {
		return ""
	}
	parts := make([]string, 0, len(kv)/2)
	for i := 0; i < len(kv)-1; i += 2 {
		parts = append(parts, fmt.Sprintf("%v=%v", kv[i], kv[i+1]))
	}
	return strings.Join(parts, " ")
}

func (l *Logger) Infow(msg string, kv ...any) {
	log.Printf("[INFO] %s %s\n", msg, l.formatKV(kv...))
}

func (l *Logger) Errorw(msg string, kv ...any) {
	log.Printf("[ERROR] %s %s\n", msg, l.formatKV(kv...))
}

func (l *Logger) Fatalw(msg string, kv ...any) {
	log.Fatalf("[FATAL] %s %s\n", msg, l.formatKV(kv...))
}
