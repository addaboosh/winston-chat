package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/url"
	"os"
	"os/signal"
)

func CreateTwitchConnection() {
	// Handle Args [program, channel name]
	channel := os.Args[1]

	//Log setup
	initLog()

	//handle OS interrupt for graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	//chan for happy shutdown
	done := make(chan struct{})

	log.Println("Hello, addboosh!")
	log.Println("Building connection details....")
 
	addr := "irc-ws.chat.twitch.tv:443"
	u := url.URL{Scheme: "wss", Host: addr}

	log.Printf("Connecting to: %v", u.String())

	c, r, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		log.Println("Could not connect to URL", u.String())
	}
	log.Printf("HTTP resp: %v", r.Status)

	defer c.Close()
	c.WriteMessage(1, []byte("CAP REQ :twitch.tv/membership twitch.tv/tags"))
	c.WriteMessage(1, []byte("NICK justinfan821"))
	c.WriteMessage(1, []byte("PASS ANONYAUTH"))

	// Join channel
	joinmsg := fmt.Sprintf("JOIN #%s", channel)
	c.WriteMessage(1, []byte(joinmsg))

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Printf("Error Reading msg: %v", err)
				return
			}
			log.Printf("%v", string(message[:]))
		}
	}()

	for {
		select {
		case <-done:
			log.Println("we done")
			return
		case <-interrupt:
			log.Println("interrupted you bastard")

			// Attempt to cleanly close conn by sending a WS close message
			// and awaiting (w/timeout) for the conn to close.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))

			if err != nil {
				log.Printf("Error closing websocket: %v", err)
				return
			}
			return
		}
	}
}

func initLog() {
	logfn := fmt.Sprintf("./data/chatlog-%v.txt", log.Lshortfile)
	w, err := os.OpenFile(logfn, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic("Unable to open log file")
	}
	mw := io.MultiWriter(os.Stdout, w)
	log.SetOutput(mw)

}
