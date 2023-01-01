package main

import "2ND_PRACTICE/pkg/router"

func main() {
	r := router.Router()

	r.Run(":8081")
}
