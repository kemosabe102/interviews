package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func NewConfig(configFilePath string) {
	viper.SetConfigName("config")
	viper.AddConfigPath(configFilePath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	ConfigOption("thermometer_sensor_classification.max_std_dev.ultra_precise", 3.0)
	ConfigOption("thermometer_sensor_classification.max_std_dev.very_precise", 5.0)
	ConfigOption("thermometer_sensor_classification.mean_range_from_reference", 0.5)

	ConfigOption("humidity_sensor_classification.max_percent_difference.keep", 1.0)

	fileOption(ConfigOption("log_file_path", "./pkg/evaluator/testLogs.txt"))
}

func ConfigOption(key string, defaultValue interface{}) string {
	viper.SetDefault(key, defaultValue)

	return key
}

// Asserts that the chosen value exists on the local file system by panicking if it doesn't
func fileOption(key string) {
	chosenValue := viper.GetString(key)

	if _, err := os.Stat(chosenValue); err != nil {
		panic(fmt.Errorf("chosen option %s does not exist", chosenValue))
	}
}