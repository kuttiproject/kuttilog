package kuttilog

// Logger contains methods for implementation-specific logging.
// Exact destination, formatting etc. is implementation-specific.
// Kuttilog will call each print method only if needed. It will
// also append a level prefix string as the first item output.
// Implementations can provide their own prefix strings via the
// LevelPrefix method.
type Logger interface {
	// MaxLevel returns the maximum level permitted by this logger.
	MaxLevel() int
	// LevelPrefix returns rhe logger-supplied prefix string for a
	// level. It may return an empty string.
	LevelPrefix(level int) string
	// Print prints to the log.
	// Arguments are handled in the manner of fmt.Print.
	Print(v ...interface{})
	// Printf prints to the log.
	// Arguments are handled in the manner of fmt.Printf.
	Printf(format string, v ...interface{})
	// Println prints to the log.
	// Arguments are handled in the manner of fmt.Println.
	Println(v ...interface{})
}
