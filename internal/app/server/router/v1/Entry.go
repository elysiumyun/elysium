package v1

import (
	"github.com/elysiumyun/elysium/internal/app/server/controller"
	"github.com/gin-gonic/gin"
)

func Entry(engine *gin.RouterGroup) {
	apiRoutes := engine.Group(prefix + "/entry")
	apiRoutes.GET("/query", controller.EntryQuery)
}
