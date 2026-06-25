package pcietable

type PCITable struct {
	Sources []string
	Vendors map[string]*Vendor //all vendors have different ids
}

func (table *PCITable) FindVendor(id string) (*Vendor, bool) {
	ven, ok := table.Vendors[id]
	return ven, ok
}

func (table *PCITable) RegisterVendor(ven *Vendor) {
	if _, ok := table.Vendors[ven.ID]; ok {
		//log.Printf("Attempted to add vender already in table: %s\n", ven.Name)
		return
	}
	table.Vendors[ven.ID] = ven
}
