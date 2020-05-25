package bomb

import (
	"github.com/hajimehoshi/ebiten"
	"math/rand"
	"time"
)

type Bomb struct {
	image *ebiten.Image
	xPos float64
	yPos float64
	isPickedUp bool
}

func New(spriteImage *ebiten.Image, screenWidth,screenHeight float64) *Bomb{
	rand.Seed(time.Now().UnixNano())
	return &Bomb{
		image : spriteImage,
		xPos : rand.Float64() * screenWidth,
		yPos : rand.Float64() * screenHeight,
		isPickedUp : false,
	}
}

//func Draw() image, ops
func (b *Bomb) GetImage() *ebiten.Image {
	return b.image
}

func (b *Bomb) GetPos() (float64 ,float64){
	return b.xPos, b.yPos
}
