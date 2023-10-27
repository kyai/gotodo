package main

import (
	"gotodo/serve"

	_ "gotodo/serve/migrations"
)

func main() {
	serve.HttpHandler()
}
