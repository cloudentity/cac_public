package cac

import (
	"github.com/cloudentity/cac/internal/cac/client"
	"github.com/cloudentity/cac/internal/cac/config"
	"github.com/cloudentity/cac/internal/cac/logging"
	"github.com/cloudentity/cac/internal/cac/storage"
	"golang.org/x/exp/slog"
)

type Application struct {
	Config  *config.Configuration
	Client  *client.Client
	Storage storage.Storage
}

func InitApp(configPath string) (app *Application, err error) {
	app = &Application{}

	if app.Config, err = config.InitConfig(configPath); err != nil {
		return app, err
	}

	if err = logging.InitLogging(app.Config.Logging); err != nil {
		return app, err
	}

	slog.Debug("config", "c", app.Config.Client)

	if app.Client, err = client.InitClient(app.Config.Client); err != nil {
		return app, err
	}

	if app.Storage, err = storage.InitMultiStorage(app.Config.Storage); err != nil {
		return app, err
	}

	slog.Info("Initiated application")

	return app, nil
}
