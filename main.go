package main

import (
	"app/routes"
	"fmt"
)

func main() {
	e, cfg := routes.NewRouter()

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)))
}
