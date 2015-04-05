package meechum

import (
	"fmt"
)

type Backend interface {
	SetKey(key string, value []byte) error
	SetTtlKey(key string, value []byte, ttl int) error
	GetKey(key string) ([]byte, error)
	ListDirectory(dir string) ([]string, error)
	DeleteKey(key string) error
	UpdateKey(key string, value []byte) error
}

func NewBackend(b string, host string) (Backend, error) {
	if b == "" {
		return nil, fmt.Errorf("Only Consul or Etcd are available")
	}
	switch b {
	case "consul":
		c, err := NewConsul(host)
		if err != nil {
			return nil, err
		}
		return c, nil
	case "etcd":
		c, err := NewEtcd(host)
		if err != nil {
			return nil, err
		}
		return c, nil
	}
	return nil, nil
}
