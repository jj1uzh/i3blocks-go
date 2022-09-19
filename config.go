package main

import (
	"i3blocks/block"
	i3b "i3blocks/i3blocks"
)

var Blocks = []i3b.Block{
	block.Clock{Format: "2006-01-02(Mon) 15:04"},
}
