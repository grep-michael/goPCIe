package pcietable

import (
	"fmt"
	"log"
)

type Vendor struct {
	ID      string
	Name    string
	Devices map[string]*Device //devices within an vendor all have different ids, so i thought
}

func (vendor *Vendor) FindDevice(id string) (*Device, bool) {
	dev, ok := vendor.Devices[id]
	return dev, ok
}

func (vendor *Vendor) addDevice(dev *Device, table *PCITable) {
	if _, ok := vendor.Devices[dev.ID]; !ok {
		vendor.Devices[dev.ID] = dev
		dev.Vendor = vendor
		dev.VendorID = vendor.ID
	} else {
		log.Printf("Attempted to add duplicate device(%s) to vendor(%s)\n", dev.Name, vendor.Name)
	}
	table.RegisterDevice(dev)
}
func (vendor *Vendor) AddChild(obj any) error {
	dev, ok := obj.(*Device)
	if !ok {
		return fmt.Errorf("Vendor Addchild: %T is not *Device type", obj)
	}
	if _, ok := vendor.Devices[dev.ID]; !ok {
		vendor.Devices[dev.ID] = dev
		dev.Vendor = vendor
		dev.VendorID = vendor.ID
	}
	return nil
}
