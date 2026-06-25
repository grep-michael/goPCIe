package pcietable

type Device struct {
	DeviceID   string
	Name       string
	Vendor     *Vendor `json:"-"`
	VendorID   string
	SubDevices []*Device
}

func (dev *Device) AddChild(line string) Child {
	subDev := lineToSubDevice(line)
	dev.SubDevices = append(dev.SubDevices, subDev)
	return subDev
}
