package test

import (
	"bytes"
	"github.com/go-gl/gl/all-core/gl"
	"image"
	"image/draw"
	"image/png"
	"kami/render"
	"kami/util"
)

//TODO move to designated file
func BindIndices(size int, data []int32) {
	var vbo uint32
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, size, gl.Ptr(data), gl.STATIC_DRAW)
}

//TODO move to designated file
func LoadTexture(fileName string) uint32 {
	data := util.ReadAsset(fileName)
	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		util.CheckErr(err)
		return 0
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return 0
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.GenTextures(1, &texture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(rgba.Rect.Size().X), int32(rgba.Rect.Size().Y), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))

	return texture
}

//TODO move to designated file
func LoadProjectionMatrix(shader *render.ShaderProgram) {
	shader.UseShader()
	matrixID := shader.CreateUniformLocation("projectionMatrix")
	gl.UniformMatrix4fv(matrixID, 1, false, &render.MainCamera.Projection[0])
	render.CheckGlError()
	//cleanup
	gl.UseProgram(0)
}
