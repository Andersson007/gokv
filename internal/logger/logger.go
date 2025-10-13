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
	level LogLevel
	msg string
}

var currentLevel = INFO	// Change this to filter logs

func BuildEntry(level LogLevel, msg string) LogEntry {
	return LogEntry{level: level, msg: msg}
}

func logMsg(entry LogEntry, v ...any) {
	if entry.level < currentLevel {
		return	// Skip logs below currentLevel
	}
	levelStr := [...]string{"DEBUG", "INFO", "WARNING", "ERROR"}[entry.level]
	log.Printf("[%s] %s", levelStr, fmt.Sprintf(entry.msg, v...))
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
