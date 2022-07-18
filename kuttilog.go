package kuttilog

import (
	"log"
	"os"
)

// Default Log levels
const (
	Error   = -1
	Quiet   = 0
	Minimal = 1
	Info    = 2
	Verbose = 3
	Debug   = 4
)

var (
	loglevel      = Info
	logger        Logger
	defaultLogger = &defaultlogger{
		log: log.New(os.Stdout, "", 0),
	}
)

// MaxLevel returns the highest level permitted by the current logger.
func MaxLevel() int {
	return logger.MaxLevel()
}

// LogLevel returns the current log level.
func LogLevel() int {
	return loglevel
}

// SetLogLevel sets the current log level.
// If level is below 0 or greater than MaxLevel, it is not changed.
func SetLogLevel(newlevel int) {
	if newlevel >= 0 && newlevel <= MaxLevel() {
		loglevel = newlevel
	}
}

// V returns true if the specified level is between -1 and the current
// log level, false otherwise.
func V(level int) bool {
	return level <= loglevel && level >= -1
}

func printWith(printfunc func(...interface{}), level int, v ...interface{}) {
	if V(level) {
		prefix := logger.LevelPrefix(level)
		if prefix != "" {
			output := append([]interface{}{logger.LevelPrefix(level)}, v...)
			printfunc(output...)
		} else {
			printfunc(v...)
		}
	}
}

// Print prints to the log, if the level is <= current log level.
// Arguments are handled in the manner of fmt.Print.
func Print(level int, v ...interface{}) {
	printWith(logger.Print, level, v...)
}

// Printf prints to the log, if the level is <= current log level.
// Arguments are handled in the manner of fmt.Printf.
func Printf(level int, format string, v ...interface{}) {
	if V(level) {
		prefix := logger.LevelPrefix(level)
		if prefix != "" {
			output := append([]interface{}{logger.LevelPrefix(level)}, v...)
			logger.Printf("%v "+format, output...)
		} else {
			logger.Printf(format, v...)
		}
	}
}

// Println prints to the log, if the level is <= current log level.
// Arguments are handled in the manner of fmt.Println.
func Println(level int, v ...interface{}) {
	printWith(logger.Println, level, v...)
}

// SetLogger sets the current logger.
func SetLogger(l Logger) {
	logger = l
}

// ResetLogger resets to the default logger.
func ResetLogger() {
	logger = defaultLogger
}

func init() {
	ResetLogger()
}
