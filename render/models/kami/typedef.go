package kami

import (
	"kami/render"

	"github.com/go-gl/mathgl/mgl32"
)

type Model struct {
	Parts []ModelPart
}

type ModelPart struct {
	Name     string
	Position mgl32.Vec3
	Rotation mgl32.Quat

	Vao render.VertexArrayObject

	Vertices      []float32
	TextureCoords []float32
	Normals       []float32

	Indices []uint32
}
