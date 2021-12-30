package main

import (
	"log"

	"github.com/nikitamirzani323/go_wl_api/db"
	"github.com/nikitamirzani323/go_wl_api/routes"
)

func main() {
	db.Init()
	app := routes.Init()
	log.Fatal(app.Listen(":7073"))
}
