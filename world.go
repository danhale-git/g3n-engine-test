package main

import (
	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/engine/window"
)

// InitWorld creates a new scene with camera and lighting
func InitWorld(a *app.Application) (scene *core.Node, cam *camera.Camera) {
	// Create scene and camera
	scene = core.NewNode()
	cam = initCamera(scene, a)

	// Create lighting
	initLighting(scene)

	// Create debug widgets
	initDebug(scene)

	// Set background color to blue
	a.Gls().ClearColor(0.5, 0.7, 0.8, 1.0)

	// Set the scene to be managed by the gui manager
	gui.Manager().Set(scene)

	// Player movement user unput
	a.Subscribe(window.OnKeyDown, MovePlayer)

	return
}

func initCamera(scene *core.Node, a *app.Application) (cam *camera.Camera) {
	// Create perspective camera
	cam = camera.New(1)
	cam.SetPosition(0, 0, 3)
	scene.Add(cam)

	// Set up orbit control for the camera
	camera.NewOrbitControl(cam)

	// Set up callback to update viewport and camera aspect ratio when the window is resized
	onResize := func(evname string, ev interface{}) {
		// Get framebuffer size and update viewport accordingly
		width, height := a.GetSize()
		a.Gls().Viewport(0, 0, int32(width), int32(height))
		// Update the camera's aspect ratio
		cam.SetAspect(float32(width) / float32(height))
	}
	a.Subscribe(window.OnWindowSize, onResize)
	onResize("", nil)

	return
}

func initLighting(scene *core.Node) {
	// Create and add lights to the scene
	scene.Add(light.NewAmbient(&math32.Color{1.0, 1.0, 1.0}, 0.8))
	pointLight := light.NewPoint(&math32.Color{1, 1, 1}, 5.0)
	pointLight.SetPosition(1, 0, 2)
	scene.Add(pointLight)
}

func initDebug(scene *core.Node) {
	// Create and add an axis helper to the scene
	scene.Add(helper.NewAxes(0.5))
}
