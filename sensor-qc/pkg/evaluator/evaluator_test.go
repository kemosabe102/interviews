package evaluator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	reference        = 1.0
	allowedTolerance = 0.5
	goodMean         = 1.1
	falseMean        = 1.6

	logFilePath = "testLogs.txt"
)

var (
	thermometerEvalCriteria = ThermometerEvaluationCriteria{
		UltraPreciseStdDev: 3.0,
		VeryPreciseStdDev: 5.0,
		MeanToleranceFromReference: 0.5,
	}

	humidityEvalCriteria = HumidityEvaluationCriteria{
		KeepMaxMeanToleranceFromReference: 1.0,
	}
)

func TestIsMeanWithinTolerance_ShouldReturnTrue(t *testing.T) {
	assert.Truef(t, isMeanWithinTolerance(goodMean, reference, allowedTolerance), "should return true")
}

func TestIsMeanWithinTolerance_ShouldReturnFalse(t *testing.T) {
	assert.Falsef(t, isMeanWithinTolerance(falseMean, reference, allowedTolerance), "should return false")
}

func TestEvaluateThermometerSensorLogs_UltraPrecise(t *testing.T) {
	// assert
	parsedLogs, _ := ParseLogFile(logFilePath)
	sensorData := SensorQualityControlLogs{}
	for _, v := range parsedLogs.Logs {
		if v.SensorInfo.SensorName == "temp-up" {
			sensorData = v
		}
	}

	// act
	sensorClassification := evaluateThermometerSensorLogs(parsedLogs.Reference, sensorData, thermometerEvalCriteria)

	// assert
	assert.Equalf(t, "ultra precise", sensorClassification,
		"sensor %v with stdDev of %v and mean of %v should be ultra precise and test data of %v", sensorData.SensorInfo.SensorName, sensorData.StdDev(), sensorData.Mean())
}

func TestEvaluateThermometerSensorLogs_VeryPrecise(t *testing.T) {
	// assert
	parsedLogs, _ := ParseLogFile(logFilePath)
	sensorData := SensorQualityControlLogs{}
	for _, v := range parsedLogs.Logs {
		if v.SensorInfo.SensorName == "temp-vp" {
			sensorData = v
		}
	}

	// act
	sensorClassification := evaluateThermometerSensorLogs(parsedLogs.Reference, sensorData, thermometerEvalCriteria)

	// assert
	assert.Equalf(t, "very precise", sensorClassification,
		"sensor %v with stdDev of %v and mean of %v should be very precise", sensorData.SensorInfo.SensorName, sensorData.StdDev(), sensorData.Mean())
}

func TestEvaluateThermometerSensorLogs_Precise(t *testing.T) {
	// assert
	parsedLogs, _ := ParseLogFile(logFilePath)
	sensorData := SensorQualityControlLogs{}
	for _, v := range parsedLogs.Logs {
		if v.SensorInfo.SensorName == "temp-p" {
			sensorData = v
		}
	}

	// act
	sensorClassification := evaluateThermometerSensorLogs(parsedLogs.Reference, sensorData, thermometerEvalCriteria)

	// assert
	assert.Equalf(t, "precise", sensorClassification,
		"sensor %v with stdDev of %v and mean of %v should be precise", sensorData.SensorInfo.SensorName, sensorData.StdDev(), sensorData.Mean())
}

func TestEvaluateHumiditySensorLogs_Keep(t *testing.T) {
	// assert
	parsedLogs, _ := ParseLogFile(logFilePath)
	sensorData := SensorQualityControlLogs{}
	for _, v := range parsedLogs.Logs {
		if v.SensorInfo.SensorName == "hum-keep" {
			sensorData = v
		}
	}

	// act
	sensorClassification := evaluateHumiditySensorLogs(parsedLogs.Reference, sensorData, humidityEvalCriteria)

	// assert
	assert.Equalf(t, "keep", sensorClassification,
		"sensor %v with mean of %v should be kept", sensorData.SensorInfo.SensorName, sensorData.Mean())
}

func TestEvaluateHumiditySensorLogs_Discard(t *testing.T) {
	// assert
	parsedLogs, _ := ParseLogFile(logFilePath)
	sensorData := SensorQualityControlLogs{}
	for _, v := range parsedLogs.Logs {
		if v.SensorInfo.SensorName == "hum-discard" {
			sensorData = v
		}
	}

	// act
	sensorClassification := evaluateHumiditySensorLogs(parsedLogs.Reference, sensorData, humidityEvalCriteria)

	// assert
	assert.Equalf(t, "discard", sensorClassification,
		"sensor %v with mean of %v should be discarded", sensorData.SensorInfo.SensorName, sensorData.Mean())
}
