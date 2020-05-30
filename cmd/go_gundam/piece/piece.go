package piece

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
)

type piece interface {
	GetPos() (float64, float64)
	GetImage() *ebiten.Image
	Draw() (*ebiten.Image, *ebiten.DrawImageOptions)
}

func printLocation(p piece){
	fmt.Println(p.GetPos())
}