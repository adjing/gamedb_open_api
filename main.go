package main

import (
	"fmt"
	"gamedb_open_api/src/api"
)

// package gamedb_open_api

func main() {

	var a = api.InitSystemData("a")
	fmt.Println(a)
	// gamedb_open_api.InitGinRoute()
}

/*
windows:
> go env -w GOOS=windows
> go env GOOS
*/
