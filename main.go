package main

import (
	_ "gotodo/serve"
	_ "gotodo/serve/config"
	_ "gotodo/serve/migrations"

	"gotodo/serve/services"
)

func main() {
	services.HttpHandler()
}
