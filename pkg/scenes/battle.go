package scenes

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/joshjiang/go-gundam/pkg/piece"
	p "github.com/joshjiang/go-gundam/pkg/piece/player"
	u "github.com/joshjiang/go-gundam/pkg/util"
	"log"
)

type BattleScene struct {
	enemy                     piece.Piece
	player                    *p.Player
	background, battleCircles *ebiten.Image
}

func NewBattleScene(playerOne *p.Player, pi piece.Piece) *BattleScene {
	var err error

	//TODO make battle background
	battleBackground, _, err := ebitenutil.NewImageFromFile("../assets/battle/battle-bg.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	battleCircles, _, err := ebitenutil.NewImageFromFile("../assets/battle/battle-circles.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	return &BattleScene{
		player:        playerOne,
		enemy:         pi,
		background:    battleBackground,
		battleCircles: battleCircles,
	}
}

func (b *BattleScene) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(2.75, 2.75)
	op.GeoM.Translate(0, 0)
	if err := screen.DrawImage(b.background, op); err != nil {
		log.Fatal(err)
	}
	fmt.Println(b.background.Bounds().Max.Y)
	if err := screen.DrawImage(b.battleCircles, op); err != nil {
		log.Fatal(err)
	}
	if err := screen.DrawImage(b.enemy.DrawBattle()); err != nil {
		log.Fatal(err)
	}
	if err := screen.DrawImage(b.player.DrawBattle()); err != nil {
		log.Fatal(err)
	}
}

func (b *BattleScene) drawDialog(screen *ebiten.Image) {
	// 	piecename wants to battle
}

func (b *BattleScene) Update(state *u.GameState) error {
	// TODO if !player.Health == 0 or !enemy.Health ==0 { 1. move scene 2. battle scene}

	return nil
}
