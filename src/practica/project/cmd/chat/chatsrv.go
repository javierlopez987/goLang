package main

import (
	"fmt"
	"os"
	"flag"

	"github.com/javierlopez987/goLang/internal/config"
	"github.com/javierlopez987/goLang/internal/database"
	"github.com/javierlopez987/goLang/internal/service/chat"

	"github.com/gin-gonic/gin"

)

func main()  {

	// Lectura de configuracion
	cfg := readConfig()

	// Instanciacion de db
	db, err := database.NewDatabase(cfg)
	defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Instanciacion de un servicio y le inyecta la configuracion y la base de datos
	service, _ := chat.New(db, cfg)
	httpService := chat.NewHTTPTransport(service)

	r:= gin.Default()
	httpService.Register(r)
	r.Run()
}

func readConfig() *config.Config {
	configFile := flag.String("config", "./config.yaml", "este es el servicio de configuracion")
	flag.Parse()
	
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return cfg
}