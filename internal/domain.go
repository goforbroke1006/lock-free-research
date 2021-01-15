package internal

type NumbersCounter interface {
	Add(n int64) error
	Get() (int64, error)
}

type TCPListener interface {
	Run()
	OnMessage(func(message string, reply *string))
	Stop() error
}
