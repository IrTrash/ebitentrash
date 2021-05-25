package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"

	"trash.com/Sprite"
	"trash.com/unit"
)

var sprites map[string]*Sprite.Sprite
var units map[string]*unit.Unit

type Game struct{}

func (g *Game) Update() error {
	for _, sp := range sprites {
		sp.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, sp := range sprites {
		screen.DrawImage(sp.Img, sp.Getdrawoption())
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	sprites = map[string]*Sprite.Sprite{}
	sprites["s1"] = &Sprite.Sprite{Imgsrc: "ddd.png", X: 200, Y: 100, Direction: 240 * (3.141592 / 180), Speed: 0.4} //여기도 y가 커지면 아래로가는축인듯
	sprites["s2"] = &Sprite.Sprite{Imgsrc: "ddd.png", X: 400, Y: 100, Direction: 90 * (3.141592 / 180), Speed: 0.5}

	for _, sp := range sprites {
		sp.Init()
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Trash")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
