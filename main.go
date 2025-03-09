package main

import (
	"github.com/maryakotova/gophermart/internal/config"
	// "github.com/maryakotova/gophermart/internal/logger"
)

func main() {

	flags, err := config.ParseFlags()
	if err != nil {
		panic(err)
	}

	// log, err := logger.Initialize()
	// if err != nil {
	// 	panic(err)
	// }

}
