package pcietable

import "strings"

type Vendor struct {
	VendorID string
	Name     string
	Devices  map[string]*Device //devices within an vendor all have different ids, so i thought
}

func (vendor *Vendor) FindDevice(id string) (*Device, bool) {
	id = strings.TrimPrefix(id, "0x")
	dev, ok := vendor.Devices[id]
	return dev, ok
}

func (vendor *Vendor) AddChild(line string) Child {
	dev := lineToDevice(line)
	if _, ok := vendor.Devices[dev.DeviceID]; !ok {
		vendor.Devices[dev.DeviceID] = dev
		dev.Vendor = vendor
		dev.VendorID = vendor.VendorID
	}
	return dev
}
