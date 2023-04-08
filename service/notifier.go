package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AccessToken struct {
	Access_token string `config.json:"access_token"`
}

func Notify(msg string) {
	config := ReadConfig()
	// using mobile app to notify you weather you has committed today
	wechat_url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
	full_url := fmt.Sprintf(wechat_url, config.Corpid, config.Corpsecret)
	client := http.Client{}
	req, err := http.NewRequest("GET", full_url, nil)
	HandleError(err)

	resp, err := client.Do(req)
	HandleError(err)
	defer resp.Body.Close()
	_bytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err)
	var token = &AccessToken{}
	json.Unmarshal(_bytes, token)
	access_token := token.Access_token
	json_dict := make(map[string]interface{})
	json_dict["touser"] = config.Touser
	json_dict["msgtype"] = "text"
	json_dict["agentid"] = config.Agentid
	text := make(map[string]string)
	text["content"] = msg
	json_dict["text"] = text
	json_dict["safe"] = 0
	json_dict["enable_id_trans"] = 0
	json_dict["duplicate_check_interval"] = 1800
	post_data_str, err := json.Marshal(json_dict)
	HandleError(err)
	send_msg_url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", access_token)
	req, err = http.NewRequest("POST", send_msg_url, bytes.NewBuffer([]byte(post_data_str)))
	HandleError(err)
	resp, err = client.Do(req)
	HandleError(err)

}
