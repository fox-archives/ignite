package main

import (
	"bufio"
	"errors"
	"os"
	"syscall"

	"github.com/faiface/mainthread"
	"github.com/hyperupcall/ignite/launcher"
	"github.com/hyperupcall/ignite/minirpc"
	"github.com/hyperupcall/ignite/util"
)

var showLauncher bool = false

func run() {
	mainthread.CallNonBlock(func() {
		for {
			if showLauncher {
				launcher.Launch()
				showLauncher = false
			}
		}
	})

	filePath := minirpc.GetFilePath()
	err := syscall.Mkfifo(filePath, 0400)
	if !errors.Is(err, os.ErrExist) { // TODO
		util.Handle(err)
	}
	file, err := os.Open(filePath)

	s := bufio.NewScanner(file)
	for {
		for s.Scan() {
			text := s.Text()

			if minirpc.MessageIsWakeup(text) {
				showLauncher = true
			}
		}
		util.Handle(s.Err())
	}
}

func main() {
	mainthread.Run(run)
}
