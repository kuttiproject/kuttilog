// Package kuttilog implements a simple levelled logging system.
// The SetLevel function sets the current level. The Print,
// Printf and Println functions all take a level parameter, and
// only log if it is less than or equal to the current level.
//
// The package provides a Logger interface, so that the actual
// output can vary between implementations. The SetLogger  method
// can be called with an appropriate implementation. There is
// a default implementation, which uses a standard output logger
// created using the Go standard log package.
//
// Levels begin at 0, and can go up to a implemetation-specific
// maximum. The default logger provides five levels: Quiet(0),
// Minimal(1), Info(2), Verbose(3) and Debug(4).
package kuttilog
