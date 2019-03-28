package nbiot

import (
	"io"
	"log"

	"github.com/telenordigital/nbiot-go"
)

func test() {
	client, err := nbiot.New()
	if err != nil {
		log.Fatal("Error creating client:", err)
	}

	collection, err := client.CreateCollection(nbiot.Collection{
		Tags: map[string]string{
			"name": "example collection",
		},
	})
	if err != nil {
		log.Fatal("Error creating collection: ", err)
	}

	imsi := "0345892703458"
	imei := "1487252347803"
	device, err := client.CreateDevice(collection.CollectionID, nbiot.Device{
		IMSI: &imei,
		IMEI: &imsi,
		Tags: map[string]string{
			"name": "example device",
		},
	})
	if err != nil {
		log.Fatal("Error creating device: ", err)
	}

	stream, err := client.DeviceOutputStream(collection.CollectionID, *device.DeviceID)
	if err != nil {
		log.Fatal("Error creating stream: ", err)
	}
	defer stream.Close()

	for {
		data, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error receiving data: ", err)
		}

		log.Print("received payload: ", data.Payload)
	}
}
