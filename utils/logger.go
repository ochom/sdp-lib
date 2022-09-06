package utils

import (
	"fmt"
	"log"
	"os"
)

// Logger ...
type Logger struct {
	k *log.Logger
}

// NewLogger ...
func NewLogger() *Logger {
	return &Logger{log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)}
}

// Print ...
func (c *Logger) Print(txt string) {
	_ = c.k.Output(3, txt)
}

// Info ...
func (c *Logger) Info(v ...any) {
	c.Print("INFO: " + fmt.Sprint(v...))
}

// Warn ...
func (c *Logger) Warn(v ...any) {
	c.Print("WARN: " + fmt.Sprint(v...))
}

// Error ...
func (c *Logger) Error(v ...any) {
	c.Print("ERROR: " + fmt.Sprint(v...))
}

// Fatal ...
func (c *Logger) Fatal(v ...any) {
	c.Print("FATAL: " + fmt.Sprint(v...))
	os.Exit(1)
}
