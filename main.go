package main

import (
	"Renting/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8081")
}
