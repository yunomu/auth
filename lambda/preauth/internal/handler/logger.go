package handler

type Logger interface {
	Error(error, string)
}

type defaultLogger struct{}

func (*defaultLogger) Error(_ error, _ string) {}
