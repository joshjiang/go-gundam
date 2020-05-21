package player

import (
	"github.com/hajimehoshi/ebiten"
)

// Player ...
type Player struct {
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
}

// New ...
func New(spriteImage *ebiten.Image, xpos, ypos, speed float64) *Player {
	return &Player{
		image: spriteImage,
		xPos:  xpos,
		yPos:  ypos,
		speed: speed,
	}
}

// Move ...
func (p *Player) Move(x1, x2, y1, y2 float64) {
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

	// if abs(p.yPos - bombOne.yPos) <= 20 && abs(p.xPos - bombOne.xPos) <= 20{
	// 	bombOne.isPickedUp = true

	// 	bombOne.xPos, bombOne.yPos = p.xPos, p.yPos
	// }
}
