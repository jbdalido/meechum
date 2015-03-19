package client

import (
	"fmt"
	"github.com/jbdalido/meechum"
)

type Backend interface {
	Connect() error
	Add(key string, value interface{}) error
	Edit(key string, value interface{}) error
	Delete(key string, value interface{}) error
	Watch(key string) (error, interface{})
}

// HERE IS A CLIENT TO THE API

type Client struct {
	Backend Backend
}

type Node struct {
	Hostname  string
	IPAddress []string
}

func NewClient(b Backend) (*Client, error) {
	if b == nil {
		return nil, fmt.Errorf("You need to select a backend, either etcd of Conseul")
	}

	err := b.Connect()
	if err != nil {
		return nil, err
	}

	return &Client{
		Backend: b,
	}, nil
}

// Node functions

func (c *Client) EditNode(node string, value *meechum.Node) error {
	return nil
}

func (c *Client) AddNode(node string, value *meechum.Node) error {
	return nil
}

func (c *Client) DeleteNode(node string) error {
	return nil
}

// Checks functions

func (c *Client) EditCheck(node string, value *meechum.Check) error {
	return nil
}

func (c *Client) AddCheck(node string, value *meechum.Check) error {
	return nil
}

func (c *Client) DeleteCheck(node string, value *meechum.Check) error {
	return nil
}

// Groups functions

func (c *Client) EditGroup(node string, value *meechum.Group) error {
	return nil
}

func (c *Client) AddGroup(node string, value *meechum.Group) error {
	return nil
}

func (c *Client) DeleteGroup(node string, value *meechum.Group) error {
	return nil
}
