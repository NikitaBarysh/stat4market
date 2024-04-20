package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/NikitaBarysh/stat4market/internal/app"
	"github.com/NikitaBarysh/stat4market/internal/handler"
	"github.com/NikitaBarysh/stat4market/internal/repository"
	"github.com/NikitaBarysh/stat4market/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := repository.NewDB(ctx,
		viper.GetString("clickhouse.address"),
		viper.GetString("clickhouse.port"),
		viper.GetString("clickhouse.database"),
		viper.GetString("clickhouse.username"),
		viper.GetString("clickhouse.password"))
	if err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()

	rep := repository.NewRepository(db)
	srv := service.NewService(rep)
	hand := handler.NewHandler(srv)

	go srv.Scheduler(ctx)

	hand.InitRouters(router)

	server := new(app.Server)

	go func() {
		if err = server.Run(viper.GetString("port"), router); err != nil {
			log.Fatal("Err to start server: ", err)
		}
	}()

	termSig := make(chan os.Signal, 1)
	signal.Notify(termSig, syscall.SIGTERM, syscall.SIGINT)
	<-termSig

	if err = server.ShutDown(); err != nil {
		log.Fatal("err to shutdown", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
