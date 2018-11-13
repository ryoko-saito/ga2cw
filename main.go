package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	analytics "google.golang.org/api/analytics/v3"
)

type Config struct {
	ProfileId string `json:"profile_id"`
	RoomId    string `json:"room_id"`
	ApiKey    string `json:"api_key"`
}

func main() {
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	var config Config
	err = json.Unmarshal(b, &config)
	if err != nil {
		log.Fatal(err)
	}
	client, err := google.DefaultClient(
		oauth2.NoContext,
		"https://www.googleapis.com/auth/analytics.readonly")
	if err != nil {
		log.Fatalf("Unable to read client : %v", err)
	}

	service, err := analytics.New(client)
	if err != nil {
		log.Fatalf("Unable to Access Google Analytics: %v", err)
	}

	result, err := service.Data.Ga.Get("ga:"+config.ProfileId, "yesterday", "yesterday", "ga:pageviews").Dimensions("ga:date").Do()

	if err != nil {
		log.Fatalf("Unable to get data: %v", err)
	}

	var date string //日付
	var pv string   //PV
	for _, row := range result.Rows {
		date = row[0]
		pv = row[1]
	}

	url := "https://api.chatwork.com/v2/rooms/" + config.RoomId + "/messages"

	//bodyの後には文字列がくる
	param := "body=" + date + " " + "GA PV:" + pv
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(param))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("X-ChatWorkToken", config.ApiKey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
}
