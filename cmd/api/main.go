package main

import (
	"fmt"
	"hook_pipe/internal/core/settings"
	"hook_pipe/internal/server"
)

func main() {
	fmt.Println("hook_pipe v0.0.1")

	settings.LoadDotEnv()

	settings.LoadEnvs()

	server.Run()
}
