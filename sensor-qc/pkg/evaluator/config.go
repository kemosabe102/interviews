package evaluator

import "github.com/spf13/viper"

func loadThermometerEvaluationConfig() ThermometerEvaluationCriteria {
	thermometerStdDevEvaluationCriteria := viper.Sub("thermometer_sensor_classification.max_std_dev")
	if thermometerStdDevEvaluationCriteria == nil {
		panic("temperature evaluation configuration not found")
	}

	return ThermometerEvaluationCriteria{
		UltraPreciseStdDev: thermometerStdDevEvaluationCriteria.GetFloat64("ultra_precise"),
		VeryPreciseStdDev: thermometerStdDevEvaluationCriteria.GetFloat64("very_precise"),
		MeanToleranceFromReference: viper.GetFloat64("thermometer_sensor_classification.mean_range_from_reference"),
	}
}

func loadHumidityEvaluationConfig() HumidityEvaluationCriteria {
	humidityEvaluationCriteria := viper.Sub("humidity_sensor_classification.max_percent_difference")
	if humidityEvaluationCriteria == nil {
		panic("humidity evaluation configuration not found")
	}

	return HumidityEvaluationCriteria{
		KeepMaxMeanToleranceFromReference: humidityEvaluationCriteria.GetFloat64("keep"),
	}
}
