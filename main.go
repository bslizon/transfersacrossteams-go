package main

import (
	"log"

	"github.com/bslizon/transfersacrossteams-go/converter"
)

const (
	OLD_TEAM_ID = `Txxxxxxxx3`
	NEW_TEAM_ID = `3xxxxxxxxV`

	OLD_KEY_ID = `Jxxxxxxxx8`
	NEW_KEY_ID = `6xxxxxxxxA`

	OLD_CLIENT_ID = `com.xxxxxxxxxx.appstore`
	NEW_CLIENT_ID = `com.xxxxxxxxxx.appstore`

	OLD_KEY_PEM = `-----BEGIN PRIVATE KEY-----
MIGxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxU3
-----END PRIVATE KEY-----`

	NEW_KEY_PEM = `-----BEGIN PRIVATE KEY-----
MIGxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxx/V
-----END PRIVATE KEY-----`
)

func main() {
	testOldSub := `001xxx.xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx.xx44`
	testNewSub := `001xxx.xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx.xx49`

	newSub := convert(testOldSub)

	log.Println("old sub:", testOldSub, "----->", "new sub:", newSub)

	if newSub == testNewSub {
		log.Println("success")
	} else {
		log.Println("fail")
	}
}

func convert(oldSub string) string {
	oldClient := converter.New(OLD_TEAM_ID, OLD_KEY_ID, OLD_KEY_PEM, OLD_CLIENT_ID)
	if err := oldClient.Init(); err != nil {
		panic(err)
	}

	newClient := converter.New(NEW_TEAM_ID, NEW_KEY_ID, NEW_KEY_PEM, NEW_CLIENT_ID)
	if err := newClient.Init(); err != nil {
		panic(err)
	}

	transferSub, err := oldClient.FromSubToTransferSub(oldSub, newClient.TeamID)
	if err != nil {
		panic(err)
	}

	newSub, err := newClient.FromTransferSubToSub(transferSub)
	if err != nil {
		panic(err)
	}

	return newSub
}
