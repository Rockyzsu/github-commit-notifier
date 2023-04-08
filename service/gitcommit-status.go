package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const GIT_URL = "%s/search/commits?q=author:%s+committer-date:%s"

func GetContentWithHeader(url string) (string, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	HandleError(err)
	req.Header.Add("Authorization", "token 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")
	resp, err := client.Do(req)
	HandleError(err)
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err)

	return string(bytes), nil
}

type CommitInfo struct {
	Total_count int `config.json:"total_count"`
}

func ParseResult(text string) bool {
	var resp = &CommitInfo{}
	json.Unmarshal([]byte(text), resp)
	count := resp.Total_count
	if count > 0 {
		return true
	} else {
		return false
	}

}

func CommitStatus(date string) bool {
	req_url := fmt.Sprintf(GIT_URL, "https://api.github.com", "Rockyzsu", date)
	content, _ := GetContentWithHeader(req_url)
	has_commit := ParseResult(content)
	if !has_commit {
		return false
	} else {
		return true
	}
}
