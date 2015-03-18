package meechum

type Etcd struct {
	Host string
}

func (e *Etcd) SetKey(key string, value interface{}) error {
	return nil
}

func (e *Etcd) SetTtlKey(key string, value interface{}, ttl int) error {
	return nil
}

func (e *Etcd) GetKey(key string) (interface{}, error) {
	return nil
}

func (e *Etcd) ListDirectory(dir string) (interface{}, error) {
	return nil
}

func (e *Etcd) DeleteKey(key string) error {
	return nil
}

func (e *Etcd) UpdateKey(key string, value interface{}) error {
	return nil
}
