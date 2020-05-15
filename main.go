package main

import (
	"time"

	"github.com/g3n/engine/math32"

	"github.com/g3n/engine/app"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/renderer"
)

// Player is the player object
var Player *core.Node

func main() {

	// Create application
	a := app.App()
	scene, cam := InitWorld(a)

	//	Add a model to the scene from a .obj file
	Player = ImportModel()
	scene.Add(Player)

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
	// Move the model along the Z axis
	var pos math32.Vector3 = Player.Position()
	Player.SetPosition(pos.X, pos.Y, pos.Z+0.01)
}
