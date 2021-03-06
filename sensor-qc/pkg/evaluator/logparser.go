package evaluator

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// ParseLogFile scans each line of the log file then records reference data and parses the sensor's info and readings
func ParseLogFile(logFilePath string) (ParsedLogFile, error) {
	file, err := os.Open(logFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var parsedLogs ParsedLogFile
	var sensorInfoAndLogs SensorQualityControlLogs

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineContent := strings.Split(scanner.Text(), " ")

		switch {
		case strings.Contains(scanner.Text(), "reference"):
			parsedLogs.Reference = parseReferenceData(lineContent)
		case strings.Contains(scanner.Text(), "thermometer") || strings.Contains(scanner.Text(), "humidity"):
			if len(sensorInfoAndLogs.SensorInfo.SensorName) > 0 {
				parsedLogs.Logs = append(parsedLogs.Logs, sensorInfoAndLogs)
			}
			sensorInfoAndLogs = SensorQualityControlLogs{}
			sensorInfoAndLogs.SensorInfo = parseSensorInfo(lineContent)
		default:
			sensorInfoAndLogs.SensorLogs = append(sensorInfoAndLogs.SensorLogs, parseSensorReading(lineContent))
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	parsedLogs.Logs = append(parsedLogs.Logs, sensorInfoAndLogs)

	return parsedLogs, nil
}

func parseReferenceData(line []string) ReferenceData {
	degrees, err := strconv.ParseFloat(line[1], 64)
	if err != nil {
		log.Println(err)
	}

	humidity, err := strconv.ParseFloat(line[2], 64)
	if err != nil {
		log.Println(err)
	}

	return ReferenceData{
		DegreesInFahrenheit: degrees,
		RelativeHumidity:    humidity,
	}
}

func parseSensorInfo(line []string) SensorInformation {
	return SensorInformation{
		SensorType: line[0],
		SensorName: line[1],
	}
}

func parseSensorReading(line []string) SensorReading {
	reading, err := strconv.ParseFloat(line[1], 64)
	if err != nil {
		log.Println(err)
	}

	return SensorReading{
		DateTime:  line[0],
		Result:    reading,
	}
}
