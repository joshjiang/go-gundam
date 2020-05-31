package piece

import (
	"github.com/hajimehoshi/ebiten"
)

type Piece interface {
	DrawMap() (*ebiten.Image, *ebiten.DrawImageOptions)
}