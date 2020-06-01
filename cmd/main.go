package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/joshjiang/go-gundam/pkg"
	u "github.com/joshjiang/go-gundam/pkg/util"
	"log"
)

func main() {
	game := &pkg.Game{}
	update := game.Update

	if err := ebiten.Run(update, u.ScreenWidth, u.ScreenHeight, 1, "Hello world!"); err != nil {
		log.Fatal(err)
	}
}
