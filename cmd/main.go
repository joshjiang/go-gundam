package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/joshjiang/go-gundam/cmd/game"
	"github.com/joshjiang/go-gundam/cmd/pieces/bomb"
	"github.com/joshjiang/go-gundam/cmd/player"
	"log"
)

const (
	screenWidth, screenHeight = 1070, 1520
)

var (
	background *ebiten.Image
	gundam     *ebiten.Image
	napalm     *ebiten.Image
	players    []*player.Player
	bombs      []*bomb.Bomb
)

func main() {
	var err error
	gundam, _, err = ebitenutil.NewImageFromFile("../assets/gundam.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	napalm, _, err = ebitenutil.NewImageFromFile("../assets/napalm.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	playerOne := player.New(gundam, screenWidth/2.0, screenHeight/2.0, 4)
	players := append(players, playerOne)
	bombOne := bomb.New(napalm, screenWidth, screenHeight)
	bombs := append(bombs, bombOne)
	game := game.New(players, bombs)

	if err := ebiten.Run(game.Update, screenWidth, screenHeight, .5, "Hello world!"); err != nil {
		log.Fatal(err)
	}
}
