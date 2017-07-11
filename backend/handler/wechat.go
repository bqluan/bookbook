package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type codeToOpenIDResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
}

func (h *apiHandler) CodeToOpenID(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	state := r.FormValue("state")

	resp, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=wxbf0bffa50d1ce293&secret=5d652482179ffbebd9fed0edec9b2edd&code=%s&grant_type=authorization_code", code))
	if err != nil {
		http.Error(w, fmt.Sprintf("handler: %s", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	var c codeToOpenIDResponse
	if err := dec.Decode(&c); err != nil {
		http.Error(w, fmt.Sprintf("handler: %s", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/#/book/%s?openid=%s", state, c.OpenID), http.StatusFound)
}
