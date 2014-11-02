package amap

import "sync"

// int -> string

type AtomicMapIntString struct {
	m   map[int]string
	tex sync.RWMutex
}

func NewAtomicMapIntString() *AtomicMapIntString {
	return &AtomicMapIntString{
		m: make(map[int]string),
	}
}

func (m *AtomicMapIntString) Get(key int) string {
	m.tex.RLock()
	defer m.tex.RUnlock()
	return m.m[key]
}

func (m *AtomicMapIntString) Get2(key int) (string, bool) {
	m.tex.RLock()
	defer m.tex.RUnlock()
	v, ok := m.m[key]
	return v, ok
}

func (m *AtomicMapIntString) Set(key int, val string) {
	m.tex.Lock()
	defer m.tex.Unlock()
	m.m[key] = val
}

func (m *AtomicMapIntString) Del(key int) {
	m.tex.Lock()
	defer m.tex.Unlock()
	delete(m.m, key)
}

// int64 -> string

type AtomicMapInt64String struct {
	m   map[int64]string
	tex sync.RWMutex
}

func NewAtomicMapInt64String() *AtomicMapInt64String {
	return &AtomicMapInt64String{
		m: make(map[int64]string),
	}
}

func (m *AtomicMapInt64String) Get(key int64) string {
	m.tex.RLock()
	defer m.tex.RUnlock()
	return m.m[key]
}

func (m *AtomicMapInt64String) Get2(key int64) (string, bool) {
	m.tex.RLock()
	defer m.tex.RUnlock()
	v, ok := m.m[key]
	return v, ok
}

func (m *AtomicMapInt64String) Set(key int64, val string) {
	m.tex.Lock()
	defer m.tex.Unlock()
	m.m[key] = val
}

func (m *AtomicMapInt64String) Del(key int64) {
	m.tex.Lock()
	defer m.tex.Unlock()
	delete(m.m, key)
}

// int64 -> interface{}

type AtomicMapInt64Iface struct {
	m   map[int64]interface{}
	tex sync.RWMutex
}

func NewAtomicMapInt64Iface() *AtomicMapInt64Iface {
	return &AtomicMapInt64Iface{
		m: make(map[int64]interface{}),
	}
}

func (m *AtomicMapInt64Iface) Get(key int64) interface{} {
	m.tex.RLock()
	defer m.tex.RUnlock()
	return m.m[key]
}

func (m *AtomicMapInt64Iface) Get2(key int64) (interface{}, bool) {
	m.tex.RLock()
	defer m.tex.RUnlock()
	v, ok := m.m[key]
	return v, ok
}

func (m *AtomicMapInt64Iface) Set(key int64, val interface{}) {
	m.tex.Lock()
	defer m.tex.Unlock()
	m.m[key] = val
}

func (m *AtomicMapInt64Iface) Del(key int64) {
	m.tex.Lock()
	defer m.tex.Unlock()
	delete(m.m, key)
}

// int -> interface{}

type AtomicMapIntIface struct {
	m   map[int]interface{}
	tex sync.RWMutex
}

func NewAtomicMapIntIface() *AtomicMapIntIface {
	return &AtomicMapIntIface{
		m: make(map[int]interface{}),
	}
}

func (m *AtomicMapIntIface) Get(key int) interface{} {
	m.tex.RLock()
	defer m.tex.RUnlock()
	return m.m[key]
}

func (m *AtomicMapIntIface) Get2(key int) (interface{}, bool) {
	m.tex.RLock()
	defer m.tex.RUnlock()
	v, ok := m.m[key]
	return v, ok
}

func (m *AtomicMapIntIface) Set(key int, val interface{}) {
	m.tex.Lock()
	defer m.tex.Unlock()
	m.m[key] = val
}

func (m *AtomicMapIntIface) Del(key int) {
	m.tex.Lock()
	defer m.tex.Unlock()
	delete(m.m, key)
}
