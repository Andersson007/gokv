// Package logger provides logger service
package logger

import (
	"fmt"
	"log"
	"os"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
)

type LogEntry struct {
	Level LogLevel
	Msg string
}

var currentLevel = INFO	// Change this to filter logs

func (e LogEntry) Set(level LogLevel, msg string) LogEntry {
	e.Level = level
	e.Msg = msg
	return e
}

func logMsg(entry LogEntry, v ...any) {
	if entry.Level < currentLevel {
		return	// Skip logs below currentLevel
	}
	levelStr := [...]string{"DEBUG", "INFO", "WARNING", "ERROR"}[entry.Level]
	log.Printf("[%s] %s", levelStr, fmt.Sprintf(entry.Msg, v...))
}

func Log(ch chan LogEntry, logFilePath string, logFileFlags int, logFmtFlags int) {
	logFile, err := os.OpenFile(logFilePath, logFileFlags, 0644)
	if err != nil {
		fmt.Fprintln(os.Stdout, "Error opening log file:", err)
		return
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	// TODO Pass these flags from main
	log.SetFlags(logFmtFlags)

	for entry := range ch {
		logMsg(entry)
	}
}
