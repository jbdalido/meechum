package meechum

import (
	"os"
)

type Node struct {
	Hostname string
}

func NewNode() *Node {
	n := &Node{}
	host, err := os.Hostname()
	if err != nil {
		n.Hostname = "localhost"
	}
	n.Hostname = host

	return n
}

func (n *Node) String() string {
	return n.Hostname
}
