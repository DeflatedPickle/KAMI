package render

import (
	"fmt"
	"kami/util"

	"github.com/go-gl/gl/all-core/gl"
)

const logSize = 512

var isGlInit = false

func LoadVBO(vbo *VertexBufferObject) {
	InitGL()
	gl.GenBuffers(1, &vbo.Handle)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo.Handle)
	gl.BufferData(gl.ARRAY_BUFFER, len(vbo.Vertices)*4, gl.Ptr(vbo.Vertices), gl.STATIC_DRAW)

	//cleanup
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func InitGL() {
	if !isGlInit {
		util.CheckErr(gl.Init())
		// util.Log.Println(fmt.Sprintf("OpenGL Version %v", gl.GoStr(gl.GetString(gl.VERSION))))
		isGlInit = true
	}
}

func LoadVAO(vao *VertexArrayObject) {
	InitGL()
	if vao.BufferCount <= 0 {
		util.ErrLog.Panicln("tried to create empty VAO!")
	}
	vao.VertexBuffers = make([]VertexBufferObject, vao.BufferCount)
	gl.GenVertexArrays(vao.BufferCount, &vao.Handle)
}

func CheckGlError() {
	for err := gl.GetError(); err != gl.NO_ERROR; err = gl.GetError() {
		util.ErrLog.Println(fmt.Sprintf("OpenGL ERROR: %v", err))
	}
}
