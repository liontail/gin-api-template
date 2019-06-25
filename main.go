package main

import (
	"log"

	"github.com/liontail/go-api-template/router"
)

func main() {
	r := router.SetupRouter()
	if err := r.Run(port); err != nil {
		log.Panic(err)
	}
}
