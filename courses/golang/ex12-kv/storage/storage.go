package storage

import (
	"errors"
	"sync"
)

type Storage struct {
	sync.RWMutex
	kv map[string]string
}

func NewStorage() *Storage {
	return &Storage{
		kv: make(map[string]string),
	}
}

// func checkEmpty(str string, fieldName string) error {
// 	if str == "" {
// 		return errors.New(fieldName + " can not be empty")
// 	}
// 	return nil
// }

func (s *Storage) Get(key string) (string, error) {
	s.RLock()
	defer s.RUnlock()

	value := s.kv[key]
	if value == "" {
		return "", errors.New("key '" + key + "' does not exist")
	}

	return value, nil
}

func (s *Storage) Add(key string, value string) error {
	s.Lock()
	defer s.Unlock()

	if key == "" {
		return errors.New("key can not be empty")
	}

	if value == "" {
		return errors.New("value can not be empty")
	}

	if s.kv[key] != "" {
		return errors.New("key '" + key + "' already exists")
	}

	s.kv[key] = value

	return nil
}

func (s *Storage) Delete(key string) (string, error) {
	s.Lock()
	defer s.Unlock()

	if key == "" {
		return "", errors.New("key can not be empty")
	}

	value := s.kv[key]

	if value == "" {
		return "", errors.New("key '" + key + "' does not exist")
	}

	delete(s.kv, key)

	return value, nil
}

func (s *Storage) Update(key string, value string) error {
	s.Lock()
	defer s.Unlock()

	if key == "" {
		return errors.New("key can not be empty")
	}

	if value == "" {
		return errors.New("value can not be empty")
	}

	if s.kv[key] == "" {
		return errors.New("key '" + key + "' does not exist")
	}

	s.kv[key] = value

	return nil
}
