package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type User struct {
	Email     string `mapstructure:"email" json:"email"`
	Password  string `mapstructure:"password" json:"password"`
	GrantType string `mapstructure:"grant_type" json:"grant_type"`
}

type UserResponse struct {
	AccessToken string `json:"access_token"`
}

type UserToken struct {
	Email       string
	AccessToken string
}

//Login ...
func (user *User) Login(loginURL string) string {
	client := &http.Client{}

	payload, err := json.Marshal(user)

	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, loginURL, bytes.NewReader(payload))

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

	var userResponse UserResponse

	json.Unmarshal(body, &userResponse)

	return userResponse.AccessToken
}
