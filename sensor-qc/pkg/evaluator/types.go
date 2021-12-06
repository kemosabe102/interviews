package evaluator

// ParsedLogFile is a collection of the parsed logs
type ParsedLogFile struct {
	Reference ReferenceData
	Logs      []SensorQualityControlLogs
}

// ReferenceData contains the reference information used to evaluate device readings
type ReferenceData struct {
	DegreesInFahrenheit float64
	RelativeHumidity    float64
}

// SensorQualityControlLogs contains the device logs from the quality control test
type SensorQualityControlLogs struct {
	SensorInfo SensorInformation
	SensorLogs []SensorReadings
}

// SensorInformation contains details about the device
type SensorInformation struct {
	SensorType string
	SensorName string
}

// SensorReadings is a single device log that contains the datetime and a reading from the device
type SensorReadings struct {
	DateTime  string
	Result    float64
}

// ThermometerEvaluationCriteria contains the information used to classify a thermometer
type ThermometerEvaluationCriteria struct {
	UltraPreciseStdDev float64
	VeryPreciseStdDev float64
	MeanToleranceFromReference float64
}

// HumidityEvaluationCriteria contains the information used to classify a humidity sensor
type HumidityEvaluationCriteria struct {
	KeepMaxMeanToleranceFromReference float64
}

// EvaluatedSensorOutput contains the results of the evaluated sensors
type EvaluatedSensorOutput struct {
	SensorResults map[string]string
}
