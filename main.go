package main

import (
	"fmt"
	"nicetry/server"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	v := viper.New()
	v.AddConfigPath(".")
	v.AutomaticEnv()
	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	var C server.Config

	err = v.Unmarshal(&C)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v\n", err)
	} else {
		fmt.Println("config:", C)
	}
	app := fiber.New()
	server.SetupRoutes(app, C)

	app.Listen(fmt.Sprintf(":%s", C.APP_PORT))
}
