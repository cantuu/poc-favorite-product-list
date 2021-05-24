package controllers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

// SetupRouter build gin engine
func (server *Server) SetupRoutes() *gin.Engine {

	router := gin.New()

	data := router.Group("/api/product")

	data.GET("/health-check", HealthCheck)

	data.Use(
		apmgin.Middleware(router),
		CorsMiddleware(),
	)

	data.GET("/GetById/:id", server.GetProduct)
	data.GET("/GetAll", server.GetAllProducts)
	data.PUT("/:id", server.UpdateProduct)
	data.DELETE("/:id", server.DeleteProduct)
	data.POST("/", server.CreateProduct)

	return router
}

func CorsMiddleware() func(c *gin.Context) {
	allowedHeader := []string{
		"X-Application-Key",
		"Authorization",
		"X-Amz-Date",
		"X-Api-Key",
		"X-Amz-Security-Token",
		"X-Bifrost-Authorization",
	}

	config := cors.DefaultConfig()
	config.AddAllowHeaders(allowedHeader...)
	config.AddExposeHeaders("X-Request-Id")
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	return cors.New(config)
}

// Handler Simple health checker
// @Summary Check the application health
// @Description get health-check
// @Success 200
// @Failure 500
// @Failure default
// @Router /platform/consumer-registration/health-check [get]
func HealthCheck(context *gin.Context) {
	context.SecureJSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}
