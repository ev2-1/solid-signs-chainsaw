package main

import (
	proxy "github.com/HimbeerserverDE/mt-multiserver-proxy"
	signLib "github.com/ev2-1/mt-multiserver-signs"

	"github.com/anon55555/mt"
	"io"
	"time"
	"fmt"
)

var thatSign = [3]int16{2, 10, -5}
var thatOtherSign = [3]int16{4, 10, -5}

var signsStrings = []string{"mcl_signs:wall_sign", "mcl_signs:standing_sign22_5", "mcl_signs:standing_sign45", "mcl_signs:standing_sign67_5"}

// AOIDs held by plugin
var AOIDs = make(map[string]map[mt.AOID]bool)

func init() {
	go signLib.LoadCharMap() // download charmap
	signLib.Maths()

	// get online players every second
	go func() {
		for {
			//		updateSigns()
			time.Sleep(time.Second * 10)
		}
	}()

	RegisterSign(&Sign{
		Pos: &SignPos{
			Pos:    [3]int16{2, 10, -5},
			Server: "hub",
			Wall: true,
			Rotation: signLib.South,
		},
		Text: "Other Server\n|[%s/10]|",
		Color: "black",
		Dyn: []DynContent{
			&Padding{
				Prepend: true,
				Length:  2,
				Filler:  '0',
				Content: &PlayerCnt{
					Srv: "otherserver",
				},
			},
		},
	})

	proxy.RegisterChatCmd(proxy.ChatCmd{
		Name:        "signstick",
		Perm:        "default",
		Help:        "no",
		Usage:       "no",
		TelnetUsage: "nada",
		Handler: func(cc *proxy.ClientConn, _ io.Writer, args ...string) string {
			updateSigns()

			return ""
		},
	})

	proxy.RegisterChatCmd(proxy.ChatCmd{
		Name:        "ready",
		Perm:        "default",
		Help:        "no",
		Usage:       "no",
		TelnetUsage: "nada",
		Handler: func(cc *proxy.ClientConn, _ io.Writer, args ...string) string {
			updateSignText()

			signsMu.RLock()
			defer signsMu.RUnlock()

			add := []mt.AOAdd{}

			for _, s := range signs[cc.ServerName()] {
				add = append(add, signLib.GenerateSignAOAdd(s.cachedText, s.Color, s.Pos.Pos, s.Pos.Rotation, s.Pos.Wall, s.aoid-1))
			}

			if len(add) != 0 {
				cc.SendCmd(&mt.ToCltAORmAdd{
					Add: add,
				})
			}
		
			readyClients[cc.Name()] = true

			return fmt.Sprintf("initialized %d signs", len(add))
		},
	})

	proxy.RegisterChatCmd(proxy.ChatCmd{
		Name:        "rmsigns",
		Perm:        "default",
		Help:        "no",
		Usage:       "no",
		TelnetUsage: "nada",
		Handler: func(cc *proxy.ClientConn, _ io.Writer, args ...string) string {
			if AOIDs[cc.Name()] == nil {
				return "no!"
			}

			rm := []mt.AOID{}
			for k := range AOIDs[cc.Name()] {
				cc.FreeAOID(k)
				delete(AOIDs[cc.Name()], k)
				rm = append(rm, k)
			}

			AOIDs[cc.Name()] = nil

			cc.SendCmd(&mt.ToCltAORmAdd{
				Remove: rm,
			})

			return "aber hallo"
		},
	})

	onDig := func(cc *proxy.ClientConn, cmd *mt.ToSrvInteract) bool {
		switch c := cmd.Pointed.(type) {
		case *mt.PointedNode:
			if c.Under == thatSign {
				cc.Hop("otherserver")
			}
			if c.Under == thatOtherSign {
				cc.Hop("hub")
			}
		}

		return false
	}

	for _, sign := range signsStrings {
		proxy.RegisterNodeHandler(&proxy.NodeHandler{
			Node:  sign,
			OnDig: onDig,
		})
	}
}
