package formatconverter

type Logger interface {
	Printf(format string, v ...interface{})
}
