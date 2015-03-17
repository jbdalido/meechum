package meechum

type Etcd struct {
	Host string
}

func (e *Etcd) SetKey(key string, value interface{}) error {

}

func (e *Etcd) SetTtlKey(key string, value interface{}, ttl int) error {

}

func (e *Etcd) GetKey(key string) (interface{}, error) {

}

func (e *Etcd) ListDirectory(dir string) (interface{}, error) {

}

func (e *Etcd) DeleteKey(key string) error {

}

func (e *Etcd) UpdateKey(key string, value interface{}) error {

}
