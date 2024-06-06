package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/benavio/meta-map.git/internal/cleints/telega"
	"github.com/benavio/meta-map.git/internal/fsdo"
	"github.com/benavio/meta-map.git/internal/lib/sl"
)

const (
	envLocal  = "local"
	envDev    = "dev"
	envProd   = "Prod"
	tgBotHost = "api.telegram.org"
)

func main() {

	log := SetupLogger(envLocal)
	tgClient := telega.New(MustToken(), tgBotHost)

	arr, err := fsdo.GetExif()
	if err != nil {
		log.Error("main", sl.Err(err))
	}

	fmt.Println(arr[1][0].Date)

}

func MustToken() string {
	token := flag.String("token-bot-token",
		"",
		"token fo access",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token broken")
	}
	return *token
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
