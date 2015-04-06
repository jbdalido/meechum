package meechum

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

type Consul struct {
	Host   string
	Client *api.Client
}

type ConsulResponse struct {
	key []byte
}

func NewConsul(host string) (Consul, error) {

	client, err := api.NewClient(&api.Config{Address: host})
	if err != nil {
		return Consul{}, fmt.Errorf("Consul: %s", err)
	}

	return Consul{
		Client: client,
		Host:   host,
	}, nil
}

func (c Consul) SetKey(key string, value []byte) error {
	kp := &api.KVPair{
		Key:   key,
		Value: value,
	}
	time, err := c.Client.KV().Put(kp, nil)
	if err != nil {
		log.Printf("Consul[%d]: %s", time, err)
		return fmt.Errorf("Consul[%d]: %s", time, err)
	}
	return nil
}

func (c Consul) SetTtlKey(key string, value []byte, ttl int) error {
	kp := &api.KVPair{
		Key:   key,
		Value: []byte(value),
	}
	time, err := c.Client.KV().Put(kp, nil)
	if err != nil {
		return fmt.Errorf("Consul[%d]: %s %d", time.RequestTime, err, ttl)
	}
	return nil
}

func (c Consul) GetKey(key string) ([]byte, error) {
	kp, time, err := c.Client.KV().Get(key, nil)
	if err != nil {
		return nil, fmt.Errorf("Consul[%d]: %s", time.RequestTime, err)
	}
	if kp == nil {
		return nil, fmt.Errorf("Consul[%d] : Key %s does not exist", time.RequestTime, key)
	}
	return kp.Value, nil
}

func (c Consul) ListDirectory(dir string) ([]string, error) {
	keys, time, err := c.Client.KV().Keys(dir, "", nil)
	if err != nil {
		return nil, fmt.Errorf("Consul[%d]: %s", time.RequestTime, err)
	}
	return keys, nil
}

func (c Consul) DeleteKey(key string) error {
	time, err := c.Client.KV().Delete(key, nil)
	if err != nil {
		return fmt.Errorf("Consul[%d]: %s", time.RequestTime, err)
	}
	return nil
}

func (c Consul) UpdateKey(key string, value []byte) error {
	err := c.SetKey(key, value)
	if err != nil {
		return fmt.Errorf("Consul[%d]: %s", err)
	}
	return nil
}
