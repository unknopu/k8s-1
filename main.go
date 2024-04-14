package main

import (
	"log"
	"testapi/core/app"
	"testapi/core/config"
	conn "testapi/core/connection"
	r "testapi/router"
)

func main() {
	cf, err := config.InitConfig("./configs")
	if err != nil {
		panic(err.Error())
	}

	//
	// r.NewInsecure(cf).Start(cf.Config.DEBUGPORT)

	database, err := conn.New(&conn.Config{
		Host:         cf.Config.Database.Host,
		Port:         cf.Config.Database.Port,
		User:         cf.Config.Database.Username,
		Password:     cf.Config.Database.Password,
		DatabaseName: cf.Config.Database.Name,
		Debug:        true,
	})
	if err != nil {
		log.Println("Postgres was failed, %v", err.Error())
	}

	option := r.Option{
		AppContext: &app.Context{
			Db:     database,
			Config: cf.Config,
		},
	}

	if err := r.NewWithOptions(option).Start(cf.Config.ServerPort); err != nil {
		panic(err.Error())
	}
}
