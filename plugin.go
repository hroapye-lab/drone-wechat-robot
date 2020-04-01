package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	Repo struct {
		Owner string `json:"owner"`
		Name  string `json:"name"`
	}

	Build struct {
		Tag     string `json:"tag"`
		Event   string `json:"event"`
		Number  int    `json:"number"`
		Commit  string `json:"commit"`
		Ref     string `json:"ref"`
		Branch  string `json:"branch"`
		Author  string `json:"author"`
		Message string `json:"message"`
		Status  string `json:"status"`
		Link    string `json:"link"`
		Started int64  `json:"started"`
		Created int64  `json:"created"`
	}

	Config struct {
		Key     string `json:"key"`
		MsgType string `json:"msgtype"`
		Content string `json:"content"`
	}

	Job struct {
		Started int64 `json:"started"`
	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Config Config
		Job    Job
	}
)

func (p Plugin) Exec() error {
	var b []byte

	url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + p.Config.Key

	if p.Config.MsgType == "text" {
		text := struct {
			Content string `json:"content"`
		}{p.Config.Content}
		data := struct {
			MsgType string `json:"msgtype"`
			Text    struct {
				Content string `json:"content"`
			} `json:"text"`
		}{p.Config.MsgType, text}

		b, _ = json.Marshal(data)
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		fmt.Printf("Error: failed to create the request. %s\n", err)
	}
	request.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("Error: failed to excute the HTTP request. %s\n", err)
		return err
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	responseBody, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(responseBody))

	return nil
}
