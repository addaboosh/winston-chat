package api

import (
	"log"
	"os"
	"os/signal"

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

func (s *Server) NewTwitchConnection () *TwitchConnection{

	conn := &TwitchConnection{
		cfg: s.cfg.TwitchConfiguration,
		channels: make([]string,0),
		interrupt: make(chan os.Signal, 1),
		done: make(chan struct{}),
	}
	signal.Notify(conn.interrupt, os.Interrupt)
	return conn
}

func Connect(t *TwitchConnection) {
	log.Println("Starting Connection....")
}	


	
