package main

import (
	"github.com/maryakotova/metrics/internal/config"
	"github.com/maryakotova/metrics/internal/logger"
)

func main() {

	flags, err := config.ParseFlags()
	if err != nil {
		panic(err)
	}

	log, err := logger.Initialize("")
	if err != nil {
		panic(err)
	}

}
