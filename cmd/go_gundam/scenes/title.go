package scenes

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/joshjiang/go-gundam/cmd/go_gundam/util"
)

type TitleScene struct {
	count      int
}

func (s *TitleScene) Update(state *util.GameState) error {
	s.count++
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.SceneManager.GoTo(NewMapScene())
		return nil
	}
	return nil
}

func (s *TitleScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Press Space!")
}
