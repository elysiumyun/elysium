package router

import (
	v1 "github.com/elysiumyun/elysium/internal/app/server/router/v1"
	"github.com/elysiumyun/elysium/internal/pkg/controller"
	"github.com/gin-gonic/gin"
)

// define root routes
func rootRoutes(engine *gin.Engine, context string) {
	apiRoutes := engine.Group("/")
	apiRoutes.GET("health", controller.HealthChecks)
}

// registry all routes
func SetRouter(engine *gin.Engine, context string) {
	rootRoutes(engine, context)

	// service context
	subApiRoutes := engine.Group(context)

	// api ver v1
	v1.Entry(subApiRoutes)
}
