package main

import "github.com/spf13/viper"

// Config ...
type Config struct {
	LoginURL               string   `mapstructure:"login_url"`
	Users                  []User   `mapstructure:"users"`
	Item                   Item     `mapstructure:"item"`
	AddToCartURL           string   `mapstructure:"add_to_cart_url"`
	GetCartURL             string   `mapstructure:"get_cart_url"`
	DeleteItemURL          string   `mapstructure:"delete_item_url"`
	PutShippingPlanURL     string   `mapstructure:"put_shipping_plan_url"`
	GetShippingPlansURL    string   `mapstructure:"get_shipping_plans_url"`
	SelectPaymentMethodURL string   `mapstructure:"select_payment_method_url"`
	CheckoutURL            string   `mapstructure:"checkout_url"`
	Checkout               Checkout `mapstructure:"checkout"`
	OrderQuantity          int      `mapstructure:"order_quantity"`
	Deal                   Deal     `mapstructure:"deal"`
}

// ReadConfig ...
func ReadConfig() Config {
	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	var config Config

	err = viper.Unmarshal(&config)

	if err != nil {
		panic(err)
	}

	return config
}
