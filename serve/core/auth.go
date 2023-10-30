package core

import (
	"encoding/json"
	"time"
)

var AppKey string

type Auth struct {
	Uid  int64
	Time int64
}

func Entoken(uid int64) (string, error) {
	bs, err := json.Marshal(Auth{
		Uid:  uid,
		Time: time.Now().Unix(),
	})
	return AesEncrypt(string(bs), AppKey), err
}

func Detoken(token string) (*Auth, error) {
	auth := &Auth{}
	err := json.Unmarshal([]byte(AesDecrypt(token, AppKey)), auth)
	return auth, err
}
