package meechum

type Consul struct {
	Host string
}

func (c *Consul) SetKey(key string, value interface{}) error {
	return nil
}

func (c *Consul) SetTtlKey(key string, value interface{}, ttl int) error {
	return nil
}

func (c *Consul) GetKey(key string) (interface{}, error) {
	return nil
}

func (c *Consul) ListDirectory(dir string) (interface{}, error) {
	return nil
}

func (c *Consul) DeleteKey(key string) error {
	return nil
}

func (c *Consul) UpdateKey(key string, value interface{}) error {
	return nil
}
