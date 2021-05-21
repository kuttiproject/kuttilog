package kuttilog

import "log"

var (
	levelprefixes = []string{
		"",
		"",
		"",
		"[Verbose]",
		"[DEBUG]",
	}
)

type defaultlogger struct {
}

func (d defaultlogger) MaxLevel() int {
	return Debug
}

func (d defaultlogger) LevelPrefix(level int) string {
	return levelprefixes[level]
}

func (d defaultlogger) Print(v ...interface{}) {
	log.Println(v...)
}

func (d defaultlogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (d defaultlogger) Println(v ...interface{}) {
	log.Println(v...)
}

func init() {
	log.SetFlags(0)
}
