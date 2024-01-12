package main

import (
	"log/slog"
	"os"
	"os/signal"
	"sso_pet/internal/app"
	"sso_pet/internal/config"
	"sso_pet/internal/lib/logger/pretty"
	"syscall"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)
	log.Info("starting app", slog.Any("config", cfg))

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTtl)

	go application.GRPCSrv.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	flag := <-stop

	application.GRPCSrv.Stop(flag.String())
	log.Info("application stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettyLogger()
	case envProd:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettyLogger() *slog.Logger {
	opts := pretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
