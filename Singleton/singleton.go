package singleton

//使用sync.Once保证一次

import (
	"sync"
)

var single *Singleton

type Singleton struct {
	name string
}

func (s *Singleton) SetName(name string) {
	s.name = name
}

func (s *Singleton) GetName() string {
	return s.name
}

var once sync.Once

func GetSingleton() *Singleton {
	once.Do(func() {
		single = &Singleton{name: "initName"}
	})
	return single
}
