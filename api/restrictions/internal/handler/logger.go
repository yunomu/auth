package handler

type Logger interface {
	Error(error, string, *Request, ...interface{})
}

type defaultLogger struct{}

func (*defaultLogger) Error(_ error, _ string, _ *Request, _ ...interface{}) {}
