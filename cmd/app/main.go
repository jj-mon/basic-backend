package main

import (
	"basic-backend/internal/app"
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	a, err := app.New(ctx)
	if err != nil {
		log.Fatalf("main: %q", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("main: %q", err)
	}
}
