package go_gundam

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/joshjiang/go-gundam/cmd/go_gundam/scenes"
	"github.com/joshjiang/go-gundam/cmd/go_gundam/util"
)

type Game struct {
	sceneManager *util.SceneManager
}

func (g *Game) Update(r *ebiten.Image) error {
	if g.sceneManager == nil {
		g.sceneManager = &util.SceneManager{}
		g.sceneManager.GoTo(&scenes.TitleScene{})
	}

	// g.sceneManager.Update() makes the sceneManager draw to the screen
	if err := g.sceneManager.Update(); err != nil{
		return err
	}
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	g.sceneManager.Draw(r)
	return nil
}
