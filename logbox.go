package main

import (
	"github.com/HimbeerserverDE/mt-multiserver-proxy"
	"github.com/anon55555/mt"

	//	"strings"
	"fmt"
)

const (
	logboxCloseBtn = "button_exit[1.5,7;3,1;quit;Close]"
)

// Logbox is a formspec box for debug logging do a client
// useful e.g. in transitions, or while a complex thing is executed
// use NewLogbox
type Logbox chan<- string

func NewLogbox(cc *proxy.ClientConn, title string) Logbox {
	ch := make(chan string)

	cc.SendCmd(&mt.ToCltShowFormspec{
		Formspec: "size[6,7.5]\nlabel[0,-0.1;" + title + "]\ntextarea[0.3,0.3;5.7,7.2;;;]\n" + logboxCloseBtn,
	})

	// start updating thread
	go func(ch chan string) {
		var content string

		for {
			select {
			case <-ch:
				fmt.Println("ded")
				return

			default:
				str := <-ch
				content += str + "\n"
				cc.SendCmd(&mt.ToCltShowFormspec{
					Formspec: "size[6,7.5]\nlabel[0,-0.1;" + title + "]\ntextarea[0.3,0.3;5.7,7.2;;;" + content + "]\n" + logboxCloseBtn,
				})
				fmt.Println("added", str)
			}
		}
	}(ch)

	return Logbox(ch)
}
