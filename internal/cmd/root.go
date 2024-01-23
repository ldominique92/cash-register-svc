package cmd

import (
	"cash-register-svc/internal/app"
	"cash-register-svc/internal/infrastructure"
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type AppConfig struct {
	ProductsSourceFile string            `mapstructure:"products_source_file"`
	DiscountRules      map[string]string `mapstructure:"discount_rules"`
}

var RootCmd = &cobra.Command{
	Use:   "cash-register-svc",
	Short: "###  Cash Register Amenitiz ###",
	Long:  `###  Cash Register Amenitiz ###`,
}

var cashRegisterApp *app.CashRegisterApp

func init() {
	appConfig, err := loadAppConfig()

	productRepository := infrastructure.NewProductFileRepository(appConfig.ProductsSourceFile)
	productCache := infrastructure.NewProductsInMemoryCache()

	newCashRegisterApp, err := app.NewCashRegisterApp(productRepository, productCache, appConfig.DiscountRules)
	if err != nil {
		log.Fatal(fmt.Errorf("newCashRegisterApp could not be initialized: %w", err))
	}
	cashRegisterApp = &newCashRegisterApp
}

func loadAppConfig() (AppConfig, error) {
	var appConfig AppConfig

	viper.AddConfigPath(".")
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
