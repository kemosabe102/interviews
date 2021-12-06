package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"sensor-qc/pkg/evaluator"
)

const (
	defaultConfigFilePath = "./"
	configFilePathUsage   = "config file directory (eg. '/app/sensor-qc/'). Config file must be named 'config.yml'."
)

var (
	configFilePath string
)

func init() {
	flag.StringVar(&configFilePath, "conf", defaultConfigFilePath, configFilePathUsage)

	flag.Parse()
}

func main() {
	viper.AutomaticEnv()
	logFilePath := viper.GetString("log_file_path")

	NewConfig(configFilePath)

	parsedLogs, err := evaluator.ParseLogFile(logFilePath)
	if err != nil {
		log.Fatal(err)
	}

	results := evaluator.EvaluateSensorLog(parsedLogs)
	fmt.Println(results.SensorResults)
}
