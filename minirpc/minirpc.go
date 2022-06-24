package minirpc

import (
	"os"
)

func GetFilePath() string {
	return "/tmp/uwu"
}

func SendWakeup(f *os.File) {
	f.WriteString("wakeup")
}

func MessageIsWakeup(text string) bool {
	return text == "wakeup"
}
