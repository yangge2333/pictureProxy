package main

import (
	_ "pictureProxy/boot"
	_ "pictureProxy/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
