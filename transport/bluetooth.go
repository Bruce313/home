package transport

import (
	"github.com/muka/go-bluetooth/api"
	"log"
)

func GetDevices()  {
	ds, err := api.GetDevices()
	if err != nil {
		panic(err)
	}
	log.Printf("devices: %s", ds)
}
