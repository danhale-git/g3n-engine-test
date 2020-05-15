package main

import (
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/loader/obj"
)

// ImportModel imports a model file into the scene
func ImportModel() (group *core.Node) {
	// Decode model in in OBJ format
	dec, err := obj.Decode("submarine.obj", "submarine.mtl")
	if err != nil {
		panic(err.Error())
	}

	// Create a new node with all the objects in the decoded file and adds it to the scene
	group, err = dec.NewGroup()
	if err != nil {
		panic(err.Error())
	}
	group.SetScale(0.3, 0.3, 0.3)
	group.SetPosition(0.0, -0.8, -0.2)

	return
}
