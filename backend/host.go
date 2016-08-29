package backend

import "github.com/berkaroad/saashard/config"

// PhysicalDBPool is data host.
type PhysicalDBPool struct {
	//
}

func NewPhysicalDBPool(hostCfg config.HostConfig) *PhysicalDBPool {
	h := &PhysicalDBPool{}
	return h
}
