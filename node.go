package meechum

type Node struct {
	Hostname string
}

func (n *Node) getHostname() string {
	return "mynode.com"
}
