package business_test

import (
	"testing"

	"github.com/rezacloner1372/hesabketab/internal/business"
)

func TestCalculateBMI(t *testing.T) {
	height := 1.75
	weight := 70.0

	expected := weight / (height * height)
	bmi, _ := business.CalculateBMI(height, weight)
	if bmi != expected {
		t.Errorf("Expected BMI: %f, got: %f", expected, bmi)
	}
}
