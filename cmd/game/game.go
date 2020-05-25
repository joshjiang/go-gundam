package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/joshjiang/go-gundam/cmd/pieces/bomb"
	"github.com/joshjiang/go-gundam/cmd/player"
	"log"
)

type Board struct {
	Players []*player.Player
	Bombs   []*bomb.Bomb
	backgroundImage *ebiten.Image
	width, height int

}

func New(players []*player.Player, bombs []*bomb.Bomb) *Board {
	var err error

	background, _, err := ebitenutil.NewImageFromFile("../assets/korea.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	return &Board {
		Players : players,
		Bombs : bombs,
		backgroundImage : background,
	}
}

func (b *Board) Update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(b.backgroundImage, op)

	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Scale(2.0, 2.0)
	playerOp.GeoM.Translate(b.Players[0].GetPos())
	screen.DrawImage(b.Players[0].GetImage(), playerOp)
	b.Players[0].Move()

	bombOp := &ebiten.DrawImageOptions{}
	bombOp.GeoM.Scale(0.2, 0.2)
	bombOp.GeoM.Translate(b.Bombs[0].GetPos())
	screen.DrawImage(b.Bombs[0].GetImage(), bombOp)

	return nil
}
