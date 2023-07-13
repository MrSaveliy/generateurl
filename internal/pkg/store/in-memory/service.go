package in_memory

import (
	"errors"
	"sync"
)

var (
	m = &memory{
		mut:  &sync.RWMutex{},
		data: map[string]string{},
	}
	ErrShortNotExist = errors.New("short not exist")
)

func GetStore() *memory {
	return m
}

type memory struct {
	mut  *sync.RWMutex
	data map[string]string
	l    int64
}

func (m *memory) Get(short string) (string, error) {
	m.mut.RLock()
	defer m.mut.RUnlock()
	v, ok := m.data[short]
	if !ok {
		return "", ErrShortNotExist
	}
	return v, nil
}

func (m *memory) Set(url, short string) error {
	m.mut.Lock()
	defer m.mut.Unlock()
	m.data[short] = url
	m.l++
	return nil
}

func (m *memory) GetLastID() (int64, error) {
	return m.l + 1, nil
}
