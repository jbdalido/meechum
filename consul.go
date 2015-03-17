package meechum

type Consul struct {
	Host string
}

func (c *Consul) SetKey(key string, value interface{}) error {

}

func (c *Consul) SetTtlKey(key string, value interface{}, ttl int) error {

}

func (c *Consul) GetKey(key string) (interface{}, error) {

}

func (c *Consul) ListDirectory(dir string) (interface{}, error) {

}

func (c *Consul) DeleteKey(key string) error {

}

func (c *Consul) UpdateKey(key string, value interface{}) error {

}
