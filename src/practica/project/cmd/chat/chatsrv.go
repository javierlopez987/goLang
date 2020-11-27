package main

import (
	"fmt"
	"os"
	"flag"
	"github.com/javierlopez987/goLang/internal/config"
	"github.com/javierlopez987/goLang/internal/service/chat"
	"github.com/javierlopez987/goLang/internal/database"
)

func main()  {

	cfg := readConfig()

	
	db, err := database.NewDatabase(cfg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	
	service, _ := chat.New(db, cfg)

	for _, m := range service.FindAll() {
		fmt.Println(m)
	}

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