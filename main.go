/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"cash-register-svc/cmd"
	"cash-register-svc/internal/app"
	"cash-register-svc/internal/infrastructure"
	"errors"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	ProductsSourceFile string // `mapstructure:"products_source_file"`
	DiscountRules      map[string]string
}

func main() {
	appConfig, err := loadAppConfig()

	productRepository := infrastructure.NewProductFileRepository(appConfig.ProductsSourceFile)
	productCache := infrastructure.NewProductsInMemoryCache()

	cashRegisterApp, err := app.NewCashRegisterApp(productRepository, productCache, appConfig.DiscountRules)
	if err != nil {
		log.Fatal(fmt.Errorf("app could not be initialized: %w", err))
	}

	cmd.Execute()
}

func loadAppConfig() (AppConfig, error) {
	var appConfig AppConfig

	viper.SetConfigName("app-config") // Register config file name (no extension)
	viper.SetConfigType("json")       // Look for specific type
	err := viper.ReadInConfig()
	if err != nil {
		return appConfig, err
	}

	err = viper.Unmarshal(&appConfig)
	if err != nil {
		return appConfig, err
	}

	if len(appConfig.ProductsSourceFile) == 0 {
		return AppConfig{}, errors.New("product source file cannot be empty")
	}

	return appConfig, nil
}
