package meechum

import (
	"fmt"
)

type Backend interface {
	SetKey(key string, value interface{}) error
	SetTtlKey(key string, value interface{}, ttl int) error
	GetKey(key string) (interface{}, error)
	ListDirectory(dir string) (interface{}, error)
	DeleteKey(key string) error
	UpdateKey(key string, value interface{}) error
}

func NewBackend(b string) (Backend, error) {
	if b == "" {
		return nil, fmt.Errorf("Only Consul or Etcd are available")
	}
	switch b {
	case "consul":
		return Consul{}, nil
	case "etcd":
		return Etcd{}, nil
	}
	return Consul{}, nil
}
