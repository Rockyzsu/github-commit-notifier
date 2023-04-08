package main

import (
	"fmt"
	"github-commit-notifier/service"
	"time"
)

func main() {
	today := time.Now().Format("2006-01-02")
	status := service.CommitStatus(today)
	if !status {
		msg := fmt.Sprintf("Today %s you did not commit any code in github!", today)
		fmt.Println(msg)
		service.Notify(msg)
	}
}
