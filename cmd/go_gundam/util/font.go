package util

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"io/ioutil"
	"log"
)

const (
	arcadeFontBaseSize = 8
)

var (
	gameFonts map[int]font.Face
)

func GetGameFonts(scale int) font.Face {
	fontBytes, err := ioutil.ReadFile("../assets/pokemon-classic.ttf")
	if err != nil {
		log.Fatal(err)
	}
	if gameFonts == nil {
		tt, err := truetype.Parse(fontBytes)
		if err != nil {
			log.Fatal(err)
		}
		gameFonts = map[int]font.Face{}
		for i := 1; i <= 4; i++ {
			const dpi = 72
			gameFonts[i] = truetype.NewFace(tt, &truetype.Options{
				Size:    float64(arcadeFontBaseSize * scale),
				DPI:     dpi,
				Hinting: font.HintingFull,
			})
		}
	}
	return gameFonts[scale]
}