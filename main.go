package main

import (
	"fmt"

	"github.com/adjing/gamedb_open_api"
)

// package gamedb_open_api

func main() {

	var a = gamedb_open_api.InitSystemData("a")
	fmt.Println(a)
	// gamedb_open_api.InitGinRoute()
}

/*
windows:
> go env -w GOOS=windows
> go env GOOS
*/
