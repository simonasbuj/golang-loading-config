package main

import (
	"golang-loading-config/config"
	"log/slog"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	//ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGKILL, syscall.SIGTERM)
	//defer stop()

	slog.Info("Let's load config")
	cfg, err := loadConfig("config/local.yml")
	if err != nil {
		slog.Error("error while loading config. stopping the app", "err", err.Error())
		os.Exit(1)
	}

	slog.Info("loaded config", "config", cfg)

}

func loadConfig(configPath string) (config.Config, error) {
	var cfg config.Config

	// first load from yaml, but only if yaml exists
	if _, err := os.Stat(configPath); err != nil {
		slog.Warn("config yaml is missing, skipping loading from yaml", "path", configPath)
	} else {
		if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
			return config.Config{}, err
		}
	}

	// now load from ENV and override defaults
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}
