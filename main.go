package main

import (
	"log"
	"testapi/core/app"
	"testapi/core/config"
	conn "testapi/core/connection"

	"testapi/core/logger"
	r "testapi/router"
)

func main() {
	cf, err := config.InitConfig("./configs")
	if err != nil {
		panic(err.Error())
	}

	logFile, err := logger.LoggerFile(cf.Config.SERVICENAME, cf.Config.LOGPATH)
	if err != nil {
		panic(err.Error())
	}
	defer logFile.Close()

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
		log.Println("status:Internal Server Error,error:code=500, message=failed to connect")
	}

	option := r.Option{
		LogFile: logFile,
		AppContext: &app.Context{
			Db:     database,
			Config: cf.Config,
		},
	}

	if err := r.NewWithOptions(option).Start(cf.Config.ServerPort); err != nil {
		panic(err.Error())
	}
}
