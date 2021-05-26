package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func openBrowser(url string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	default:
		return fmt.Errorf("unsupported platform")
	}
}

func main() {
	url := "https://build.golang.org/"

	err := openBrowser(url)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("done")
}
