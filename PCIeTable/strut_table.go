package pcietable

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
