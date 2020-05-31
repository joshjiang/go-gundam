package player

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/joshjiang/go-gundam/cmd/go_gundam/piece/village"
	"github.com/joshjiang/go-gundam/cmd/go_gundam/util"
	"log"
	"math"
)

// Player ...
type Player struct {
	currentImage, imageDown, imageUp, imageLeft, imageRight,
	imageUpLeft, imageUpRight, imageDownLeft, imageDownRight, battleImage *ebiten.Image
	xPos, yPos float64
	speed      float64
}

// New ...
func New(xpos, ypos, speed float64) *Player {
	down, _, err := ebitenutil.NewImageFromFile("../assets/gundam/down.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	up, _, err := ebitenutil.NewImageFromFile("../assets/gundam/up.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	left, _, err := ebitenutil.NewImageFromFile("../assets/gundam/left.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	right, _, err := ebitenutil.NewImageFromFile("../assets/gundam/right.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	upLeft, _, err := ebitenutil.NewImageFromFile("../assets/gundam/upLeft.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	upRight, _, err := ebitenutil.NewImageFromFile("../assets/gundam/upRight.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	downLeft, _, err := ebitenutil.NewImageFromFile("../assets/gundam/downLeft.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	downRight, _, err := ebitenutil.NewImageFromFile("../assets/gundam/downRight.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	battleImage, _, err := ebitenutil.NewImageFromFile("../assets/gundam/battle-gundam.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	return &Player{
		imageDown:      down,
		imageUp:        up,
		imageLeft:      left,
		imageRight:     right,
		imageUpLeft:    upLeft,
		imageUpRight:   upRight,
		imageDownLeft:  downLeft,
		imageDownRight: downRight,
		currentImage:   down,
		battleImage:	battleImage,
		xPos:           xpos,
		yPos:           ypos,
		speed:          speed,
	}
}

// Move ...
func (p *Player) Move() {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyUp) && ebiten.IsKeyPressed(ebiten.KeyLeft):
		p.setImage(p.imageUpLeft)
		p.yPos -= p.speed
		p.xPos -= p.speed
	case ebiten.IsKeyPressed(ebiten.KeyUp) && ebiten.IsKeyPressed(ebiten.KeyRight):
		p.setImage(p.imageUpRight)
		p.yPos -= p.speed
		p.xPos += p.speed
	case ebiten.IsKeyPressed(ebiten.KeyDown) && ebiten.IsKeyPressed(ebiten.KeyLeft):
		p.setImage(p.imageDownLeft)
		p.yPos += p.speed
		p.xPos -= p.speed
	case ebiten.IsKeyPressed(ebiten.KeyDown) && ebiten.IsKeyPressed(ebiten.KeyRight):
		p.setImage(p.imageDownRight)
		p.yPos += p.speed
		p.xPos += p.speed
	case ebiten.IsKeyPressed(ebiten.KeyUp):
		p.setImage(p.imageUp)
		p.yPos -= p.speed
	case ebiten.IsKeyPressed(ebiten.KeyDown):
		p.setImage(p.imageDown)
		p.yPos += p.speed
	case ebiten.IsKeyPressed(ebiten.KeyLeft):
		p.setImage(p.imageLeft)
		p.xPos -= p.speed
	case ebiten.IsKeyPressed(ebiten.KeyRight):
		p.setImage(p.imageRight)
		p.xPos += p.speed
	}
}

func (p *Player) Fight(v *village.Village) {
	//	TODO new fight scene transition with village and player
}

func (p *Player) GetPos() (float64, float64) {
	return p.xPos, p.yPos
}

func (p *Player) IsFighting(v *village.Village) bool {
	// if players position +bounds overlap with the village, then return true, else return false
	playerX, playerY := p.GetPos()
	villageX, villageY := v.GetPos()
	if math.Abs(playerX-villageX) < 10 && math.Abs(playerY-villageY) < 10 {
		return true
	}

	return false
}

func (p *Player) setImage(image *ebiten.Image) {
	p.currentImage = image
}

func (p *Player) DrawMap() (*ebiten.Image, *ebiten.DrawImageOptions) {
	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Scale(0.5, 0.5)
	playerOp.GeoM.Translate(p.GetPos())
	return p.currentImage, playerOp
}

func (p *Player) DrawBattle() (*ebiten.Image, *ebiten.DrawImageOptions) {
	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(util.ScreenHeight-100, 100)
	return p.battleImage, playerOp
}
