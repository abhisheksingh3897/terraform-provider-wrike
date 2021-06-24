package token

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GenerateToken(client_id, client_secret, refresh_token string) (string, error) {

	url := "https://login.wrike.com/oauth2/token?client_id=" + client_id + "&client_secret=" + client_secret + "&grant_type=refresh_token&refresh_token=" + refresh_token + "&scope=Default, wsReadWrite"
	method := "POST"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Cookie", "wrikeLocale=en; isWhitelabel=false")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var tokenbody map[string]interface{}
	err = json.Unmarshal(body, &tokenbody)
	if err != nil {
		return "", err
	}
	token := tokenbody["token_type"].(string) + " " + tokenbody["access_token"].(string)
	return token, nil
}
