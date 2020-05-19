package main

import (
	"fmt"
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
//var Player *core.Node

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

	// Add a model to the scene from a .obj file
	var player *core.Node = importModel()
	scene.Add(player)

	// Run the application
	a.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {

		//	Call update once per frame
		update(*a.KeyState(), player)

		// Render
		a.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
		renderer.Render(scene, cam)
	})
}

func update(ks window.KeyState, player *core.Node) {
	// Move player based on which keys are pressed
	movePlayer(ks, player)
}

func movePlayer(keyState window.KeyState, player *core.Node) {

	// Move the model along the Z axis
	var pos math32.Vector3 = player.Position()
	player.SetPosition(pos.X, pos.Y, pos.Z-0.01)

	switch {
	case keyState.Pressed(window.KeyA):
		fmt.Println(player.Direction(), player.Position())
		player.RotateY(0.1)
	case keyState.Pressed(window.KeyD):
		player.RotateY(-0.1)
		//TODO: Move forward and backward, see: https://github.com/g3n/g3nd/blob/4a8cdd403c411cd2acb0e62c61828fd073bef340/demos/other/tank.go#L102
	}

	keyState.Dispose()
}
