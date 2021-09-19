package main

import (
	"MyProject/ProjectALTA/JobDir/configs"
	"MyProject/ProjectALTA/JobDir/routes"
)

func main() {
	configs.InitDB()
	e := routes.NewRoute()
	e.Start(":8000")
}
