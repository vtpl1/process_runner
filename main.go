package main

import (
	"os/exec"

	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Hello from Zerolog global logger")
	cmd := exec.Command("ffmpeg", "-v")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
