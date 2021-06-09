package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	authToken  string
	httpClient *http.Client
}

type User struct {
	ID        string        `json:"id"`
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	Profile   []UserProfile `json:"profiles"`
}

type UserProfile struct {
	AccountID string `json:"accountId"`
	Email     string `json:"email"`
}

type UserGet struct {
	Data []User `json:"data"`
}

func NewClient(token string) *Client {
	return &Client{
		authToken:  token,
		httpClient: &http.Client{},
	}
}

func (c *Client) GetId(email string) (string, error) {
	body, err := c.HttpRequest(fmt.Sprintf("contacts"), "GET", &strings.Reader{})
	if err != nil {
		return "", err
	}
	getuser := UserGet{}
	err = json.Unmarshal(body, &getuser)
	if err != nil {
		return "", err
	}
	for _, v := range getuser.Data {
		if v.Profile[0].Email == email {
			return v.ID, err
		}
	}
	return "", fmt.Errorf("user not found")
}

func (c *Client) GetUser(email string) (*User, error) {
	userid, err := c.GetId(email)
	if err != nil {
		return nil, err
	}
	body, err := c.HttpRequest(fmt.Sprintf("users/%v", userid), "GET", &strings.Reader{})
	if err != nil {
		return nil, err
	}
	user := UserGet{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}
	return &user.Data[0], nil
}

func (c *Client) NewUser(email string) error {
	_, err := c.GetUser(email)
	if err == nil {
		return fmt.Errorf("user already exist")
	}
	parms := url.Values{}
	parms.Add("email", email)
	body := strings.NewReader(parms.Encode())
	_, err = c.HttpRequest("invitations", "POST", body)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) HttpRequest(path, method string, body *strings.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, c.requestPath(path), body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", c.authToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respbody, nil
}

func (c *Client) requestPath(path string) string {
	return fmt.Sprintf("https://www.wrike.com/api/v4/%v", path)
}
