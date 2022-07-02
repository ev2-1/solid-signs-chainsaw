package main

import (
	proxy "github.com/HimbeerserverDE/mt-multiserver-proxy"

	"github.com/anon55555/mt"

)

func aoid(cc *proxy.ClientConn) mt.AOID {
	_, a := cc.GetFreeAOID()
	if AOIDs[cc.Name()] == nil {
		AOIDs[cc.Name()] = make(map[mt.AOID]bool)
	}

	AOIDs[cc.Name()][a] = true
	return a
}
