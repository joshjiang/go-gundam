package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/joshjiang/go-gundam/cmd/game"
	"github.com/joshjiang/go-gundam/cmd/piece/bomb"
	"github.com/joshjiang/go-gundam/cmd/piece/player"
	"github.com/joshjiang/go-gundam/cmd/piece/village"
	"log"
)

const (
	screenWidth, screenHeight = 1070, 1520
)

var (
	background *ebiten.Image
	napalm     *ebiten.Image
	players    []*player.Player
	bombs      []*bomb.Bomb
	villages   []*village.Village
)

func main() {
	//TODO Put object create calls in game?
	villageOne := village.New(300, 500)
	villages := append(villages, villageOne)
	playerOne := player.New(screenWidth/2.0, screenHeight/2.0, 4)
	players := append(players, playerOne)
	bombOne := bomb.New(screenWidth, screenHeight)
	bombs := append(bombs, bombOne)

	game := game.New(players, bombs, villages)

	if err := ebiten.Run(game.Update, screenWidth, screenHeight, .5, "Hello world!"); err != nil {
		log.Fatal(err)
	}
}
