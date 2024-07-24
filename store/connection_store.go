package store

import (
	"os"

	"github.com/addaboosh/winston-chat/config"
)

type Connection interface {
	Connect()
	Authenticate()

	AddChat() //handle n chats
	RemoveChat() //handle n chats
	PauseChat() //handle n chats

	Disconnect()
	Shutdown()

}


type TwitchConnection struct {
	cfg config.TwitchConfiguration
	channels []string
	interrupt (chan os.Signal)
	done (chan struct{})

}
