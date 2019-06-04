package main

import (
	"fmt"
	"log"
	"strings"
	"time"
	// "goal"

	"github.com/paypal/gatt"
	"github.com/paypal/gatt/examples/option"
)

var (
	espRed                  = "RED_GOAL"
	espBlue                 = "BLUE_GOAL"
	currentDevice           = ""
	goalServiceUUID         = "298ed54e6b1911e9a9231681be663d3e"
	goalCharacteristicsUUID = "298ed8d2-6b19-11e9-a923-1681be663d3e"
)

func main() {
	log.SetFlags(log.Llongfile)
	addGoalRed()
	d, err := gatt.NewDevice(option.DefaultClientOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s\n", err)
		return
	}

	// Register handlers.
	d.Handle(
		gatt.PeripheralDiscovered(onPeriphDiscovered),
		gatt.PeripheralConnected(onPeriphConnected),
		gatt.PeripheralDisconnected(onPeriphDisconnected),
	)

	d.Init(onStateChanged)
	<-done
	fmt.Println("Done")
}

var done = make(chan struct{})

func onStateChanged(d gatt.Device, s gatt.State) {
	fmt.Println("State:", s)
	switch s {
	case gatt.StatePoweredOn:
		fmt.Println("Scanning...")
		d.Scan([]gatt.UUID{}, false)
		return
	default:
		d.StopScanning()
	}
}

// OnPeriphDiscovered checks fo peripherals.
// It returns if a name does not match a predefined name
func onPeriphDiscovered(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {
	log.Println("Discovered device with ID: " + p.ID() + " and name: " + p.Name())
	if strings.ToUpper(p.Name()) == strings.ToUpper(espRed) || strings.ToUpper(p.Name()) == strings.ToUpper(espBlue) {
		currentDevice = p.Name()
		p.Device().StopScanning()
		p.Device().Connect(p)
	} else if strings.ToUpper(p.Name()) == espBlue {
		p.Device().StopScanning()
		p.Device().Connect(p)
	} else {
		return
	}
}

func onPeriphConnected(p gatt.Peripheral, err error) {
	log.Println("Connected to: " + p.Name())
	defer p.Device().CancelConnection(p)

	// MTU (Maximum Transmission Unit) is the maximum length of an ATT packet.
	if err := p.SetMTU(500); err != nil {
		log.Printf("Failed to set MTU, err: %s\n", err)
	}

	// Discovery services
	ss, err := p.DiscoverServices(nil)
	if err != nil {
		log.Printf("Failed to discover services, err: %s\n", err)
		return
	}

	// Range over discovered services and check for desired services.
	for _, s := range ss {
		msg := "Service: " + s.UUID().String()
		if len(s.Name()) > 0 {
			msg += " (" + s.Name() + ")"
		}
		log.Println(s.Name())
		log.Println(s.UUID())

		cs, err := p.DiscoverCharacteristics(nil, s)
		if err != nil {
			log.Printf("Failed to discover characteristics, err: %s\n", err)
			continue
		}

		// Range over discovered characteristics and check if desired characteristic is there (currently no check).
		// Currently we check for
		for _, c := range cs {
			if c.Name() != "" {
				log.Println(c.Name())
			}
			msg += "\nCharacteristic:  " + c.UUID().String()
			if len(c.Name()) > 0 {
				msg += " (" + c.Name() + ")"
			}
			msg += "\nProperties:    " + c.Properties().String()
			log.Println(msg)

			// Read the characteristic, if possible.
			if (c.Properties() & gatt.CharRead) != 0 {
				b, err := p.ReadCharacteristic(c)
				if err != nil {
					log.Printf("Failed to read characteristic, err: %s\n", err)
					continue
				}
				log.Printf("    value         %x | %q\n", b, b)
			}

			// Discovery descriptors
			ds, err := p.DiscoverDescriptors(nil, c)
			if err != nil {
				log.Printf("Failed to discover descriptors, err: %s\n", err)
				continue
			}

			for _, d := range ds {
				msg := "  Descriptor      " + d.UUID().String()
				if len(d.Name()) > 0 {
					msg += " (" + d.Name() + ")"
				}
				log.Println(msg)

				// Read descriptor (could fail, if it's not readable)
				b, err := p.ReadDescriptor(d)
				if err != nil {
					log.Printf("Failed to read descriptor, err: %s\n", err)
					continue
				}
				log.Printf("    value         %x | %q\n", b, b)
			}

			// Subscribe the characteristic, if possible.
			if (c.Properties() & (gatt.CharIndicate)) != 0 {
				f := func(c *gatt.Characteristic, b []byte, err error) {
					log.Printf("indicated: % X | %q\n", b, b)
					log.Println(string(b))
					if true {
						log.Println("writing stuff")
						p.WriteCharacteristic(c, b, true)
					}
				}
				if err := p.SetIndicateValue(c, f); err != nil {
					log.Printf("Failed to subscribe characteristic, err: %s\n", err)
					continue
				}
			}

		}
		log.Println()
	}

	log.Printf("Waiting for 5 seconds to get some notifiations, if any.\n")
	time.Sleep(5 * time.Second)
}

func onPeriphDisconnected(p gatt.Peripheral, err error) {
	log.Println("Disconnected")
	close(done)
}
