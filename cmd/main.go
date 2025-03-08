package main

import (
	"context"
	"fmt"
	"github.com/alishojaeiir/GoShort/config"
	"github.com/alishojaeiir/GoShort/internal"
	cfgloader "github.com/alishojaeiir/GoShort/pkg/cfg_loader"
	"os"
	"path/filepath"
)

func main() {
	var cfg config.Config

	// Load config
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v", err)
	}

	options := cfgloader.Option{
		Prefix:       "GOSH_",
		Delimiter:    ".",
		Separator:    "__",
		YamlFilePath: filepath.Join(workingDir, "config", "config-local.yml"),
		CallbackEnv:  nil,
	}

	if err := cfgloader.Load(options, &cfg); err != nil {
		fmt.Sprintln("Failed to load config: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := internal.Setup(ctx, cfg)
	app.Start()
}
