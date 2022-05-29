package main

import (
	proxy "github.com/HimbeerserverDE/mt-multiserver-proxy"
	signLib "github.com/ev2-1/mt-multiserver-signs"

	"fmt"
	"github.com/anon55555/mt"
	"io"
	"strings"
	"time"
)

var thatSign = [3]int16{2, 10, -5}
var thatOtherSign = [3]int16{4, 10, -5}

var signsStrings = []string{"mcl_signs:wall_sign", "mcl_signs:standing_sign22_5", "mcl_signs:standing_sign45", "mcl_signs:standing_sign67_5"}
var blocklist = make(map[mt.AOID]bool) // AOs not to send clt

var AOIDs = make(map[string]map[mt.AOID]bool)

func log(cc *proxy.ClientConn, cmd *mt.ToSrvInteract) bool {
	cc.Log("<-", "interaction", cmd.Action.String())

	switch pos := cmd.Pointed.(type) {
	case *mt.PointedNode:
		cc.Log(" ^", "Node", pos.Under[0], pos.Under[1], pos.Under[2])
	}

	return false
}

func init() {
	go signLib.LoadCharMap() // download charmap
	signLib.Maths()

	// get online players every second
	go func() {
		for {
			updateOnlinePlayers()
			//		updateSigns()
			time.Sleep(time.Second * 10)
		}
	}()

	proxy.RegisterChatCmd(proxy.ChatCmd{
		Name:        "addsigns",
		Perm:        "default",
		Help:        "no",
		Usage:       "no",
		TelnetUsage: "nada",
		Handler: func(cc *proxy.ClientConn, _ io.Writer, args ...string) string {
			text := "%d test\nd\ntest\nddddddddddddddd"

			aoid := func() mt.AOID {
				_, a := cc.GetFreeAOID()
				if AOIDs[cc.Name()] == nil {
					AOIDs[cc.Name()] = make(map[mt.AOID]bool)
				}
				AOIDs[cc.Name()][a] = true
				return a
			}

			a := []mt.AOAdd{}
			for i := 0; i <= int(signLib.West67_5); i++ {
				a = append(a, signLib.GenerateSignAOAdd(fmt.Sprintf(text, i), "black", [3]int16{-2, 9, int16(-4) - int16(i)}, signLib.Rotate(i), false, aoid()))
			}

			a = append(a, signLib.GenerateSignAOAdd(fmt.Sprintf(text, 1), "black", [3]int16{-2, 15, -4}, signLib.North, true, aoid()))
			a = append(a, signLib.GenerateSignAOAdd(fmt.Sprintf(text, 2), "black", [3]int16{-2, 18, -4}, signLib.East,  true, aoid()))
			a = append(a, signLib.GenerateSignAOAdd(fmt.Sprintf(text, 3), "black", [3]int16{-2, 20, -4}, signLib.South, true, aoid()))
			a = append(a, signLib.GenerateSignAOAdd(fmt.Sprintf(text, 4), "black", [3]int16{-2, 22, -4}, signLib.West,  true, aoid()))

			cmd := &mt.ToCltAORmAdd{
				Add: a,
			}
			cc.SendCmd(cmd)

			return "hallo"
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
				AOIDs[cc.Name()][k] = false
				rm = append(rm, k)
			}

			AOIDs[cc.Name()] = nil

			cc.SendCmd(&mt.ToCltAORmAdd{
				Remove: rm,
			})

			return "aber hallo"
		},
	})

	RegisterSign(&Sign{
		Pos: &SignPos{
			Pos:    [3]int16{2, 10, -5},
			Server: "hub",
		},
		Text: "Other Server\n|%s[/10]|",
		Dyn: []DynContent{
			&PlayerCnt{
				Srv: "otherserver",
			},
		},
	})

	/*if p == thatOtherSign { // *that* sign #02
		registerPlayerSign(id, &PlayerSign{
			Server:      "otherserver",
			ServerAbout: "hub",
			Text:        []string{"Back", "[%s/10]"},
			Wall:        wall,
			Type:        0,
		})
	}*/

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

	proxy.RegisterAOHandler(&proxy.AOHandler{
		OnAOAdd: func(cc *proxy.ClientConn, id mt.AOID, add *mt.AOAdd) bool {

			var sign, wall bool
			//		loop:
			logsign(cc, id, add, sign, wall)

			pos := add.InitData.Pos
			for _, msg := range add.InitData.Msgs {
				switch m := msg.(type) {
				case *mt.AOCmdPos:
					pos = m.Pos.Pos
				}
			}

			p := toIntPos(pos)
			cc.Log("sign", "at", pos[0], pos[1], pos[2])

			for _, s := range signs {
				if !s.found {
					if s.Pos.Pos == p {
						s.found = true
						s.Pos.wall = wall

						signsID[id] = s
					}
				}
			}

			return false
		},
		OnAOMsg: func(cc *proxy.ClientConn, id mt.AOID, msg mt.AOMsg) bool {
			if blocklist[id] {
				return true
			} else {
				return false
			}
		},
	})
}

func logsign(cc *proxy.ClientConn, id mt.AOID, add *mt.AOAdd, sign, wall bool) {
	for _, msg := range add.InitData.Msgs {
		switch m := msg.(type) {
		case *mt.AOCmdProps:
			if !sign {
				sign = strings.Contains(m.Props.Infotext, "mcl_signs")
			}
			if sign && !wall {
				blocklist[id] = true
				//						return true // DO NOT SEND CLIENT SIGNS!!11!
				cc.Log("Props: (" + uint16String(uint16(id)) + ")\n" + AOString(m.Props))
				wall = strings.Contains(m.Props.Infotext, "wall_sign")
				//						break loop
			}

		case *mt.AOCmdAnim:
			cc.Log("AOCmdAnim", fmt.Sprintf("Frames: %d&%d@speed:%f; blendering at %f noloop %t", m.Anim.Frames[0], m.Anim.Frames[1], m.Anim.Speed, m.Anim.Blend, m.Anim.NoLoop))
		case *mt.AOCmdArmorGroups:
			cc.Log("AOCmdArmorGroups", armorGroupsString(m.Armor))
		case *mt.AOCmdAttach:
			cc.Log("AOCmdAttach", attachString(m.Attach))
		case *mt.AOCmdTextureMod:
			cc.Log("AOCmdTextureMod", m.Mod)

		default:
			cc.Log(fmt.Sprintf("Sth else %T", m))
		}
	}
	cc.Log("---")
}
