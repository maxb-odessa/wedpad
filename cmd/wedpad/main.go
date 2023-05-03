// websockets.go
package main

import (
	"fmt"
	"net/http"
	"time"

	ws "github.com/gorilla/websocket"
)

var upgrader = ws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("conn!")
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("err conn")
			return
		}
		i := 0
		for {
			//_, _, err = conn.ReadMessage()
			//if err != nil {
			//	return
			//}
			//fmt.Println("got msg")

			// Write message back to browser
			time.Sleep(1 * time.Second)
			if err = conn.WriteMessage(ws.TextMessage, []byte("message")); err != nil {
				fmt.Println("send failed")
				return
			}
			fmt.Println("send msg")
			if i > 5 {
				conn.Close()
				return
			}
			i++
		}
	})

	http.Handle("/", http.FileServer(http.Dir("page")))

	http.ListenAndServe(":8080", nil)
}
