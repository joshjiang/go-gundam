package piece

import (
	"github.com/hajimehoshi/ebiten"
)

type Piece interface {
	DrawMap() (*ebiten.Image, *ebiten.DrawImageOptions)
	DrawBattle() (*ebiten.Image, *ebiten.DrawImageOptions)
	GetPos() (float64,float64)
}