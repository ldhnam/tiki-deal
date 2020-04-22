package main

import (
	"fmt"
	"sync"
)

func main() {
	config := ReadConfig()

	users := config.Users

	if users == nil {
		fmt.Println("Users are empty")
	}

	var wg sync.WaitGroup
	var accessTokens []string

	for _, user := range users {
		accessToken := user.Login(config.LoginURL)
		accessTokens = append(accessTokens, accessToken)
	}

	for _, accessToken := range accessTokens {
		preCart := GetCart(accessToken, config.GetCartURL)
		preCart.DeleteItems(accessToken, config.DeleteItemURL)
	}

	for _, accessToken := range accessTokens {
		AddToCart(accessToken, config.AddToCartURL, config.Item)
	}

	for i := 0; i < config.OrderQuantity; i++ {
		for _, accessToken := range accessTokens {
			wg.Add(1)
			go handle(accessToken, config, &wg)

			wg.Wait()
		}
	}
}

func handle(accessToken string, config Config, wg *sync.WaitGroup) {
	defer wg.Done()

	cart := GetCart(accessToken, config.GetCartURL)
	cart.PutShippingPlan(accessToken, config.PutShippingPlanURL, 1)
	cart.SelectPaymentMethod(accessToken, config.SelectPaymentMethodURL)
	if config.Deal.IsCheck {
		for {
			cart := GetCart(accessToken, config.GetCartURL)
			fmt.Println(fmt.Sprintf("sub total = %d, payment method = %s, shipping address = %s, shipping plan = %d", cart.SubTotal, cart.PaymentMethod.Method, cart.ShippingAddress.Street, cart.ShippingPlan.ID))
			if cart.PaymentMethod.Method == "" {
				cart.SelectPaymentMethod(accessToken, config.SelectPaymentMethodURL)
			}
			if cart.SubTotal == config.Deal.Price {
				cart.Checkout(accessToken, config.CheckoutURL, config.Checkout)
				break
			}
		}
	} else {
		cart.Checkout(accessToken, config.CheckoutURL, config.Checkout)
	}
}
