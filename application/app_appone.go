package application

import (
	"demo-app/domain_mydomain/controller/userapi"
	"demo-app/domain_mydomain/gateway/inmemory"
	"demo-app/domain_mydomain/usecase/registeruser"
	"demo-app/shared/driver"
	"demo-app/shared/infrastructure/config"
	"demo-app/shared/infrastructure/logger"
	"demo-app/shared/infrastructure/server"
	"demo-app/shared/util"
)

type appone struct {
	httpHandler *server.GinHTTPHandler
	controller  driver.Controller
}

func (c appone) RunApplication() {
	c.controller.RegisterRouter()
	c.httpHandler.RunApplication()
}

func NewAppone() func() driver.RegistryContract {

	const appName = "appone"

	return func() driver.RegistryContract {

		cfg := config.ReadConfig()

		appID := util.GenerateID(4)

		appData := driver.NewApplicationData(appName, appID)

		log := logger.NewSimpleJSONLogger(appData)

		httpHandler := server.NewGinHTTPHandler(log, cfg.Servers[appName].Address, appData)

		datasource := inmemory.NewGateway(log, appData, cfg)

		return &appone{
			httpHandler: &httpHandler,
			controller: &userapi.Controller{
				Log:                log,
				Config:             cfg,
				Router:             httpHandler.Router,
				RegisterUserInport: registeruser.NewUsecase(datasource),
			},
		}

	}
}
