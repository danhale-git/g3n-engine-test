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
//var Player *core.Node

var application *app.Application

func main() {

	// Create application
	application = app.App()

	// Create scene and camera
	scene := core.NewNode()
	cam := createCamera(scene, application)

	// Create lighting
	createLighting(scene)

	// Create debug widgets
	createDebugWidgets(scene)

	// Set background color to blue
	application.Gls().ClearColor(0.5, 0.7, 0.8, 1.0)

	// Set the scene to be managed by the gui manager
	gui.Manager().Set(scene)

	// Add a model to the scene from a .obj file
	var player *core.Node = importModel()
	scene.Add(player)

	// Run the application
	application.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {

		//	Call update once per frame
		update(player, deltaTime)

		// Render
		application.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
		renderer.Render(scene, cam)
	})
}

func update(player *core.Node, deltaTime time.Duration) {
	// Move player based on which keys are pressed
	movePlayer(player, deltaTime)
}

func movePlayer(player *core.Node, deltaTime time.Duration) {

	keyState := application.KeyState()

	if keyState.Pressed(window.KeyA) {
		player.RotateY(float32(deltaTime.Seconds()))
	}

	if keyState.Pressed(window.KeyD) {
		player.RotateY(float32(-deltaTime.Seconds()))
	}

	if keyState.Pressed(window.KeyW) {
		forward := nodeForward(player, deltaTime)
		newPosition := player.Position()
		newPosition.Add(forward)
		player.SetPositionVec(&newPosition)
	}

	if keyState.Pressed(window.KeyS) {
		forward := nodeForward(player, deltaTime)
		newPosition := player.Position()
		newPosition.Add(forward.Negate())
		player.SetPositionVec(&newPosition)
	}

	keyState.Dispose()
}

func nodeForward(n *core.Node, deltaTime time.Duration) *math32.Vector3 {
	var d *math32.Vector3 = &math32.Vector3{}
	n.WorldDirection(d)
	return d.MultiplyScalar(float32(deltaTime.Seconds()))
}
