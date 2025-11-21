package main

import "github.com/amlcx/tablero/backend/wiring"

func main() {
	app := wiring.NewApp()

	if err := app.Start(); err != nil {
		panic(err)
	}
}
