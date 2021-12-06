package main

import (
	"math/rand"
	"strconv"
	"strings"
)

// createFloatString can be used to create a random sample of test data for sensor readings
func createFloatString() {
	var valuesToString []string
	values := randFloats(40, 50, 5)

	for _, float := range values{
		valuesToString = append(valuesToString, strconv.FormatFloat(float, 'f', 1, 64) )
	}

	println(strings.Join(valuesToString, ","))
}

func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64() * (max - min)
	}
	return res
}