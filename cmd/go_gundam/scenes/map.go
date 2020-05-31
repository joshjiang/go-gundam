package scenes

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	b "github.com/joshjiang/go-gundam/cmd/go_gundam/piece/bomb"
	p "github.com/joshjiang/go-gundam/cmd/go_gundam/piece/player"
	v "github.com/joshjiang/go-gundam/cmd/go_gundam/piece/village"
	u "github.com/joshjiang/go-gundam/cmd/go_gundam/util"

	"log"
)

var (
	villages []*v.Village
	bombs    []*b.Bomb
)

type MapScene struct {
	Player          *p.Player
	Bombs           []*b.Bomb
	Villages        []*v.Village
	backgroundImage *ebiten.Image
	width, height   int
}

func NewMapScene() *MapScene {
	var err error

	background, _, err := ebitenutil.NewImageFromFile("../assets/korea.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	villageOne := v.New(300, 500)
	villages := append(villages, villageOne)
	playerOne := p.New(u.ScreenWidth/2.0, u.ScreenHeight/2.0, 4)
	player := playerOne
	bombOne := b.New(u.ScreenWidth, u.ScreenHeight)
	bombs := append(bombs, bombOne)

	return &MapScene{
		Player:          player,
		Bombs:           bombs,
		Villages:        villages,
		backgroundImage: background,
	}
}

func (m *MapScene) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	op.GeoM.Scale(.5,.5)
	screen.DrawImage(m.backgroundImage, op)
	screen.DrawImage(m.Bombs[0].DrawMap())
	screen.DrawImage(m.Villages[0].DrawMap())
	screen.DrawImage(m.Player.DrawMap())
}

func (m *MapScene) Update(state *u.GameState) error {
	m.Player.Move()
	fmt.Println(m.Villages[0].GetPos())
	fmt.Println(m.Player.GetPos())
	if m.Player.IsFighting(m.Villages[0]) {
		//give player each piece an isBattling option
		state.SceneManager.GoTo(NewBattleScene(m.Player, m.Villages[0]))

		//TODO transition to battle scene
	}
	return nil
}
