package kuttilog

import (
	"log"
)

var levelprefixes = []string{
	"",
	"",
	"",
	"[Verbose]",
	"[DEBUG]",
}

type defaultlogger struct {
	log *log.Logger
}

func (d *defaultlogger) MaxLevel() int {
	return Debug
}

func (d *defaultlogger) LevelPrefix(level int) string {
	return levelprefixes[level]
}

func (d *defaultlogger) Print(v ...interface{}) {
	d.log.Println(v...)
}

func (d *defaultlogger) Printf(format string, v ...interface{}) {
	d.log.Printf(format, v...)
}

func (d *defaultlogger) Println(v ...interface{}) {
	d.log.Println(v...)
}
