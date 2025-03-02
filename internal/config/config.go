package config

import (
	"flag"
	"os"
)

type Flags struct {
	RunAddress           string
	DatabaseURI          string
	AccrualSystemAddress string
}

func ParseFlags() *Flags {

	var flags Flags

	flag.StringVar(&flags.RunAddress, "a", "localhost:8080", "адрес и порт запуска сервиса")
	flag.StringVar(&flags.DatabaseURI, "d", "", "адрес подключения к базе данных")
	flag.StringVar(&flags.AccrualSystemAddress, "r", "", "адрес системы расчёта начислений")

	if envRunAddress := os.Getenv("RUN_ADDRESS"); envRunAddress != "" {
		flags.RunAddress = envRunAddress
	}

	if envDatabaseURI := os.Getenv("DATABASE_URI"); envDatabaseURI != "" {
		flags.DatabaseURI = envDatabaseURI
	}

	if envAccrualSystemAddress := os.Getenv("ACCRUAL_SYSTEM_ADDRESS"); envAccrualSystemAddress != "" {
		flags.AccrualSystemAddress = envAccrualSystemAddress
	}

	return &flags
}
