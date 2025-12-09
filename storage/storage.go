package storage

// Storage - интерфейс для любого хранилища
type Storage interface {
	Save(data []byte) error
	Load() ([]byte, error)
}
