package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/joshjiang/go-gundam/code/player"
)

type Board struct {
	Players []*player.Player
	Bombs   []*bomb.Bomb
}

func New() *Board {}

func (b *Board) Update(screen *ebiten.Image) error {}
