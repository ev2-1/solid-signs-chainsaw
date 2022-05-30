package main

import (
	proxy "github.com/HimbeerserverDE/mt-multiserver-proxy"
	signLib "github.com/ev2-1/mt-multiserver-signs"

	"github.com/anon55555/mt"
	//"strings"
	"sync"
	"fmt"
)

var readyClients = make(map[string]bool)

func aoid(cc *proxy.ClientConn) mt.AOID {
	_, a := cc.GetFreeAOID()
	if AOIDs[cc.Name()] == nil {
		AOIDs[cc.Name()] = make(map[mt.AOID]bool)
	}

	AOIDs[cc.Name()][a] = true
	return a
}

type SignPos struct {
	Pos      [3]int16
	Wall     bool
	Rotation signLib.Rotate
	Server   string
}

type Sign struct {
	Pos   *SignPos
	aoid  mt.AOID

	Text  string
	Color string
	Dyn  []DynContent

	cachedText string
	changed bool
}

var signs = make(map[string][]*Sign)
var signsMu sync.RWMutex

func RegisterSign(ps *Sign) {
	signsMu.Lock()
	defer signsMu.Unlock()

	_, ps.aoid = proxy.GetServerAOId(ps.Pos.Server)

	signs[ps.Pos.Server] = append(signs[ps.Pos.Server], ps)
}

func updateSignText() {
	signsMu.Lock()
	defer signsMu.Unlock()

	for _, srv := range signs {
		for _, s := range srv {
			var dyn []any
			for _, d := range s.Dyn {
				dyn = append(dyn, d.Evaluate(s.Text, s.Pos))
			}	

			text := fmt.Sprintf(s.Text, dyn...)

			s.changed = !(text == s.cachedText)
			if s.changed {
				s.cachedText = text
			}
		}
	}
}

func updateSigns() {
	updateSignText()
	
	signsMu.RLock()
	defer signsMu.RUnlock()

	sendCache := make(map[string][]mt.IDAOMsg)

	for srv, signs := range signs {
		for _, s := range signs {
			if s.changed {
				sendCache[srv] = append(sendCache[srv], mt.IDAOMsg{
					ID: s.aoid,
					Msg: signLib.GenerateTextureAOMod(s.cachedText, s.Pos.Wall, s.Color),
				})
			}
		}
	}

	for clt := range proxy.Clts() {
		if !readyClients[clt.Name()] {
			break
		}
	
		srv := clt.ServerName()
	
		if len(sendCache[srv]) != 0 {
			clt.SendCmd(&mt.ToCltAOMsgs{
				Msgs: sendCache[srv],
			})
		}
	}
}


func toIntPos(pos mt.Pos) (p [3]int16) {
	for i := range p {
		p[i] = int16(pos[i] / 10)
	}
	return
}
