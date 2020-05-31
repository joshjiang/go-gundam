package bomb

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
	"math/rand"
	"time"
)

type Bomb struct {
	image      *ebiten.Image
	xPos       float64
	yPos       float64
	isPickedUp bool
}

func New(screenWidth, screenHeight float64) *Bomb {
	napalm, _, err := ebitenutil.NewImageFromFile("../assets/napalm.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano())
	return &Bomb{
		image:      napalm,
		xPos:       rand.Float64() * screenWidth,
		yPos:       rand.Float64() * screenHeight,
		isPickedUp: false,
	}
}

func (b *Bomb) GetImage() *ebiten.Image {
	return b.image
}

func (b *Bomb) GetPos() (float64, float64) {
	return b.xPos, b.yPos
}

func (b *Bomb) DrawMap() (*ebiten.Image, *ebiten.DrawImageOptions) {
	bombOp := &ebiten.DrawImageOptions{}
	bombOp.GeoM.Scale(0.2, 0.2)
	bombOp.GeoM.Translate(b.GetPos())
	return b.image, bombOp
}

func (b *Bomb) DrawBattle() (*ebiten.Image, *ebiten.DrawImageOptions) {
	bombOp := &ebiten.DrawImageOptions{}
	bombOp.GeoM.Scale(0.2, 0.2)
	bombOp.GeoM.Translate(b.GetPos())
	return b.image, bombOp
}