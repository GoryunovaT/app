package storage

import "os"

type Storage struct {
	filename string
}

func NewStorage(filename string) *Storage {
	if filename == "" {
		filename = "calendar.json"
	}
	return &Storage{filename: filename}
}

func (s *Storage) Save(data []byte) error {
	return os.WriteFile(s.filename, data, 0644)
}
func (s *Storage) Load() ([]byte, error) {
	return os.ReadFile(s.filename)
}
