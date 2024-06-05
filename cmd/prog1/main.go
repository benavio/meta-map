package main

import (
	"fmt"
	"log/slog"
	"meta-map/internal/fsdo"
	"meta-map/internal/lib/sl"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "Prod"
)

func main() {

	log := SetupLogger(envLocal)

	arr, err := fsdo.GetExif()
	if err != nil {
		log.Error("main", sl.Err(err))
	}

	fmt.Println(arr[1][0].Date)

}

// Init logger
func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
	}
	return log
}
