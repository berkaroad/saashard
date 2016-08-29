package backend

import "github.com/berkaroad/saashard/config"

// PhysicalDBNode is data node.
type PhysicalDBNode struct {
}

func NewPhysicalDBNode(nodeCfg config.NodeConfig) *PhysicalDBNode {
	n := &PhysicalDBNode{}
	return n
}
