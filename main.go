package main

import (
	"internal/controller"
	"internal/model"
)

func main() {
	DB := model.InitDb()
	controller.Main(DB)
}
