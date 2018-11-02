package store

type Store interface {
	Get(filename string) ([]byte, error)
	Put(filename string, buf []byte) error
	Remove(filename string) error
}
