package main

import (
	"fmt"

	"github.com/adjing/gamedb_open_api/role"
	"github.com/adjing/gamedb_open_api/sys"
)

func main() {
	var b = role.GetRoleName("a")
	fmt.Println(b)

	var a = sys.GetGUID()
	fmt.Println(a)
}
