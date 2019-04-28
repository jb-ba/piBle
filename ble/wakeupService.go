package ble

import (
	"fmt"
	"github.com/paypal/gatt"
)

func NewWakeupService() *gatt.Service {
	n:=0
	s := gatt.NewService(gatt.MustParseUUID("6b759a6a-b416-4c89-9433-dcb9cc8c3e57"))
	s.AddCharacteristic(gatt.MustParseUUID("7b534b61-fb20-4823-8171-ed791ee7694f")).HandleReadFunc(
		func(rsp gatt.ResponseWriter, req *gatt.ReadRequest) {
			fmt.Fprintf(rsp, "count: %d", n)
			n++
		})

	//s.AddCharacteristic(gatt.MustParseUUID("7b534b61-fb20-4823-8171-ed791ee7694f")).HandleWriteFunc(
	//	func(r gatt.Request, data []byte) (status byte) {
	//		log.Println("Wrote:", string(data))
	//		return gatt.StatusSuccess
	//	})
	//
	//s.AddCharacteristic(gatt.MustParseUUID("1c927b50-c116-11e3-8a33-0800200c9a66")).HandleNotifyFunc(
	//	func(r gatt.Request, n gatt.Notifier) {
	//		cnt := 0
	//		for !n.Done() {
	//			fmt.Fprintf(n, "Count: %d", cnt)
	//			cnt++
	//			time.Sleep(time.Second)
	//		}
	//	})

	return s
}
