package main

import (
	"time"

	"github.com/g3n/engine/app"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/renderer"
)

func main() {

	// Create application
	a := app.App()
	scene, cam := InitWorld(a)

	scene.Add(ImportModel())

	// Run the application
	a.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {

		//	Call update once per frame
		update()

		// Render
		a.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
		renderer.Render(scene, cam)
	})
}

func update() {
}
