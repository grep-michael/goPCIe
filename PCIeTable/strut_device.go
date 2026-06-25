package pcietable

import "fmt"

type Device struct {
	ID         string
	Name       string
	Vendor     *Vendor `json:"-"`
	VendorID   string
	SubDevices []*Device
}

func (dev *Device) addSubDevice(subDev *Device) {
	dev.SubDevices = append(dev.SubDevices, subDev)
}

func (dev *Device) AddChild(obj any) error {
	subDev, ok := obj.(*Device)
	if !ok {
		return fmt.Errorf("Device  Addchild: %T isnt a *Device type", obj)
	}
	dev.SubDevices = append(dev.SubDevices, subDev)
	return nil
}
