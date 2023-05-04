package server

import (
	"net/http"

	ws "github.com/gorilla/websocket"

	"github.com/maxb-odessa/sconf"
	"github.com/maxb-odessa/slog"
)

var toClient chan []byte

func init() {
	toClient = make(chan []byte, 8)
}

func Send(s []byte) {
	select {
	case toClient <- s:
	default:
		slog.Debug(5, "Client is not connected, discarding the message")
	}
}

var onConnectCb func()

func OnConnect(f func()) {
	onConnectCb = f
}

func Run() {

	var upgrader = ws.Upgrader{
		ReadBufferSize:  8192,
		WriteBufferSize: 8192,
	}

	wsHandler := func(w http.ResponseWriter, r *http.Request) {

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			slog.Err("Websocket upgrade failed: %s", err)
			return
		}

		slog.Info("Websocket connected: %s", conn.RemoteAddr())

		if onConnectCb != nil {
			onConnectCb()
		}

		defer func() {
			slog.Info("Websocket connection closed: %s", conn.RemoteAddr())
			conn.Close()
		}()

		reader := func() {
			for {
				// catch remote connection close
				_, _, err := conn.ReadMessage()
				if err != nil {
					return
				}
			}
		}

		go reader()

		for {

			select {
			case msg, ok := <-toClient:
				if !ok {
					return
				}
				if err = conn.WriteMessage(ws.TextMessage, msg); err != nil {
					slog.Err("Websocket send() failed: %s", err)
					return
				} else {
					slog.Debug(5, "Websocket sent: %q", string(msg))
				}
			}
		}
	}

	http.HandleFunc("/ws", wsHandler)

	pageDir := sconf.StrDef("paths", "page", "page")
	slog.Debug(9, "Serving HTTP dir: %s", pageDir)
	http.Handle("/", http.FileServer(http.Dir(pageDir)))

	http.ListenAndServe(sconf.StrDef("net", "listen", "localhost:8080"), nil)
}
