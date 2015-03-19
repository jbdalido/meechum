package meechum

type Consul struct {
	Host string
}

type ConsulResponse struct {
	key []byte
}

func NewConsul() Backend {
	c := Consul{}

	return c
}

func (c Consul) Connect(host string) error {
	return nil
}

func (c Consul) SetKey(key string, value interface{}) error {
	return nil
}

func (c Consul) SetTtlKey(key string, value interface{}, ttl int) error {
	return nil
}

func (c Consul) GetKey(key string) ([]byte, error) {
	return nil, nil
}

func (c Consul) ListDirectory(dir string) ([]string, error) {
	return nil, nil
}

func (c Consul) DeleteKey(key string) error {
	return nil
}

func (c Consul) UpdateKey(key string, value interface{}) error {
	return nil
}
