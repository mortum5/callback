package main

import (
	"log/slog"
	"os"
	"os/signal"

	"github.com/mortum5/callback/internal/db"
	"github.com/mortum5/callback/internal/server"
	"github.com/mortum5/callback/internal/service"
)

func main() {
	db := db.New()
	db.Migrate()

	service := service.New(db)
	slog.Info("service created")

	service.Start()
	slog.Info("service started")

	server := server.New(service)
	slog.Info("server created")

	server.Start()
	slog.Info("server started")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	slog.Info("canceled")

}
