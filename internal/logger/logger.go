package logger

import (
	"log"
	"os"
)

const (
	Reset  = "\033[0m"    //Off-color
	Green  = "\033[1;32m" // Info - 1
	Yellow = "\033[1;33m" // Warn - 2
	Red    = "\033[1;31m" // Error - 3
	White  = "\033[1;37m" // Trace - 4
)

type LogLevel int

const (
	SilentLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	TraceLevel
)

type Interface interface {
	Info(string)
	Warn(string)
	Error(string)
	Trace(string)
}

type Logger struct {
	minLogLevel LogLevel
	log         map[LogLevel]*log.Logger
}

func (l *Logger) write(level LogLevel, message string) {
	if l.minLogLevel <= level {
		l.log[level].Println(message)
	}
}

// Define loggers function

func (l *Logger) Info(message string) {
	l.write(InfoLevel, message)
}

func (l *Logger) Warn(message string) {
	l.write(WarnLevel, message)
}

func (l *Logger) Error(message string) {
	l.write(ErrorLevel, message)
}

func (l *Logger) Trace(message string) {
	l.write(TraceLevel, message)
}

// Init the logger
func NewLogger(level LogLevel) Interface {
	flags := log.Lmsgprefix | log.Ldate | log.Ltime
	return &Logger{
		minLogLevel: level,
		log: map[LogLevel]*log.Logger{
			InfoLevel:  log.New(os.Stdout, Green+"[INFO] "+Reset, flags),
			WarnLevel:  log.New(os.Stdout, Yellow+"[WARN] "+Reset, flags),
			ErrorLevel: log.New(os.Stderr, Red+"[ERROR] "+Reset, flags),
			TraceLevel: log.New(os.Stdout, White+"[TRACE] "+Reset, flags),
		},
	}
}
