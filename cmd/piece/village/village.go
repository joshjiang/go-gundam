package village

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

type Village struct {
	xPos, yPos float64
	image      *ebiten.Image
	health     float64
}

func New(xpos, ypos float64) *Village{
	house, _, err := ebitenutil.NewImageFromFile("../assets/village.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	return &Village{
		xPos:   xpos,
		yPos:   ypos,
		image:  house,
		health: 100,
	}
}

func (v *Village) GetPos() (float64, float64) {
	return v.xPos, v.yPos
}

func (v *Village) Draw() (*ebiten.Image, *ebiten.DrawImageOptions) {
	villageOp := &ebiten.DrawImageOptions{}
	villageOp.GeoM.Translate(v.GetPos())
	villageOp.GeoM.Scale(0.5, 0.5)
	return v.image, villageOp
}
