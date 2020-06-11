package glog

import (
	std_log "log"
	"strings"
)

var (
	_ = std_log.Print
)

type (
	Level int
)

const (
	TRACE Level = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

var levelColor = map[Level]string{
	TRACE: "36",
	DEBUG: "32",
	INFO:  "33",
	WARN:  "31",
	ERROR: "31",
	FATAL: "35",
}

func NewLevel(s string) Level {
	s = strings.ToLower(s)
	switch s {
	case "trace":
		return TRACE
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warn":
		return WARN
	case "error":
		return ERROR
	case "fatal":
		return FATAL
	default:
		std_log.Printf("invalid log level '%s'", s)
		return TRACE
	}
}

func (l Level) Color() string {
	return levelColor[l]
}
