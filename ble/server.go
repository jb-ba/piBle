package ble

import (
	"fmt"
	"github.com/paypal/gatt"
	"github.com/paypal/gatt/examples/option"
	"github.com/paypal/gatt/examples/service"
	"log"
)

type server struct {
	Name string
	UUID string
}

func (s *server) DefaultServerSpec() {
	s.Name = "iaf-goal-server"
	s.UUID = "6b759a6a-b416-4c89-9433-dcb9cc8c3e57"
}

func New() {
	d, err := gatt.NewDevice(option.DefaultServerOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s", err)
	}

	// Register optional handlers.
	d.Handle(
		gatt.CentralConnected(func(c gatt.Central) { fmt.Println("Connect: ", c.ID()) }),
		gatt.CentralDisconnected(func(c gatt.Central) { fmt.Println("Disconnect: ", c.ID()) }),
	)

	// A mandatory handler for monitoring device state.
	onStateChanged := func(d gatt.Device, s gatt.State) {
		fmt.Printf("State: %s\n", s)
		switch s {
		case gatt.StatePoweredOn:
			log.Println("powered on")

			s := server{}
			s.DefaultServerSpec()

			// Setup GAP and GATT services for Linux implementation.
			d.AddService(service.NewGapService(s.Name)) // no effect on OS X
			d.AddService(service.NewGattService())        // no effect on OS X

			// A simple count service for demo.
			s1 := NewBatteryService()
			d.AddService(s1)

			// Advertise device name and service's UUIDs.
			d.AdvertiseNameAndServices(s.Name, []gatt.UUID{s1.UUID()})

			// Advertise as an OpenBeacon iBeacon
			d.AdvertiseIBeacon(gatt.MustParseUUID(s.UUID), 1, 2, -59)


		default:
		}
	}

	d.Init(onStateChanged)
	select {}
	log.Println("Bye")

}
