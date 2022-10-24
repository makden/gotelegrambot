package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	botToken := "5178242009:AAHbLC1Gk4l42kTt4Tx_6qcTA_WyLW-aBxA"
	botApi := "https://api.telegram.org/bot"
	botUrl := botApi + botToken

	for {
		updates, err := getUpdates(botUrl)

		if err != nil {
			log.Println("Error while updating", err.Error())
		}

		fmt.Println(updates)
	}

}

func getUpdates(botUrl string) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates")

	fmt.Println(botUrl + "/getUpdates")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var res RestResponse
	err = json.Unmarshal(body, &res) // Почему просто = а не :=

	if err != nil {
		return nil, err
	}

	return res.Result, nil
}

func respond() {

}
