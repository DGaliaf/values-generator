package redis

type Storage interface {
	Create(value []byte) (string, []byte, error)
	Get(id string) ([]byte, error)
}
