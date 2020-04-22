package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const SHIPPING_PLAN_STANDARD = 1

type Item struct {
	Products []Product `mapstructure:"products" json:"products"`
}

type Product struct {
	ProductID int `mapstructure:"product_id" json:"product_id"`
	Quantity  int `mapstructure:"quantity" json:"qty"`
}

type Cart struct {
	Items           []CartItem      `json:"items"`
	ShippingPlan    ShippingPlan    `json:"shipping_plan"`
	SubTotal        int             `json:"subtotal"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
	PaymentMethod   PaymentMethod   `json:"payment_method"`
}

type ShippingAddress struct {
	Street string `json:"street"`
}

type PaymentMethod struct {
	Method string `json:"method"`
}

type CartItem struct {
	ID          string `json:"id"`
	ProductName string `json:"product_name"`
}

type ShippingPlan struct {
	ID int64 `json:"id"`
}

type Deal struct {
	IsCheck bool `mapstructure:"isCheck"`
	Price   int  `mapstructure:"price"`
}

// AddToCart ...
func AddToCart(accessToken, url string, item Item) {
	client := &http.Client{}

	payload, err := json.Marshal(item)

	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-access-token", accessToken)

	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func GetCart(accessToken, url string) Cart {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		panic(err)
	}
	req.Header.Set("x-access-token", accessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var cart Cart
	json.Unmarshal(body, &cart)
	return cart
}

// DeleteCart ...
func (cart *Cart) DeleteItems(accessToken, url string) {
	client := http.Client{}

	for _, item := range cart.Items {
		fmt.Println(item)
		deleteURL := fmt.Sprintf(url, item.ID)
		fmt.Println(deleteURL)

		req, err := http.NewRequest(http.MethodDelete, deleteURL, nil)

		if err != nil {
			panic(err)
		}
		req.Header.Set("x-access-token", accessToken)

		resp, err := client.Do(req)

		if err != nil {
			panic(err)
		}

		fmt.Println(resp.Body)
	}
}

func (cart *Cart) PutShippingPlan(accessToken, url string, shippingPlanID int64) {
	client := &http.Client{}

	shippingPlanURL := fmt.Sprintf(url, shippingPlanID)

	req, err := http.NewRequest(http.MethodPut, shippingPlanURL, nil)

	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-access-token", accessToken)

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Body)
}

func (cart *Cart) Checkout(accessToken, url string, checkout Checkout) {
	client := &http.Client{}

	payload, err := json.Marshal(checkout)

	fmt.Println(string(payload))

	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))

	if err != nil {
		panic(err)
	}
	req.Header.Set("x-access-token", accessToken)
	req.Header.Add("User-Agent", "Tiki/vn.app.tiki(4.43.0; Android Version 28 9)")

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Body)
}

func (cart *Cart) GetShippingPlans(accessToken, url string) int64 {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		panic(err)
	}
	req.Header.Set("x-access-token", accessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var shippingPlans []ShippingPlan

	json.Unmarshal(body, &shippingPlans)

	for _, shippingPlan := range shippingPlans {
		if shippingPlan.ID == 3 {
			return shippingPlan.ID
		}
		if shippingPlan.ID == 5 {
			return shippingPlan.ID
		}

	}

	return SHIPPING_PLAN_STANDARD
}

func (cart *Cart) SelectPaymentMethod(accessToken, url string) {
	req, err := http.NewRequest(http.MethodPut, url, nil)
	req.Header.Set("x-access-token", accessToken)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
