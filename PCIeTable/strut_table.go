package pcietable

type PCITable struct {
	Sources []string
	Vendors map[string]*Vendor //all vendors have different ids
	Classes map[string]*Class
}

func (table *PCITable) FindVendor(id string) (*Vendor, bool) {
	ven, ok := table.Vendors[id]
	return ven, ok
}
func (table *PCITable) RegisterClass(class *Class) {
	if _, ok := table.Classes[class.ClassID]; ok {
		return
	}
	table.Classes[class.ClassID] = class
}
func (table *PCITable) RegisterVendor(ven *Vendor) {
	if _, ok := table.Vendors[ven.VendorID]; ok {
		return
	}
	table.Vendors[ven.VendorID] = ven
}
