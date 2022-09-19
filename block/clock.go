package block

import (
	i3b "i3blocks/i3blocks"
	"time"
)

type Clock struct {
	Format   string
	Duration time.Duration
}

func (b *Clock) fill() {
	if b.Format == "" {
		b.Format = time.RFC3339
	}
	if b.Duration <= 0 {
		b.Duration = 1 * time.Minute
	}
}

func (b Clock) Run(i int, c chan i3b.BlockUpdate) {
	b.fill()
	tick := time.Tick(b.Duration)
	var o i3b.BlockObj
	u := i3b.BlockUpdate{I: i, B: &o}

	o.FullText = time.Now().Format(b.Format)
	c <- u
	for {
		select {
		case <-tick:
			o.FullText = time.Now().Format(b.Format)
			c <- u
		}
	}
}
