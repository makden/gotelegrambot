package main

type Configs struct {
	Auth auth
}

type auth struct {
	Token string //`default: "5178242009:AAHbLC1Gk4l42kTt4Tx_6qcTA_WyLW-aBxA"`
	Api   string //`default: "https://api.telegram.org/bot"`
}
