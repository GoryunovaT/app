package storage

import "os"

type JsonStorage struct {
	filename string
}

func NewStorage(filename string) *JsonStorage {
	if filename == "" {
		filename = "calendar.json"
	}
	return &JsonStorage{filename: filename}
}

func (s *JsonStorage) Save(data []byte) error {
	return os.WriteFile(s.filename, data, 0644)
}
func (s *JsonStorage) Load() ([]byte, error) {
	return os.ReadFile(s.filename)
}
