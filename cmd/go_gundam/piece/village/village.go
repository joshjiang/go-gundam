package village

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/joshjiang/go-gundam/cmd/go_gundam/util"
	"log"
)

type Village struct {
	xPos, yPos float64
	image, battleImage      *ebiten.Image
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
		//TODO Update village battle image
		battleImage: house,
		health: 100,
	}
}

func (v *Village) GetPos() (float64, float64) {
	return v.xPos, v.yPos
}

func (v *Village) DrawMap() (*ebiten.Image, *ebiten.DrawImageOptions) {
	villageOp := &ebiten.DrawImageOptions{}
	villageOp.GeoM.Scale(0.25, 0.25)
	villageOp.GeoM.Translate(v.GetPos())
	return v.image, villageOp
}

func (v *Village) DrawBattle() (*ebiten.Image, *ebiten.DrawImageOptions) {
	villageOp := &ebiten.DrawImageOptions{}
	villageOp.GeoM.Translate(util.ScreenWidth-200, 100)
	return v.battleImage, villageOp
}