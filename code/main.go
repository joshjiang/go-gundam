package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil" // This is required to draw debug texts.
	"log"
	"math/rand"
)

// Goal #1: move square around on screen with keyboard
const (
	screenWidth, screenHeight = 1070, 1520
)

var (
	err        error
	background *ebiten.Image
	gundam *ebiten.Image
	napalm *ebiten.Image
	playerOne player
	bombOne bomb
)

type player struct {
	image *ebiten.Image
	xPos, yPos float64
	speed	float64
}

type bomb struct {
	image *ebiten.Image
	xPos, yPos float64
	isPickedUp bool
}

func abs(x float64) float64{
	if x < 0{
		val := x * -1
		return val
	}
	return x
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

func grabBomb(){
	if abs(playerOne.yPos - bombOne.yPos) >= 10 && abs(playerOne.xPos - bombOne.xPos) >= 10{
		bombOne.isPickedUp = true
		for {
			bombOne.xPos, bombOne.yPos = playerOne.xPos, playerOne.yPos + 20
		}
	}
}

// Run this code once at startup
func init() {
	background, _, err = ebitenutil.NewImageFromFile("../static/korea.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	gundam, _, err = ebitenutil.NewImageFromFile("../static/gundam.png", ebiten.FilterDefault)
	if err != nil{
		log.Fatal(err)
	}

	napalm, _, err = ebitenutil.NewImageFromFile("../static/napalm.png", ebiten.FilterDefault)
	if err != nil{
		log.Fatal(err)
	}

	playerOne = player{gundam, screenWidth/2.0, screenHeight/2.0, 4}
	bombOne = bomb{napalm, rand.Float64()*screenWidth,rand.Float64()*screenHeight, false}
}

func update(screen *ebiten.Image) error{
	movePlayer()
	grabBomb()

	if ebiten.IsDrawingSkipped() {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0,0)
	screen.DrawImage(background,op)

	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Scale(2.0,2.0)
	playerOp.GeoM.Translate(playerOne.xPos, playerOne.yPos)
	screen.DrawImage(playerOne.image, playerOp)

	bombOp := &ebiten.DrawImageOptions{}
	bombOp.GeoM.Scale(0.2,0.2)
	bombOp.GeoM.Translate(bombOne.xPos, bombOne.yPos)
	screen.DrawImage(bombOne.image, bombOp)

	return nil
}

func main(){
	if err := ebiten.Run(update, screenWidth, screenHeight, .5, "Hello world!"); err!=nil{
		panic(err)
	}
}