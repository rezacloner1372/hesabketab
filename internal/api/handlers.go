package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rezacloner1372/hesabketab/internal/business"
)

type BMIRequest struct {
	Height float64 `json:"height" binding:"required"` // in meters
	Weight float64 `json:"weight" binding:"required"` // in kilograms
}

func BMICalculator(c *gin.Context) {
	var req BMIRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input. Please provide height(m) and weight(kg)."})
		return
	}

	bmiValue, category := business.CalculateBMI(req.Height, req.Weight)

	healthTip := "Maintain your current lifestyle and consult a healthcare provider for personalized advice."
	c.JSON(http.StatusOK, gin.H{
		"bmi":       bmiValue,
		"category":  category,
		"healthTip": healthTip,
	})
}
