package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

func main() {
	fmt.Println("GoMud-Client Launching")

	// Init SDL
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("init SDL", err)
		return
	}

	// Create the window
	window, err := sdl.CreateWindow(
		"GoMud Client",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		windowWidth, windowHeight,
		sdl.WINDOW_OPENGL)

	if err != nil {
		fmt.Println("init window", err)
		return
	}

	defer window.Destroy() // Destroy window at the end

	// Create the renderer
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	if err != nil {
		fmt.Println("init renderer", err)
		return
	}

	defer renderer.Destroy() // Destroy renderer at the end

	playerInstance, err := NewPlayer(renderer)

	if err != nil {
		fmt.Println("playerinstance", err)
	}

	// Draw loop
	for {
		// Check for sdlEventQueue
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

			switch event.(type) {

			case *sdl.QuitEvent:
				return

			}

		}

		renderer.SetDrawColor(255, 0, 0, 255)
		renderer.Clear()

		playerInstance.draw(renderer)
		playerInstance.update()

		renderer.Present()
	}

}
