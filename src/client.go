package src

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client interface {
	Signup(username, password string) (SignupResponse, error)
	EndGame(players Players) error
}

type client struct {
	http.Client
	baseUrl string
}

func (c *client) Signup(username, password string) (SignupResponse, error) {
	var result SignupResponse
	data := url.Values{
		"username": {username},
		"password": {password},
	}

	reqUrl, urlErr := url.Parse(c.baseUrl)
	if urlErr != nil {
		return result, urlErr
	}

	reqUrl.Path = "/api/v1/user/signup"

	res, err := c.Client.PostForm(reqUrl.String(), data)

	if err != nil {
		return result, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error(err)
		}
	}(res.Body)

	err = json.NewDecoder(res.Body).Decode(&result)

	if res.StatusCode != http.StatusCreated {
		return result, errors.New(result.Message)
	}

	return result, nil
}

func (c *client) EndGame(players Players) error {
	data, err := json.Marshal(PlayerEndGameRequest{Players: players})

	if err != nil {
		log.Fatal(err)
	}

	reqUrl, urlErr := url.Parse(c.baseUrl)
	if urlErr != nil {
		return urlErr
	}

	reqUrl.Path = "/api/v1/endgame"

	resp, err := c.Client.Post(reqUrl.String(), "application/json", bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return err
	}

	fmt.Println(res["status"])
	if res["status"] != "success" {
		fmt.Println(res)
	}

	return nil
}

func NewClient(baseUrl string, timeout time.Duration) *client {
	return &client{
		baseUrl: baseUrl,
		Client:  http.Client{Timeout: time.Second * timeout},
	}
}
