package userapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"demo-app/domain_mydomain/usecase/registeruser"
	"demo-app/shared/infrastructure/logger"
	"demo-app/shared/model/payload"
	"demo-app/shared/util"
)

// registerUserHandler ...
func (r *Controller) registerUserHandler() gin.HandlerFunc {

	type request struct {
		Username string `json:"username"`
	}

	type response struct {
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		if err := c.BindJSON(&jsonReq); err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req registeruser.InportRequest
		req.Username = jsonReq.Username

		r.Log.Info(ctx, util.MustJSON(req))

		res, err := r.RegisterUserInport.Execute(ctx, req)
		if err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		_ = res

		r.Log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
