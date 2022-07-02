package main

import (
	"github.com/HimbeerserverDE/mt-multiserver-proxy"
	"github.com/anon55555/mt"

	"io"
	"strings"
)

var (
	testForm = Form{
		Title: "testformtitle",
		Name:  "testformname",
		Fields: []Field{
			&Passwd{
				Name:  "testpassname",
				Label: "testpasslable",
				//CloseOnEnter: false, // TODO
			},
			&TextField{
				Name:    "testtxtname",
				Label:   "testtxtlable",
				Default: "i like chess",
			},
			&TextArea{
				Name:    "textareaname",
				Label:   "textarealabel",
				Default: "default_value",

				Size: Offset{10, 5},
			},
		},
	}
)

func init() {
	proxy.RegisterPacketHandler(&proxy.PacketHandler{
		CltHandler: func(cc *proxy.ClientConn, pkt *mt.Pkt) bool {

			return false
		},
		SrvHandler: func(sc *proxy.ServerConn, pkt *mt.Pkt) bool {
			switch cmd := pkt.Cmd.(type) {
			case *mt.ToCltShowFormspec:
				sc.Log("->", cmd.Formspec)
			}

			return false
		},
	})

	proxy.RegisterChatCmd(proxy.ChatCmd{
		Name: "showspec",
		Handler: func(cc *proxy.ClientConn, _ io.Writer, args ...string) string {
			fs := strings.ReplaceAll(strings.Join(args[1:], " "), "\\n", "\n")

			for clt := range proxy.Clts() {
				if clt.Name() == args[0] {
					cc.SendCmd(&mt.ToCltShowFormspec{
						Formspec: fs,
						Formname: "testspec",
					})
				}
			}

			return fs
		},
	})

	// help menu but formspec
	proxy.RegisterChatCmd(proxy.ChatCmd{
		Name: "helpspec",
		Help: "show help menu using formspecs",
		Handler: func(cc *proxy.ClientConn, _ io.Writer, args ...string) (ret string) {
			if len(args) == 0 && cc != nil {
				// head
				ret = "size[13,7.5]\n"
				ret += "label[0,-0.1;Available commands: (see also: /help <cmd>)]\n"

				// tablehead
				ret += "tablecolumns[color;tree;text;text]\n"
				// start of table body
				ret += "table[0,0.5;12.8,4.8;list;"
				ret += "#FFF,0,Command,Usage"

				// sort chatcmds by plugin:
				cmds := make(map[string][]proxy.ChatCmd)

				for _, cmd := range proxy.ChatCmds() {
					plugin := cmd.Plugin()

					cmds[plugin] = append(cmds[plugin], cmd)
				}

				// generate list
				for plugin, cmds := range cmds {
					// each plugin name:
					ret += ",#7AF,0," + plugin + ","

					for _, cmd := range cmds {
						ret += ",#7F7,1," + cmd.Name + "," + Escape(cmd.Usage)
					}
				}

				// end of table
				ret += ";0]\n"

				// a box
				ret += "box[0,5.5;12.8,1.5;#000]\n"
				ret += "textarea[0.3,5.5;13.05,1.9;;;for more information, click on any entry in the list u stupid.\nclick twice if you dare to!]\n"

				// a exit
				ret += "button_exit[5,7;3,1;quit;Close (if u to dumb to press <ESC>)]\n"

				cc.SendCmd(&mt.ToCltShowFormspec{
					Formname: "help",
					Formspec: ret,
				})

				return ""
			}

			return "not yet"
		},
	})

	proxy.RegisterChatCmd(proxy.ChatCmd{
		Name: "testform",
		Help: "show the testform",
		Handler: func(cc *proxy.ClientConn, _ io.Writer, args ...string) string {
			cc.SendCmd(testForm.Formspec())

			return ""
		},
	})

	var box Logbox

	proxy.RegisterChatCmd(proxy.ChatCmd{
		Name: "logbox",
		Help: "show help menu using formspecs",
		Handler: func(cc *proxy.ClientConn, _ io.Writer, args ...string) (ret string) {
			box = NewLogbox(cc, "textbox")

			return
		},
	})

	proxy.RegisterChatCmd(proxy.ChatCmd{
		Name: "log2box",
		Help: "show help menu using formspecs",
		Handler: func(cc *proxy.ClientConn, _ io.Writer, args ...string) (ret string) {
			box <- strings.Join(args, " ")

			return
		},
	})
}
