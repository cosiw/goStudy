package main

import (
	"fmt"
	"go-server/pkg/api"
	"go-server/pkg/db"
	"go-server/pkg/router"
)

func main() {

	//dbHandler, err := db.NewAndConnectGorm("user:password@tcp(127.0.0.1:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local")
	dbHandler, err := db.NewAndConnectGorm("user:password@tcp(127.0.0.1:3308)/dev?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	apis := api.NewAPI(dbHandler)
	r := router.Router(apis)

	r.Run(":8085")
}
