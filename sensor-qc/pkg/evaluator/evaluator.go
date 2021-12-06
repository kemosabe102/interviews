package evaluator

import (
	"fmt"
	"math"
)

func EvaluateSensorLog(parsedLogs ParsedLogFile) EvaluatedSensorOutput {
	resultsMap := make(map[string]string)

	humidityEvalCriteria := loadHumidityEvaluationConfig()
	thermometerEvalCriteria := loadThermometerEvaluationConfig()

	for _, v := range parsedLogs.Logs {
		switch v.SensorInfo.SensorType {
		case "humidity":
			resultsMap[v.SensorInfo.SensorName] = evaluateHumiditySensorLogs(parsedLogs.Reference, v, humidityEvalCriteria)
		case "thermometer":
			resultsMap[v.SensorInfo.SensorName] = evaluateThermometerSensorLogs(parsedLogs.Reference, v, thermometerEvalCriteria)
		}
	}

	return EvaluatedSensorOutput{
		SensorResults: resultsMap,
	}
}

// evaluateHumiditySensorLogs returns the classification of a humidity sensor based on whether its mean deviates more
// than the maximum allowed range from the reference point
func evaluateHumiditySensorLogs(referenceData ReferenceData, sensorData SensorQualityControlLogs, evalCriteria HumidityEvaluationCriteria) string {
	mean := sensorData.Mean()
	fmt.Println(sensorData.SensorInfo.SensorName, "Mean:", mean)

	if isMeanWithinTolerance(mean, referenceData.RelativeHumidity, evalCriteria.KeepMaxMeanToleranceFromReference) {
		return "keep"
	}
	return "discard"
}

// evaluateThermometerSensorLogs returns the classification of a thermometer sensor based on whether its mean deviates more
// than the maximum allowed range from the reference point and if the standard deviation of its readings is within
// the value passed in via the configuration file
func evaluateThermometerSensorLogs(referenceData ReferenceData, sensorData SensorQualityControlLogs, evalCriteria ThermometerEvaluationCriteria) string {
	mean := sensorData.Mean()
	stdDev := sensorData.StdDev()
	fmt.Println(sensorData.SensorInfo.SensorName, "Mean:", mean)

	if !(isMeanWithinTolerance(mean, referenceData.DegreesInFahrenheit, evalCriteria.MeanToleranceFromReference)) {
		return "precise"
	}

	switch {
	case stdDev < evalCriteria.UltraPreciseStdDev:
		return "ultra precise"
	case stdDev < evalCriteria.VeryPreciseStdDev:
		return "very precise"
	default:
		return "precise"
	}
}

func isMeanWithinTolerance(mean float64, reference float64, allowedTolerance float64) bool {
	upperMeanBounds := reference + allowedTolerance
	lowerMeanBounds := reference - allowedTolerance
	if mean > lowerMeanBounds && mean < upperMeanBounds {
		return true
	}
	return false
}

func (sensorData SensorQualityControlLogs) Mean() float64 {
	var sum,mean float64

	for _, v := range sensorData.SensorLogs {
		sum += v.Result
	}
	mean = sum/float64(len(sensorData.SensorLogs))

	return mean
}

func (sensorData SensorQualityControlLogs) StdDev() float64 {
	var mean,sumOfDiff,stdDev float64

	mean = sensorData.Mean()

	for _, v := range sensorData.SensorLogs {
		sumOfDiff += math.Pow(v.Result - mean, 2)
	}

	stdDev = math.Sqrt(sumOfDiff/float64(len(sensorData.SensorLogs)))

	fmt.Println(sensorData.SensorInfo.SensorName, "StdDev:", stdDev)

	return stdDev
}