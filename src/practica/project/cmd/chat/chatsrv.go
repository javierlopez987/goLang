package main

import (
	"fmt"
	"os"
	"flag"
	"github.com/javierlopez987/goLang/internal/config"
)

func main()  {

	configFile := flag.String("config", "./config.yaml", "este es el servicio de configuracion")
	flag.Parse()
	
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(cfg.DB.Driver)
	fmt.Println(cfg.Version)

}