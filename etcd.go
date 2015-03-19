package meechum

type Etcd struct {
	Host string
}

func NewEtcd() Backend {
	c := Etcd{}

	return c
}

func (e Etcd) Connect(host string) error {
	return nil
}

func (e Etcd) SetKey(key string, value interface{}) error {
	return nil
}

func (e Etcd) SetTtlKey(key string, value interface{}, ttl int) error {
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

func (e Etcd) UpdateKey(key string, value interface{}) error {
	return nil
}
