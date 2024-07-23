package main

import (
	"barosa.fun/azure-ai-stream-backend/environment"
	"barosa.fun/azure-ai-stream-backend/server"
)

func main() {
	environment.Init()
	environment.CheckEnvVars()
	server.Init()
}
