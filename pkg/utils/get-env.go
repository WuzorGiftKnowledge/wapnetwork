package utils

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func GetEnvVariable(key string) string {

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	return fmt.Sprint(viper.Get(key))
}
