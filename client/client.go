package client

import (
	"fmt"
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
	Backend *Backend
}

type Node struct {
	Hostname  string
	IPAddress []string
}

func NewClient(b *Backend) (*Client, error) {
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

func (c *Client) EditNode(node string, value *Node) error {

}

func (c *Client) AddNode(node string, value *Node) error {

}

func (c *Client) DeleteNode(node string) error {

}

// Checks functions

func (c *Client) EditCheck(node string, value *Check) error {

}

func (c *Client) AddCheck(node string, value *Check) error {

}

func (c *Client) DeleteCheck(node string, value *Check) error {

}

// Groups functions

func (c *Client) EditGroup(node string, value *Group) error {

}

func (c *Client) AddGroup(node string, value *Group) error {

}

func (c *Client) DeleteGroup(node string, value *Group) error {

}
