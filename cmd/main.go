package main

import (
	"log/slog"
	"os"
)

func main() {
	app := &application{
		config: config{
			addr: ":8080",
		},
	}

	if err := app.serve(); err != nil {
		slog.Error("Failed to serve", "error", err)
		os.Exit(1)
	}
}
