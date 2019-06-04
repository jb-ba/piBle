package test

import (
	"crypto/tls"
	"github.com/gorilla/websocket"
	"log"
	// "strconv"
	"testing"
	// "time"
)

func TestAaddGoalRed(t *testing.T) {
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("start ws client")
	addr := "iafoosball.me:9003"
	tableID := "table-1"
	userID := "tableUser"

	u := "wss://" + addr + "/users/?tableID=" + tableID + "&userID=" + userID
	log.Println(u)
	log.Printf("connecting to %s", u)
	d := websocket.Dialer{}
	d.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	conn, _, err := d.Dial(u, nil)

	message := "{ \"command\": \"addGoal\", \"values\": { \"speed\": 12, \"side\": \"red\", \"setposition\": \"attack\"  }}"
	log.Println(message)

	err = conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return
	}
	// w.Write([]byte(message))

	if err != nil {
		log.Println(err)
	}
	// handleErr(err, "making websocket connection")
	// }
	//defer c.Close()

	//_ := &client{
	//	send: make(chan []byte, 256),
	//}

}
