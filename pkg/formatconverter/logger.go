package formatconverter

// Logger is an logger interface to log messages into output.
type Logger interface {
	Printf(format string, v ...interface{})
}
