package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"kami/render/models/kami"
	"kami/render/models/minecraftjson"
	"kami/render/models/obj"

	"github.com/urfave/cli/v2"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	/*err := glfw.Init()
	util.FCheckErr(err, "could not initialize glfw: %v")
	defer glfw.Terminate()

	glfw.DefaultWindowHints()
	glfw.WindowHint(glfw.Visible, glfw.False)

	if runtime.GOOS == "macos" {
		glfw.WindowHint(glfw.CocoaMenubar, glfw.False)
	}

	window, err := glfw.CreateWindow(1, 1, "", nil, nil)
	util.FCheckErr(err, "could not create OpenGL window: %v")
	window.MakeContextCurrent()

	render.InitGL()

	go func() {
		for !window.ShouldClose() {
			render.CheckGlError()
			glfw.PollEvents()
		}
	}()*/

	var model string

	app := &cli.App{
		Name:                 "kami",
		Version:              "v1.0.0",
		Usage:                "export OBJ or Minecraft JSON models as OpenGL VAOs",
		EnableBashCompletion: true,
		Action: func(c *cli.Context) error {
			if c.NArg() > 0 {
				model = c.Args().Get(0)
			}

			_, err := os.Stat(model)
			if err != nil {
				log.Fatal(err)
			}

			if !os.IsNotExist(err) {
				var kamiModel kami.Model

				switch file := path.Ext(model); file {
				case ".obj":
					kamiModel = obj.LoadModel(model)
				case ".json":
					kamiModel = minecraftjson.LoadModel(model)
				}

				data, err := json.Marshal(kamiModel)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(string(data))
			}

			if err != nil {
				log.Fatal(err)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
