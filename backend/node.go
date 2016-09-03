package backend

import "github.com/berkaroad/saashard/config"

// DataNode is a data node.
type DataNode struct {
	Name     string
	Database string
	DataHost *DataHost
}

// NewDataNode new node instance.
func NewDataNode(nodeCfg config.NodeConfig, dataHost *DataHost) *DataNode {
	n := new(DataNode)
	n.Name = nodeCfg.Name
	n.Database = nodeCfg.Database
	n.DataHost = dataHost
	return n
}
