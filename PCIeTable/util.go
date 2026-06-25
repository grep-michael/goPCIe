package pcietable

import (
	"strings"
)

func setUpTable(table *PCITable) {
	if table.Vendors == nil {
		table.Vendors = make(map[string]*Vendor)
	}
	if table.Sources == nil {
		table.Sources = make([]string, 0)
	}
	if table.Classes == nil {
		table.Classes = make(map[string]*Class)
	}
}

func lineToDevice(line string) *Device {
	line = strings.TrimSpace(line)
	lines := strings.SplitN(line, "  ", 2)
	return &Device{
		DeviceID:   lines[0],
		Name:       lines[1],
		SubDevices: make([]*Device, 0),
	}
}
func lineToSubDevice(line string) *Device {
	line = strings.TrimSpace(line)
	parts := strings.SplitN(line, "  ", 2)
	ids := strings.Split(parts[0], " ")
	return &Device{
		DeviceID:   ids[1],
		VendorID:   ids[0],
		Name:       parts[1],
		SubDevices: make([]*Device, 0),
	}
}

func lineToVendor(line string) *Vendor {
	line = strings.TrimSpace(line)
	lines := strings.SplitN(line, " ", 2)
	return &Vendor{
		Name:     lines[1],
		VendorID: lines[0],
		Devices:  make(map[string]*Device),
	}
}

func lineToClass(line string) *Class {
	line = strings.TrimPrefix(line, "C ")
	line = strings.TrimSpace(line)
	line = strings.ReplaceAll(line, "  ", " ")
	segments := strings.SplitN(line, " ", 2)

	class := &Class{
		ClassID: segments[0],
		Name:    segments[1],
	}
	return class
}
