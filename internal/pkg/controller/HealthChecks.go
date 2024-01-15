package controller

import (
	"net/http"
	"time"

	"github.com/elysiumyun/elysium/internal/pkg/system/service"
	"github.com/elysiumyun/elysium/pkg/info"
	"github.com/elysiumyun/elysium/pkg/timezone"
	"github.com/gin-gonic/gin"
)

func HealthChecks(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"service":      info.MicroServiceName,
			"timestamp":    timezone.Format(time.Now()),
			"dependencies": service.Dependencies.Get(),
		},
	)
}
