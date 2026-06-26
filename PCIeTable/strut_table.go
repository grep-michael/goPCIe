package pcietable

import "strings"

type PCITable struct {
	Sources []string
	Vendors map[string]*Vendor //all vendors have different ids
	Classes map[string]*Class
}

func (table *PCITable) FindVendor(id string) (*Vendor, bool) {
	id = strings.TrimPrefix(id, "0x")
	ven, ok := table.Vendors[id]
	return ven, ok
}
func (table *PCITable) FindClass(id string) (*Class, bool) {
	id = strings.TrimPrefix(id, "0x")
	foundClass, ok := table.Classes[id[0:2]]
	if !ok {
		return nil, false
	}
	for i := 2; ; i = i + 2 {
		if i+2 > len(id) {
			break
		}
		code := id[i : i+2]
		tmpClass, ok2 := foundClass.SubClasses[code]
		if !ok2 {
			break
		}
		foundClass = tmpClass
	}

	return foundClass, true
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
