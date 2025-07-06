package main

import (
	"log"

	"gorm-gen/repository"
	"gorm-gen/service"
)

func main() {
	gdb := repository.NewConnection()
	gdb.Init()
	gdb.Gen()

	service.UserService(gdb)
	log.Println("done")
}
