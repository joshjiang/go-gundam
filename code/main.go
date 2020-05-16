package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
	"github.com/hajimehoshi/ebiten/ebitenutil" // This is required to draw debug texts.
)

// Goal #1: move square around on screen with keyboard
const (
	screenWidth, screenHeight = 640, 480
)

var (
	err        error
	background *ebiten.Image
	gundam *ebiten.Image
	playerOne player
)

type player struct {
	image *ebiten.Image
	xPos, yPos float64
	speed	float64
}

func movePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyUp){
		playerOne.yPos -= playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown){
		playerOne.yPos += playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft){
		playerOne.xPos -= playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight){
		playerOne.xPos += playerOne.speed
	}
}
// Run this code once at startup
func init() {
	background, _, err = ebitenutil.NewImageFromFile("../static/yugoslavia.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	gundam, _, err = ebitenutil.NewImageFromFile("../static/Gundam.png", ebiten.FilterDefault)
	if err != nil{
		log.Fatal(err)
	}

	playerOne = player{gundam, screenWidth/2.0, screenHeight/2.0, 4}
}

func update(screen *ebiten.Image) error{
	movePlayer()

	if ebiten.IsDrawingSkipped() {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0,0)
	screen.DrawImage(background,op)

	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(playerOne.xPos,playerOne.yPos)
	screen.DrawImage(playerOne.image, playerOp)

	return nil
}

func main(){
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Hello world!"); err!=nil{
		panic(err)
	}
}