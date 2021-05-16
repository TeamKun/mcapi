package minecraft

import (
	"fmt"
	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/yggdrasil"
	"github.com/google/uuid"
	"log"
	"os"
)

func Login(address string, user string, password string) {
	resp, err := yggdrasil.Authenticate(user, password)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	id, name := resp.SelectedProfile()

	client := bot.NewClient()
	client.Auth.Name = name
	client.Auth.UUID = id
	client.Auth.AsTk = resp.AccessToken()
	//player := basic.NewPlayer(client, basic.DefaultSettings)

	basic.EventsListener{
		ChatMsg:    onChatMsg,
		Disconnect: onDisconnect,
	}.Attach(client)

	//Login
	err = client.JoinServer(address)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	//JoinGame
	err = client.HandleGame()
	if err != nil {
		log.Fatal(err)
	}
}

func onChatMsg(c chat.Message, pos byte, uuid uuid.UUID) error {
	log.Println("Chat:", c)
	return nil
}

func onDisconnect(c chat.Message) error {
	log.Println("Disconnect:", c)
	return nil
}
