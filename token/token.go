// Package token 用于调用 openapi，websocket 的 token 对象。
package token

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/tencent-connect/botgo/log"
	"gopkg.in/yaml.v3"
)

// Type token 类型
type Type string

// TokenType
const (
	TypeBot    Type = "Bot"
	TypeNormal Type = "Bearer"
	TypeQQBot  Type = "QQBot"
)

// Token 用于调用接口的 token 结构
type Token struct {
	AppID       uint64
	AccessToken string
	secret      string
	Type        Type

	expirationTime time.Time
	authToken      string
	lock           sync.RWMutex
}

//// New 创建一个新的 Token
//func New(tokenType Type) *Token {
//	return &Token{
//		Type: tokenType,
//	}
//}

func QQBotToken(appID uint64, accessToken, secret string) *Token {
	return &Token{
		AppID:       appID,
		AccessToken: accessToken,
		secret:      secret,
		Type:        TypeQQBot,
	}
}

// BotToken 机器人身份的 token
func BotToken(appID uint64, accessToken string) *Token {
	return &Token{
		AppID:       appID,
		AccessToken: accessToken,
		Type:        TypeBot,
	}
}

// UserToken 用户身份的token
func UserToken(appID uint64, accessToken string) *Token {
	return &Token{
		AppID:       appID,
		AccessToken: accessToken,
		Type:        TypeNormal,
	}
}

// GetString 获取授权头字符串
func (t *Token) GetString() string {
	if t.Type == TypeNormal {
		return t.AccessToken
	}
	// 这边要获取auth token
	if t.Type == TypeQQBot {
		return t.authAccessToken(t.AppID, t.secret)
	}
	return fmt.Sprintf("%v.%s", t.AppID, t.AccessToken)
}

// LoadFromConfig 从配置中读取 appid 和 token
func (t *Token) LoadFromConfig(file string) error {
	var conf struct {
		AppID uint64 `yaml:"appid"`
		Token string `yaml:"token"`
	}
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Errorf("read token from file failed, err: %v", err)
		return err
	}
	if err = yaml.Unmarshal(content, &conf); err != nil {
		log.Errorf("parse config failed, err: %v", err)
		return err
	}
	t.AppID = conf.AppID
	t.AccessToken = conf.Token
	return nil
}

type authAccessTokenResult struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

func (t *Token) authAccessToken(appid uint64, secret string) string {
	t.lock.RLock()
	if time.Now().Sub(t.expirationTime) > 0 {
		t.lock.RUnlock()
	} else {
		t.lock.RUnlock()
		return t.authToken
	}
	n := 0
LOOP:
	if n != 0 {
		time.Sleep(time.Millisecond * 100)
	}
	if n > 10 {
		panic("get auth token failed")
	}
	resp, err := http.Post("https://bots.qq.com/app/getAppAccessToken", "application/json", strings.NewReader(fmt.Sprintf(`{"appId":"%d","clientSecret":"%s"}`, appid, secret)))
	if err != nil {
		log.Errorf("get auth token failed, err: %v", err)
		n++
		goto LOOP
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("get auth token failed, err: %v", err)
		n++
		goto LOOP
	}
	authToken := &authAccessTokenResult{}
	if err = json.Unmarshal(body, authToken); err != nil {
		log.Errorf("get auth token failed, err: %v", err)
		n++
		goto LOOP
	}
	if len(authToken.AccessToken) == 0 {
		log.Errorf("get auth token failed, err: %v", string(body))
		n++
		goto LOOP
	}
	t.authToken = authToken.AccessToken
	ei, _ := strconv.Atoi(authToken.ExpiresIn)
	t.expirationTime = time.Now().Add(time.Duration(ei-50) * time.Second)
	return authToken.AccessToken
}
