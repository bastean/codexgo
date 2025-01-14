package roles

type Logger interface {
	Debug(string)
	Error(string)
	Fatal(string)
	Info(string)
	Success(string)
}
