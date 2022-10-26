package main

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	ChatId int `json:"id"`
}

type RestResponse struct {
	Result []Update `json:"result"`
}

type BotMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}

type Configs struct {
	Auth auth
}

type auth struct {
	Token string "5178242009:AAHbLC1Gk4l42kTt4Tx_6qcTA_WyLW-aBxA"
	Api   string "https://api.telegram.org/bot"
}
