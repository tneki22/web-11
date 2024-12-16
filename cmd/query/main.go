package main

import (
	"flag"

	"web-11/internal/query/api"
	"web-11/internal/query/config"
	"web-11/internal/query/provider"
	"web-11/internal/query/usecase"

	"log"

	_ "github.com/lib/pq"
)

func main() {
	// main - точка входа в микросервис query.
	configPath := flag.String("config-path", "C:/Users/1234/web-11-master/configs/query.yaml", "путь к файлу конфигурации")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	use := usecase.NewUsecase(cfg.Usecase.DefaultMessage, prv)
	srv := api.NewServer(cfg.IP, cfg.Port, cfg.API.MaxMessageSize, use)

	srv.Run()
}
