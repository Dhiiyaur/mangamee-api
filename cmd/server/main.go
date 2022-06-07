package main

import (
	"mangamee-api/internal/app"
	"mangamee-api/internal/config"
)

func init() {
	config.InitEnvConfiguration()
}

func main() {

	app.Start()
}
