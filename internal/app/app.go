package app

import (
	"log/slog"
	grpcapp "sso_pet/internal/app/grpc"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, port int, storagePath string, tokenTTL time.Duration) *App {
	// TODO: инициализировать хранилище

	// TODO: инициализировать сервисный слой

	grpcApp := grpcapp.New(log, port)

	return &App{
		GRPCSrv: grpcApp,
	}
}
