package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api")
	{
		apiGroup.POST("/bmi", BMICalculator)
	}
}
