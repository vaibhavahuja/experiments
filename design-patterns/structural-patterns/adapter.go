package main

import "fmt"

//for example you are working with one XML
// and library supports only JSON
// how do you convert your XML to JSON?
// You can either convert that library to accept only XML.
// but do you really have access to that source code?
// Use adapter pattern, as the name suggests

type Client struct {
}
type Computer interface {
	InsertIntoLightningPort()
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}

type Mac struct {
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts lighting signal to USB")
	w.windowMachine.insertIntoUSBPort()
}

func main() {
	client := &Client{}
	mac := &Mac{}
	client.InsertLightningConnectorIntoComputer(mac)
	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{windowsMachine}
	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
