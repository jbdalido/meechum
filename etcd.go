package meechum

type Etcd struct {
	Host string
}

func NewEtcd(host string) (Backend, error) {
	c := Etcd{}
	return c, nil
}

func (e Etcd) SetKey(key string, value []byte) error {
	return nil
}

func (e Etcd) SetTtlKey(key string, value []byte, ttl int) error {
	return nil
}

func (e Etcd) GetKey(key string) ([]byte, error) {
	return nil, nil
}

func (e Etcd) ListDirectory(dir string) ([]string, error) {
	return nil, nil
}

func (e Etcd) DeleteKey(key string) error {
	return nil
}

func (e Etcd) UpdateKey(key string, value []byte) error {
	return nil
}
