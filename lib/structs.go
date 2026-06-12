package lib

import "log"

type PCITable struct {
	Sources []string
	Vendors map[string]*Vendor   //all vendors have different ids
	Devices map[string][]*Device //devices can have the same id
}

func (table *PCITable) FindVendor(id string) (*Vendor, bool) {
	ven, ok := table.Vendors[id]
	return ven, ok
}

func (table *PCITable) FindDevice(id string) ([]*Device, bool) {
	devices, ok := table.Devices[id]
	return devices, ok
}
func (table *PCITable) RegisterDevice(dev *Device) {
	if table.Devices[dev.ID] == nil {
		table.Devices[dev.ID] = make([]*Device, 0)
	}
	table.Devices[dev.ID] = append(table.Devices[dev.ID], dev)
}
func (table *PCITable) RegisterVendor(ven *Vendor) {
	if _, ok := table.Vendors[ven.ID]; ok {
		//log.Printf("Attempted to add vender already in table: %s\n", ven.Name)
		return
	}
	table.Vendors[ven.ID] = ven
}

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
