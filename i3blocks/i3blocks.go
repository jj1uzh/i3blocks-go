package i3blocks

type Preamble struct {
	Version     int  `json:"version"`
	ClickEvents bool `json:"click_events"`
}

type BlockObj struct {
	FullText string `json:"full_text"`
}

type BlockUpdate struct {
	I int
	B *BlockObj
}

type Block interface {
	Run(int, chan BlockUpdate)
}
