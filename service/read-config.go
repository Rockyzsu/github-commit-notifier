package service

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type WechatArgInfo struct {
	Agentid    string `json:"agentid" gorm:"column:agentid"`
	Touser     string `json:"touser" gorm:"column:touser"`
	Corpid     string `json:"corpid" gorm:"column:corpid"`
	Corpsecret string `json:"corpsecret" gorm:"column:corpsecret"`
}

func ReadConfig() *WechatArgInfo {
	file, err := os.Open("service/config.json")
	HandleError(err)
	defer file.Close()
	byteData, err := ioutil.ReadAll(file)
	config := &WechatArgInfo{}
	json.Unmarshal(byteData, config)
	return config
}
