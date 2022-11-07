package userapi

import (
	"github.com/gin-gonic/gin"

	"demo-app/domain_mydomain/usecase/registeruser"
	"demo-app/shared/infrastructure/config"
	"demo-app/shared/infrastructure/logger"
)

type Controller struct {
	Router             gin.IRouter
	Config             *config.Config
	Log                logger.Logger
	RegisterUserInport registeruser.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.POST("/registeruser", r.authorized(), r.registerUserHandler())
	//r.Router.GET("/registeruser", r.authorized(), r.registerUserHandler())
}
