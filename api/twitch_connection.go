package api

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/addaboosh/winston-chat/config"
	"github.com/gorilla/websocket"
)

type Connection interface {
	Read()

	Connect()
	Authenticate()

	AddChat()    //handle n chats
	RemoveChat() //handle n chats
	PauseChat()  //handle n chats

	Disconnect()
	Shutdown()
}

const (
	DISCONNECTED = iota
	CONNECTED    = iota
	PAUSED       = iota
)

type TwitchConnection struct {
	cfg        config.TwitchConfiguration
	channels   []string
	interrupt  (chan os.Signal)
	done       (chan struct{})
	connection *websocket.Conn
	status     int
}

func (s *Server) NewTwitchConnection() *TwitchConnection {

	conn := &TwitchConnection{
		cfg:       s.cfg.TwitchConfiguration,
		channels:  make([]string, 0),
		interrupt: make(chan os.Signal, 1),
		done:      make(chan struct{}),
		status:    DISCONNECTED,
	}
	signal.Notify(conn.interrupt, os.Interrupt)
	return conn
}

func (t *TwitchConnection) Read() {
	go func() {
		defer close(t.done)
		for {
			_, message, err := t.connection.ReadMessage()
			if err != nil {
				log.Printf("Error Reading msg: %v", err)
				return
			}
			log.Printf("%v", string(message[:]))
		}
	}()
	for {
		select {
		case <-t.done:
			log.Println("we done")
			return
		case <-t.interrupt:
			log.Println("interrupted you bastard")

			// Attempt to cleanly close conn by sending a WS close message
			// and awaiting (w/timeout) for the conn to close.
			err := t.connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))

			if err != nil {
				log.Printf("Error closing websocket: %v", err)
				return
			}
			return
		}
	}
}

func (t *TwitchConnection) Connect() {
	log.Println("Building Connection Details....")
	addr := t.cfg.Url
	u := url.URL{Scheme: "wss", Host: addr}

	log.Printf("Connecting to %v", u)
	c, r, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		log.Printf("Failed to connect to %s", u.String())
	}
	if r.StatusCode != 200 {
		log.Printf("HTTP non-200 %d", r.StatusCode)
	} else {
		t.connection = c
		t.status = CONNECTED
	}
}

func (t *TwitchConnection) Authenticate() {
	t.connection.WriteMessage(1, []byte("CAP REQ :twitch.tv/membership twitch.tv/tags"))
	t.connection.WriteMessage(1, []byte("NICK justinfan821"))
	t.connection.WriteMessage(1, []byte("PASS ANONYAUTH"))
}

func (t *TwitchConnection) AddChannel(channels []string) {
	for _, channel := range channels {
		// Join channel
		joinmsg := fmt.Sprintf("JOIN #%s", channel)
		t.connection.WriteMessage(1, []byte(joinmsg))

		// Add to channel list
		t.channels = append(t.channels, channel)
	}

}
