package main

import (
	"gin_project/app"
)

func main() {
	r := app.Init()
	_ = r.Run(":8888")
}
