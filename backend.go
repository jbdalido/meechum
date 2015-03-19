package meechum

import (
	"fmt"
)

type Backend interface {
	Connect(host string) error
	SetKey(key string, value interface{}) error
	SetTtlKey(key string, value interface{}, ttl int) error
	GetKey(key string) ([]byte, error)
	ListDirectory(dir string) ([]string, error)
	DeleteKey(key string) error
	UpdateKey(key string, value interface{}) error
}

func NewBackend(b string) (Backend, error) {
	if b == "" {
		return nil, fmt.Errorf("Only Consul or Etcd are available")
	}
	switch b {
	case "consul":
		return NewConsul(), nil
	case "etcd":
		return NewEtcd(), nil
	}
	return NewConsul(), nil
}
