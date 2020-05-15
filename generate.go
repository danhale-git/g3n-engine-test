package main

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

// Plane returns a new flat plane with given width and length
func Plane(width, length float32) *graphic.Mesh {
	// Create a list of vertices
	positions := math32.NewArrayF32(0, 0)

	//	Create a color
	color := math32.Color{
		R: 1,
		G: 1,
		B: 1,
	}

	//	Add vertices in groups of 3 to create triangles
	positions.Append(1, 1, 0, color.R, color.G, color.B)
	positions.Append(0, 1, 0, color.R, color.G, color.B)
	positions.Append(0, 0, 0, color.R, color.G, color.B)

	// Create geometry
	geom := geometry.NewGeometry()
	geom.AddVBO(
		gls.NewVBO(positions).
			AddAttrib(gls.VertexPosition).
			AddAttrib(gls.VertexColor),
	)

	mat := material.NewStandard(math32.NewColor("DarkBlue"))

	return graphic.NewMesh(geom, mat)
}
