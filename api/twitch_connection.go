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
	l          *log.Logger
	cfg        config.TwitchConfiguration
	channels   []string
	interrupt  (chan os.Signal)
	done       (chan struct{})
	connection *websocket.Conn
	status     int
}

func (s *Server) NewTwitchConnection() *TwitchConnection {

	s.l.Println("Hello from twichconneciton")
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
				t.l.Printf("Error Reading msg: %v", err)
				return
			}
			t.l.Printf("%v", string(message[:]))
		}
	}()
	for {
		select {
		case <-t.done:
			t.l.Println("we done")
			return
		case <-t.interrupt:
			t.l.Println("interrupted you bastard")

			// Attempt to cleanly close conn by sending a WS close message
			// and awaiting (w/timeout) for the conn to close.
			err := t.connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))

			if err != nil {
				t.l.Printf("Error closing websocket: %v", err)
				return
			}
			return
		}
	}
}

func (t *TwitchConnection) Connect() {
	t.l.Println("Building Connection Details....")
	addr := t.cfg.Url
	u := url.URL{Scheme: "wss", Host: addr}

	t.l.Printf("Connecting to %v", u)
	c, r, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		t.l.Printf("Failed to connect to %s", u.String())
	}
	if r.StatusCode != 200 {
		t.l.Printf("HTTP non-200 %d", r.StatusCode)
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

func (t *TwitchConnection) JoinChannel(channels []string) {
	for _, channel := range channels {
		// Join channel
		joinmsg := fmt.Sprintf("JOIN #%s", channel)
		t.connection.WriteMessage(1, []byte(joinmsg))

		// Add to channel list
		t.channels = append(t.channels, channel)
	}

}
