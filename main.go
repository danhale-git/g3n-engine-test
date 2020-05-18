package main

import (
	"time"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/window"

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

	// Create scene and camera
	scene := core.NewNode()
	cam := createCamera(scene, a)

	// Create lighting
	createLighting(scene)

	// Create debug widgets
	createDebugWidgets(scene)

	// Set background color to blue
	a.Gls().ClearColor(0.5, 0.7, 0.8, 1.0)

	// Set the scene to be managed by the gui manager
	gui.Manager().Set(scene)

	//	Add a model to the scene from a .obj file
	Player = importModel()
	scene.Add(Player)

	// Player movement user unput
	a.Subscribe(window.OnKeyDown, movePlayer)

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
