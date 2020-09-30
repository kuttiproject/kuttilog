// Package kuttilog implements a simple leveled logging system.
// The SetLevel function sets the current level. The Print,
// Printf and Println functions all take a level parameter, and
// only log if it is less than or equal to the current level.
// The package provides a Logger interface, so that the actual
// output can vary between implementations. The SetLogger  method
// can be called with an appropriate implementation. There is also
// a default implementation, which uses the Go standard log package.
// Levels begin at 0, and can go up to a implemetation-specific
// maximum.
package kuttilog
