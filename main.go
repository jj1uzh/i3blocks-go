package main

import (
	"encoding/json"
	"fmt"
	"strings"

	i3b "i3blocks/i3blocks"
)

var (
	preamblev = i3b.Preamble{Version: 1, ClickEvents: true}
)

func rcv(c chan i3b.BlockUpdate) {
	block_jsons := make([]string, len(Blocks))
	for {
		select {
		case u := <-c:
			j, err := json.Marshal(u.B)
			if err != nil {
				block_jsons[u.I] = `{"full_text":"tmp"}`
			} else {
				block_jsons[u.I] = string(j)
			}
			fmt.Printf("[%v],\n", strings.Join(block_jsons, ","))
		}
	}
}

func main() {
	preamblej, err := json.Marshal(preamblev)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(preamblej))
	fmt.Println("[")

	c := make(chan i3b.BlockUpdate)
	go rcv(c)
	for i, b := range Blocks {
		go b.Run(i, c)
	}
	select {}
}
