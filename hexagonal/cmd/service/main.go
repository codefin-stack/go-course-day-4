package main

import (
	"hexagonal/internal/container"
)

func main() {
	c := container.NewService()
	if err := c.Start(); err != nil {
		panic(err)
	}
}
