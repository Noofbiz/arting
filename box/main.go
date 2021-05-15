package main

import (
	"github.com/unixpickle/model3d/model3d"
	"github.com/unixpickle/model3d/render3d"
)

func main() {
	bitmap := NewBitmapSlab("in/pup.png", nil, model3d.Coord3D{}, 0.1, 1, 1)
	// negamap := NewBitmapSlab("in/heart.png", nil, model3d.Coord3D{X: -0.1, Y: -0.1}, 0.08, 0.9, 0.9)
	// sub := &model3d.SubtractedSolid{Positive: bitmap, Negative: negamap}

	mesh := model3d.MarchingCubesSearch(bitmap, 0.01, 8)
	mesh.SaveGroupedSTL("out/pup.stl")
	render3d.SaveRandomGrid("out/pup.png", mesh, 3, 3, 300, nil)
}
