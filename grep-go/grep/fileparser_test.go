package grep

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	tempReferenceDegree = 70.0
	humidityReference   = 45.0
)

func TestParseLogFile_Reference_Degrees(t *testing.T) {
	parsedLogs, err := SearchFile(logFilePath)
	assert.NoError(t, err)

	assert.Equalf(t, tempReferenceDegree, parsedLogs.Reference.DegreesInFahrenheit, "temp reference degree should be %v", tempReferenceDegree)
}

func TestParseLogFile_Reference_Humidity(t *testing.T) {
	parsedLogs, err := SearchFile(logFilePath)
	assert.NoError(t, err)

	assert.Equalf(t, humidityReference, parsedLogs.Reference.RelativeHumidity, "humidity reference should be %v", humidityReference)
}

func TestParseLogFile_SensorInfo(t *testing.T) {
	parsedLogs, err := SearchFile(logFilePath)
	assert.NoError(t, err)

	expectedSensorName := "temp-up"
	expectedSensorType := "thermometer"

	assert.Equalf(t, expectedSensorName, parsedLogs.Logs[2].SensorInfo.SensorName, "sensor name should be %v", expectedSensorName)
	assert.Equalf(t, expectedSensorType, parsedLogs.Logs[2].SensorInfo.SensorType, "sensor type should be %v", expectedSensorType)
}

func TestParseLogFile_SensorReadings(t *testing.T) {
	parsedLogs, err := SearchFile(logFilePath)
	assert.NoError(t, err)

	expectedSensorReadingCount := 10

	assert.Equalf(t, expectedSensorReadingCount, len(parsedLogs.Logs[2].SensorLogs), "there should be %v sensor readings", expectedSensorReadingCount)
}
