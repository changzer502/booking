package config

import (
	"github.com/goccy/go-json"
	"io/ioutil"
	"net/http"
)

type Wx struct {
	AppID     string `mapstructure:"app_id" json:"app_id" yaml:"app_id"`
	AppSecret string `mapstructure:"app_secret" json:"app_secret" yaml:"app_secret"`
}

func (wxApi *Wx) GetOpenId(code string) (openIdRes *OpenIdRes, err error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + wxApi.AppID +
		"&secret=" + wxApi.AppSecret +
		"&js_code=" + code + "&grant_type=authorization_code"
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &openIdRes)
	if err != nil {
		return
	}
	return
}

func (wxApi *Wx) GetAccessToken() (accessTokenRes *AccessTokenRes, err error) {
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential" +
		"&appid=" + wxApi.AppID +
		"&secret=" + wxApi.AppSecret
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &accessTokenRes)
	if err != nil {
		return
	}
	return
}

type OpenIdRes struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	Errcode    int64  `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

type AccessTokenRes struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}
