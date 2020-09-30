package kuttilog

// Logger contains methods for implemetation-specific logging.
// Exact destination, formatting etc. is implementation-specific.
// Kuttilog will call each print method only if needed. It will
// also append a level prefix string as the first item output.
// Implementations can provide their own prefix strings.
type Logger interface {
	MaxLevel() int                // The maximum level permitted by this logger
	LevelPrefix(level int) string // The logger-supplied prefix string for a level.
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}
