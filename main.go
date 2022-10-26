package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	var conf Configs

	conf.Auth.Token = "5178242009:AAHbLC1Gk4l42kTt4Tx_6qcTA_WyLW-aBxA"
	conf.Auth.Api = "https://api.telegram.org/bot"

	botUrl := conf.Auth.Api + conf.Auth.Token
	offset := 0

	for {
		updates, err := getUpdates(botUrl, offset)

		if err != nil {
			log.Println("Error while updating", err.Error())
		}

		for _, update := range updates {
			err = respond(botUrl, update)
			offset = update.UpdateId + 1
		}

		fmt.Println(updates)
	}

}

func getUpdates(botUrl string, offset int) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))

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

func respond(botUrl string, update Update) error {
	var BotMes BotMessage
	BotMes.ChatId = update.Message.Chat.ChatId
	BotMes.Text = update.Message.Text
	buf, err := json.Marshal(BotMes)

	if err != nil {
		return err
	}

	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))

	if err != nil {
		return err
	}

	return nil
}
