package main

import (
	"github.com/hajimehoshi/ebiten"
	g "github.com/joshjiang/go-gundam/cmd/go_gundam"
	u "github.com/joshjiang/go-gundam/cmd/go_gundam/util"
	"log"
)

func main() {
	game := &g.Game{}
	update := game.Update

	if err := ebiten.Run(update, u.ScreenWidth, u.ScreenHeight, 1, "Hello world!"); err != nil {
		log.Fatal(err)
	}
}
