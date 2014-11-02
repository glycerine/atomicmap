package amap

import "sync"

type AtomicMap struct {
	m   map[int64]string
	tex sync.RWMutex
}

func NewAtomicMap() *AtomicMap {
	return &AtomicMap{
		m: make(map[int64]string),
	}
}

func (m *AtomicMap) Get(key int64) string {
	m.tex.RLock()
	defer m.tex.RUnlock()
	return m.m[key]
}

func (m *AtomicMap) Get2(key int64) (string, bool) {
	m.tex.RLock()
	defer m.tex.RUnlock()
	v, ok := m.m[key]
	return v, ok
}

func (m *AtomicMap) Set(key int64, val string) {
	m.tex.Lock()
	defer m.tex.Unlock()
	m.m[key] = val
}
