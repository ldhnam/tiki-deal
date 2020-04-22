package main

type Checkout struct {
	Payment Payment `mapstructure:"payment" json:"payment"`
}

type Payment struct {
	Method string `mapstructure:"method" json:"method"`
}
