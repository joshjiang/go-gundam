package player

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

// Player ...
type Player struct {
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
}

// New ...
func New(xpos, ypos, speed float64) *Player {
	gundam, _, err := ebitenutil.NewImageFromFile("../assets/gundam.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	return &Player{
		image: gundam,
		xPos:  xpos,
		yPos:  ypos,
		speed: speed,
	}
}

// Move ...
func (p *Player) Move() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.yPos -= p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.yPos += p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.xPos -= p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.xPos += p.speed
	}
}

func (p *Player) GetPos() (float64, float64) {
	return p.xPos, p.yPos
}

func (p *Player) GetImage() *ebiten.Image {
	return p.image
}

func (p *Player) setImage() {

}

func (p *Player) Draw() (*ebiten.Image, *ebiten.DrawImageOptions) {
	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Scale(4.0, 4.0)
	playerOp.GeoM.Translate(p.GetPos())
	return p.GetImage(), playerOp
}


