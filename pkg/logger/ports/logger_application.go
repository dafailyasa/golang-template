package ports

type LoggerApplication interface {
	Close() error
	Debug(msg string, args any)
	Info(msg string, args any)
	Warn(msg string, args any)
	Error(msg string, args any)
}
