package main

import (
	"os"

	"github.com/hyperupcall/ignite/minirpc"
	"github.com/hyperupcall/ignite/util"
)

func main() {
	filePath := minirpc.GetFilePath()
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0644)
	util.Handle(err)

	minirpc.SendWakeup(file)
	// TODO: ensure the newly (or existing) Launcher is _active_
}
