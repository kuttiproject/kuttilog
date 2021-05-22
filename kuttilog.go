package kuttilog

import (
	"log"
	"os"
)

// Default Log levels
const (
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

// Loglevel returns the current log level.
func Loglevel() int {
	return loglevel
}

// Setloglevel sets the current log level.
// If level is below 0 or greater than MaxLevel, it is not changed.
func Setloglevel(newlevel int) {
	if newlevel >= 0 && newlevel <= MaxLevel() {
		loglevel = newlevel
	}
}

// V returns true if the specified level is between 0 and the current
// log level, false otherwise.
func V(level int) bool {
	return level <= loglevel && level >= 0
}

// Print prints to the log, if the level is <= current log level.
// Arguments are handled in the manner of fmt.Print.
func Print(level int, v ...interface{}) {
	if V(level) {
		output := append([]interface{}{logger.LevelPrefix(level)}, v...)
		logger.Print(output...)
	}
}

// Printf prints to the log, if the level is <= current log level.
// Arguments are handled in the manner of fmt.Printf.
func Printf(level int, format string, v ...interface{}) {
	if V(level) {
		output := append([]interface{}{logger.LevelPrefix(level)}, v...)
		logger.Printf("%v"+format, output...)
	}
}

// Println prints to the log, if the level is <= current log level.
// Arguments are handled in the manner of fmt.Println.
// There will always be a space printed before the first argument,
// after the label prefix if any.
func Println(level int, v ...interface{}) {
	if V(level) {
		output := append([]interface{}{logger.LevelPrefix(level)}, v...)
		logger.Println(output...)
	}
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
