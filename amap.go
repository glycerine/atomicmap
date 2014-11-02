package amap

import "sync"

type AtomicMapIntString struct {
	m   map[int64]string
	tex sync.RWMutex
}

func NewAtomicMapIntString() *AtomicMapIntString {
	return &AtomicMapIntString{
		m: make(map[int64]string),
	}
}

func (m *AtomicMapIntString) Get(key int64) string {
	m.tex.RLock()
	defer m.tex.RUnlock()
	return m.m[key]
}

func (m *AtomicMapIntString) Get2(key int64) (string, bool) {
	m.tex.RLock()
	defer m.tex.RUnlock()
	v, ok := m.m[key]
	return v, ok
}

func (m *AtomicMapIntString) Set(key int64, val string) {
	m.tex.Lock()
	defer m.tex.Unlock()
	m.m[key] = val
}
