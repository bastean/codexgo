package roles

type Logger interface {
	Debug(format string, values ...any)
	Error(format string, values ...any)
	Fatal(format string, values ...any)
	Info(format string, values ...any)
	Success(format string, values ...any)
}
