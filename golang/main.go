package main

import (
	"encoding/json"
	"fmt"
	aisstream "github.com/aisstream/ais-message-models/golang/aisStream"
	"github.com/gorilla/websocket"
	"log"
)

func main() {

	url := "wss://stream.aisstream.io/v0/stream"

	// Enables compression, which aisstream.io requires to serve full message bandwidth.
	dialer := *websocket.DefaultDialer
	dialer.EnableCompression = true

	ws, _, err := dialer.Dial(url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer ws.Close()

	subMsg := aisstream.SubscriptionMessage{
		APIKey:        "<YOUR API KEY>",
		BoundingBoxes: [][][]float64{{{-90.0, -180.0}, {90.0, 180.0}}}, // bounding box for the entire world
	}

	subMsgBytes, _ := json.Marshal(subMsg)
	if err := ws.WriteMessage(websocket.TextMessage, subMsgBytes); err != nil {
		log.Fatalln(err)
	}

	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			log.Fatalln(err)
		}
		var packet aisstream.AisStreamMessage

		if err := json.Unmarshal(p, &packet); err != nil {
			// The server's first message is a SubscriptionConfirmation, which the
			// AIS message models do not describe. Skip anything that is not an
			// AIS message rather than treating it as fatal.
			continue
		}

		var shipName string
		// field may or may not be populated
		if packetShipName, ok := packet.MetaData["ShipName"]; ok {
			shipName = packetShipName.(string)
		}

		switch packet.MessageType {
		case aisstream.POSITION_REPORT:
			var positionReport aisstream.PositionReport
			positionReport = *packet.Message.PositionReport
			fmt.Printf("MMSI: %d Ship Name: %s Latitude: %f Longitude: %f\n",
				positionReport.UserID, shipName, positionReport.Latitude, positionReport.Longitude)
		}

	}

}
