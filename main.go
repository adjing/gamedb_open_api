package main

import "gamedb_open_api"

func main() {

	gamedb_open_api.InitGinRoute()
}

/*
windows:
> go env -w GOOS=windows
> go env GOOS
*/
