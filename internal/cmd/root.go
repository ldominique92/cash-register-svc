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
	ProductsSourceFile string `mapstructure:"products_source_file"`
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

	newCashRegisterApp, err := app.NewCashRegisterApp(productRepository, productCache)
	if err != nil {
		log.Fatal(fmt.Errorf("newCashRegisterApp could not be initialized: %w", err))
	}
	cashRegisterApp = &newCashRegisterApp
}

func loadAppConfig() (AppConfig, error) {
	var appConfig AppConfig

	viper.AddConfigPath(".")
	viper.SetConfigName("app-config")
	viper.SetConfigType("json")
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
