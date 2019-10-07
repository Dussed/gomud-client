package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const playerSpeed = 1

type player struct {
	tex  *sdl.Texture
	x, y float64
}

func NewPlayer(renderer *sdl.Renderer) (p player, err error) {
	img, err := sdl.LoadBMP("assets/player.bmp")

	if err != nil {
		fmt.Println("player load", err)
		return player{}, err
	}

	defer img.Free()

	p.tex, err = renderer.CreateTextureFromSurface(img)

	if err != nil {
		fmt.Println("player texture", err)
		return player{}, nil
	}

	p.x = windowWidth / 2.0
	p.y = windowHeight - 20

	return p, nil
}

func (p *player) draw(renderer *sdl.Renderer) {

	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: 20, H: 20},
		&sdl.Rect{X: int32(p.x), Y: int32(p.y), W: 20, H: 20})

}

func (p *player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		// left
		p.x -= playerSpeed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		// right
		p.x += playerSpeed
	}

}
