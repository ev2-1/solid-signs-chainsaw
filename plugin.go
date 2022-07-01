package main

import (
	proxy "github.com/HimbeerserverDE/mt-multiserver-proxy"

	"github.com/ev2-1/mt-multiserver-sign-templates"
	"github.com/ev2-1/mt-multiserver-signs"
	"github.com/anon55555/mt"
	"io"
	"os"
	"time"
	"fmt"
	"bufio"
)

var thatSign = [3]int16{2, 10, -5}
var thatOtherSign = [3]int16{4, 10, -5}

var signsStrings = []string{"mcl_signs:wall_sign", "mcl_signs:standing_sign22_5", "mcl_signs:standing_sign45", "mcl_signs:standing_sign67_5"}

// AOIDs held by plugin
var AOIDs = make(map[string]map[mt.AOID]bool)

func init() {
	go signs.LoadCharMap() // download charmap
	signs.Maths()

	template := signTemplates.GetTemplate(signTemplates.ServerPlayerCnt)
	err, signT := template(signs.SignPos{
		Pos:  [3]int16{0, 10, 0},
		Wall: true,
		Rotation: signs.North,
		Server: "hub",
	}, "otherserver")
	if err != nil {
		fmt.Println("[ERROR], cant create template thing", err)
	}
	signs.RegisterSign(signT)

	// get online players every second
	go func() {
		for {
			//		updateSigns()
			time.Sleep(time.Second * 10)
		}
	}()

	f, err := os.Open("./p.l")
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(0)
	}

	for _, sign := range signs.ParseScanner(bufio.NewScanner(f)) {
		fmt.Println("sign", sign.Text)
		signs.RegisterSign(sign)
	}

	/*signs.RegisterSign(&signs.Sign{
		Pos: &signs.SignPos{
			Pos:    [3]int16{2, 10, -5},
			Server: "hub",
			Wall: true,
			Rotation: signs.South,
		},
		Text: "Other Server\n|    [%s/10]   |",
		Color: "black",
		Dyn: []signs.DynContent{
			&signs.Padding{
				Prepend: true,
				Length:  2,
				Filler:  '0',
				Content: &signs.PlayerCnt{
					Srv: "otherserver",
				},
			},
		},
	})*/

	proxy.RegisterChatCmd(proxy.ChatCmd{
		Name:        "signstick",
		Perm:        "default",
		Help:        "no",
		Usage:       "no",
		TelnetUsage: "nada",
		Handler: func(cc *proxy.ClientConn, _ io.Writer, args ...string) string {
			signs.Update()

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
			signs.Ready(cc)
			
			return ""
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

	/*
	onDig := func(cc *proxy.ClientConn, cmd *mt.ToSrvInteract) bool {
		cc.Log("<-", fmt.Sprintf("action %s", cmd.Action.String()))
	
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
	*/
}
