package main

import (
	"github.com/hajimehoshi/ebiten"
	//"github.com/hajimehoshi/ebiten/ebitenutil" // This is required to draw debug texts.
	"image/color"
)

// Goal #1: move square around on screen with keyboard
var square *ebiten.Image

func update(screen *ebiten.Image) error{
	if square == nil {
		// create square
		square, _ = ebiten.NewImage(16,16, ebiten.FilterNearest)
	}

	square.Fill(color.White)

	opts := &ebiten.DrawImageOptions{}
	var x float64 = 64
	var y float64 = 64

	opts.GeoM.Translate(x,y)


	screen.DrawImage(square, opts)

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		opts.GeoM.Translate(x, y + 1)
		screen.DrawImage(square, opts)
	}

	return nil
}



func main(){
	if err := ebiten.Run(update, 320, 240, 2, "Hello world!"); err!=nil{
		panic(err)
	}
}