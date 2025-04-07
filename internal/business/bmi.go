package business

func CalculateBMI(height float64, weight float64) (float64, string) {
	if height <= 0 || weight <= 0 {
		return 0, "Invalid input. Height and weight must be positive numbers."
	}

	bmi := weight / (height * height)
	var category string

	switch {
	case bmi < 16.0:
		category = "Severe Thinness"
	case bmi < 17.0:
		category = "Moderate Thinness"
	case bmi < 18.5:
		category = "Mild Thinness"
	case bmi < 25.0:
		category = "Normal"
	case bmi < 30.0:
		category = "Overweight"
	case bmi < 35.0:
		category = "Obese (Class I)"
	case bmi < 40.0:
		category = "Obese (Class II)"
	default:
		category = "Obese (Class III)"
	}

	return bmi, category
}
