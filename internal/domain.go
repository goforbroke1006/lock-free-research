package internal

type NumbersCounter interface {
	Add(n int64) error
	Get() (int64, error)
}
