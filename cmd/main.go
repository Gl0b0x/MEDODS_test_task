package main

import (
	"MEDODS/configs"
	"MEDODS/pkg/app"
	"log"
)

func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	app.Run(cfg)
	//storage := make(map[string]*User)
	//guid := uuid.New().String()
	//storage[guid] = &User{guid, "vanyakyz@mail.ru", "95.220.56.245", NewRefreshToken()}
	//db := Storage{storage}
	//db.login(guid)

}
