package main

import (
	"os"

	"github.com/jawher/mow.cli"

	"github.com/forana/skycam/camera"
)

func main() {
	app := cli.App("skycam", "Skycam")
	app.Command("capture", "Capture a single image", capture)
	app.Run(os.Args)
}

func capture(cmd *cli.Cmd) {
	outputDir := cmd.StringArg("OUTPUT_DIR", "", "Relative path to output directory")
	cameraIndex := cmd.IntOpt("i index", 0, "Index of the camera to use")

	cmd.Action = func() {
		camera.Capture(*outputDir, *cameraIndex)
	}
}
