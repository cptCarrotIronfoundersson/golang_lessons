package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/cmd"
	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/app"
	internalgrpc "github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/server/grpc"
	internalhttp "github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/server/http"
	sqlstorage "github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/storage/sql"
	_ "github.com/lib/pq"
)

func init() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	conf := cmd.Config.NewConfig()
	logg := cmd.Logger

	storage := sqlstorage.New()
	// storage := memorystorage.New()
	calendar := app.New(logg, storage)
	grpcserver := internalgrpc.NewServer(logg, conf, calendar)
	httpserver := internalhttp.NewServer(logg, conf, calendar)
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		if err := httpserver.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
		if err := grpcserver.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	if err := httpserver.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
	logg.Info("calendar is running...")

	if err := grpcserver.Start(ctx); err != nil {
		logg.Error(err)
	}
}
