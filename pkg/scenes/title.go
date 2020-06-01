package scenes

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/joshjiang/go-gundam/pkg/util"
	"image/color"
)

type TitleScene struct {
	count int
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
	text.Draw(screen, "Gundam AC Fighter", util.GetGameFonts(2), 250, 250, color.White)
	text.Draw(screen, "Press SPACE to begin", util.GetGameFonts(2), 250, 280, color.White)
}
