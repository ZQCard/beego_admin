package handlers

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type JSApiAccessTokenResponse struct {
	AccessToken string        `json:"access_token"`
	ExpiresIn   time.Duration `json:"expires_in"`
	ErrCode     int           `json:"errcode"`
	ErrMsg      string        `json:"errmsg"`
}

// 网络请求微信服务器获取access_token,用于js_api
func JsApiGetAccessToken(appId, appSecret string) (accessToken string, expire time.Duration, err error) {
	client := http.Client{}
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appId + "&secret=" + appSecret
	// 提交请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", 0, err
	}

	// 处理返回结果
	response, _ := client.Do(request)
	if response.StatusCode != 200 {
		return "", 0, errors.New("微信服务器返回结果失败")
	}

	body, _ := ioutil.ReadAll(response.Body)
	var r = JSApiAccessTokenResponse{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return "", 0, errors.New("accessToken:json解析失败")
	}

	if r.ErrCode != 0 {
		return "", 0, errors.New("accessToken:" + r.ErrMsg)
	}

	return r.AccessToken, r.ExpiresIn, nil
}

type JsApiTicket struct {
	ErrCode   int           `json:"errcode"`
	ErrMsg    string        `json:"errmsg"`
	Ticket    string        `json:"ticket"`
	ExpiresIn time.Duration `json:"expires_in"`
}

// 网络请求微信服务器获取jsapi_ticket
func GetJsApiTicket(accessToken string) (ticket string, expire time.Duration, err error) {
	client := &http.Client{}
	url := "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=" + accessToken + "&type=jsapi"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", 0, err
	}
	response, _ := client.Do(request)
	if response.StatusCode != 200 {
		return "", 0, errors.New("微信服务器返回结果失败")
	}
	var r = JsApiTicket{}
	body, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &r)
	if r.ErrCode != 0 {
		return "", 0, errors.New("accessToken:" + r.ErrMsg)
	}
	return r.Ticket, r.ExpiresIn, nil
}

// 微信js signature签名算法
func GetSignature(ticket, noncestr, url string, timestamp int64) string {
	str := "jsapi_ticket=" + ticket + "&noncestr=" + noncestr + "&timestamp=" + strconv.FormatInt(timestamp, 10) + "&url=" + url
	return sha1String(str)
}

// sha1加密字符串
func sha1String(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte(nil)))
}

type GrantAccessTokenResponse struct {
	AccessToken  string        `json:"access_token"`
	ExpiresIn    time.Duration `json:"expires_in"`
	RefreshToken string        `json:"refresh_token"`
	OpenId       string        `json:"openid"`
	Scope        string        `json:"scope"`
	ErrCode      int           `json:"errcode"`
	ErrMsg       string        `json:"errmsg"`
}

// 网页授权获取access_token
func GrantGetAccessToken(appId, appSecret, code string) (accessToken string, expire time.Duration, openId string, err error) {
	client := http.Client{}
	url := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + appId + "&secret=" + appSecret + "&code=" + code + "&grant_type=authorization_code"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", 0, "", err
	}
	response, err := client.Do(request)
	if err != nil {
		return "", 0, "", err
	}
	res := GrantAccessTokenResponse{}
	body, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &res)

	if err != nil {
		return "", 0, "", errors.New("get grant access token: JSON 解析失败")
	}

	if res.ErrCode != 0 {
		return "", 0, "", errors.New("get grant access token 失败: " + res.ErrMsg)
	}

	return res.AccessToken, res.ExpiresIn, res.OpenId, nil
}

type WeChatUserInfo struct {
	OpenId string `json:"openid"`
	Nickname string `json:"nickname"`
	Sex int `json:"sex"`
	Province string `json:"province"`
	City string `json:"city"`
	Country string `json:"country"`
	HeadImgUrl string `json:"headimgurl"`
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
}

// 网页授权获取用户信息
func GrantGetUserInfo(accessToken, openId string) (userInfo *WeChatUserInfo, err error) {
	url := "https://api.weixin.qq.com/sns/userinfo?access_token=" + accessToken + "&openid=" + openId + "&lang=zh_CN"
	client := http.Client{}
	request,err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, errors.New("网页授权获取用户信息:获取用户信息请求失败")
	}
	res := WeChatUserInfo{}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("网页授权获取用户信息:读取http body失败")
	}
	err =json.Unmarshal(body, &res)
	if err != nil {
		return nil, errors.New("网页授权获取用户信息:解析json失败")
	}

	if res.ErrCode != 0 {
		return nil, errors.New("get grant access token 失败: " + res.ErrMsg)
	}

	return &res, nil
}
