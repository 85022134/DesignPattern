package singleton

//使用sync.Once保证一次

import (
	"sync"
)

var single *singleton

type singleton struct {
	name string
}

func (s *singleton) SetName(name string) {
	s.name = name
}

func (s *singleton) GetName() string {
	return s.name
}

var once sync.Once

func GetSingleton() *singleton {
	once.Do(func() {
		single = &singleton{name: "initName"}
	})
	return single
}
