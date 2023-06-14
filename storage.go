package asinodb

// newStorage instantiates a new database storage
func newStorage() *Storage {
	return &Storage{
		Data: make(map[string]interface{}),
	}
}

type Storage struct {
	Data map[string]interface{}
}

func (s *Storage) get(key string) (interface{}, error) {
	if value, exists := s.Data[key]; exists {
		return value, nil
	}

	return nil, ErrNothing
}

func (s *Storage) set(key string, value interface{}) error {
	s.Data[key] = value
	return nil
}
