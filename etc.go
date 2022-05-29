package main

import (
	proxy "github.com/HimbeerserverDE/mt-multiserver-proxy"
	//	signLib "github.com/ev2-1/mt-multiserver-signs"

	"github.com/anon55555/mt"
	//"strings"
	"sync"
	"time"
)

var onlinePlayers map[string][]*proxy.ClientConn
var onlinePlayersLastUpdate time.Time
var onlinePlayersMu sync.RWMutex

func updateOnlinePlayers() {
	onlinePlayersMu.Lock()
	defer onlinePlayersMu.Unlock()

	onlinePlayers = make(map[string][]*proxy.ClientConn)

	for clt := range proxy.Clts() {
		onlinePlayers[clt.ServerName()] = append(onlinePlayers[clt.ServerName()], clt)
	}

	onlinePlayersLastUpdate = time.Now()
}

func playersOnSrv(srv string) int {
	onlinePlayersMu.RLock()
	defer onlinePlayersMu.RUnlock()

	return len(onlinePlayers[srv])
}

type SignPos struct {
	Pos    [3]int16
	wall   bool
	Server string
}

type Sign struct {
	Pos   *SignPos
	found bool

	Text string
	Dyn  []DynContent

	cachedTexture mt.Texture
}

var signs []*Sign
var signsID = make(map[mt.AOID]*Sign)
var signsMu sync.RWMutex

func RegisterSign(ps *Sign) {
	signsMu.Lock()
	defer signsMu.Unlock()

	signs = append(signs, ps)
}

func updateSigns() {
	/*	sign := signLib.Sign()
		sign.Textures[0] = signLib.GenerateSignTexture("test", true, "black")

		for c := range proxy.Clts() {
			c.SendCmd(&mt.ToCltAORmAdd{
				Add: []mt.AOAdd{
					mt.AOAdd{
						ID: 2564,
						InitData: mt.AOInitData{
							ID:  2564,
							Pos: mt.Pos{1.0, 10.0 - 4.0},

							HP: 0,
							Msgs: []mt.AOMsg{
								&mt.AOCmdProps{
									Props: sign,
								},
							},
						},
					},
				},
			})
		}*/
}

/*func updateSigns() {
	signsMu.Lock()
	for _, sign := range signs {
		sign.cachedTexture = ""

		// resolve dyn content:
		var dyn []any
		for _, d := range sign.Dyn {
			dyn = append(dyn, d.Evaluate(sign.Text, sign.Pos))
		}

		sign.cachedTexture = signLib.GenerateSignTexture(fmt.Sprintf(sign.Text, dyn...), sign.Pos.wall, "black")
	}
	signsMu.Unlock()

	signsMu.RLock()
	defer signsMu.RUnlock()
	for clt := range proxy.Clts() {
		s := clt.ServerName()
		var msgs []mt.IDAOMsg

		for id, sign := range signsID {
			if s == sign.Pos.Server {
				msgs = append(msgs, mt.IDAOMsg{
					ID: id,
					Msg: &mt.AOCmdTextureMod{
						Mod: "^" + sign.cachedTexture,
					},
				})
			}
		}

		if len(msgs) != 0 {
			clt.SendCmd(&mt.ToCltAOMsgs{
				Msgs: msgs,
			})
		}
	}
}
*/

func toIntPos(pos mt.Pos) (p [3]int16) {
	for i := range p {
		p[i] = int16(pos[i] / 10)
	}
	return
}
