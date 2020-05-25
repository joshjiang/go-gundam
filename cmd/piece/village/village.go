package village

import "github.com/hajimehoshi/ebiten"

type Village struct {
	xPos, yPos float64
	image      *ebiten.Image
	health     float64
}

func New(image *ebiten.Image, xPos, yPos float64) *Village{
	return &Village{
		xPos:   500,
		yPos:   500,
		image:  nil,
		health: 100,
	}
}

func (v *Village) GetPos() (float64, float64) {
	return v.xPos, v.yPos
}

func (v *Village) Draw() (*ebiten.Image, *ebiten.DrawImageOptions) {
	villageOp := &ebiten.DrawImageOptions{}
	villageOp.GeoM.Translate(v.GetPos())
	return v.image, villageOp
}
