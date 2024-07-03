package persistence

import "sync"

type Persistence struct {
	store sync.Map
}

func NewPersistence() *Persistence {
	return &Persistence{}
}

func (p *Persistence) Set(key string, value interface{}) {
	p.store.Store(key, value)
}

func (p *Persistence) Get(key string) (interface{}, bool) {
	return p.store.Load(key)
}

func (p *Persistence) Delete(key string) {
	p.store.Delete(key)
}

func (p *Persistence) GetAll() map[string]interface{} {
	m := make(map[string]interface{})
	p.store.Range(func(key, value interface{}) bool {
		m[key.(string)] = value
		return true
	})
	return m
}
